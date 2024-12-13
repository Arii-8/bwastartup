package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Tangkap parameter di handler
// Handler ke Service
// Service yang menentukan Repository mana yang di-call
// Repository: FindAll, FindByUserID
// DB

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	c.Query("user_id")
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.GetCampaigns(userID)

	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// -----> jadi api nya: http://api/v1/campaign/id(bisa 1, 2 atau tergantung data)
	// handler : mapping id yang di url ke struct input => service, call formatter
	// service : inputnya struct input => menangkap id di url, manggil repo
	// repository : get campaign by id

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaignDetail, err := h.service.GetCampaignByID(input)

	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

// create campaign
func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	// analisis
	// tangkap parameter dari user perlu mappimg ke input struct
	// ambil current user dari jwt/handler
	// panggil service, parameternya input struct yang sudah dimapping dari input user (dan juga buat slug)
	// panggil respository, untuk simpan data campaign baru
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	// Memeriksa apakah terjadi error saat parsing input JSON
	if err != nil {
		// Memformat error validasi ke dalam bentuk yang mudah dipahami
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		// Membuat respons dengan pesan "failed to create campaign" dan detail error
		response := helper.APIResponse("failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	// Memeriksa apakah terjadi error saat parsing input JSON
	if err != nil {
		// Membuat respons dengan pesan "failed to create campaign" dan detail error
		response := helper.APIResponse("failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response jika berhasil create campaign
	response := helper.APIResponse("success to create campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

// update campaign
func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	// user masukkan input
	// handler
	// mapping dari input ke input struct (ada 2 mapping, yang satu dari user dan yang satunya lagi dari uri)
	// input dari user, dan juga yang ada di uri (passing ke service)
	// service (find campaign by id, tangkap parameter)
	// repository update data campaign

	var inputID campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)
	// Memeriksa apakah terjadi error saat parsing input JSON
	if err != nil {
		// Memformat error validasi ke dalam bentuk yang mudah dipahami
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		// Membuat respons dengan pesan "failed to update campaign" dan detail error
		response := helper.APIResponse("failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response jika berhasil update campaign
	response := helper.APIResponse("success to update campaign", http.StatusOK, "success", campaign.FormatCampaign(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

// upload campaign image
func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// id user tuh harusnya dapet dari jwt
	// currentUser := c.MustGet("currentUser").(user.User)
	// userID := currentUser.ID

	// Simpan gambarnya di folder "images/" berdasarkan id di filename
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Di service kita panggil repo
	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	// JWT (sementara hardcore, seakan akan user yang login ID = 1)
	// Repo ambil data user yang ID = 1

	// Repo update data user simpan lokasi file
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Campaign image successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

// upload campaign image blueprint:
// handler:
// tangkap input dan ubah ke struct input
// save image campaign ke suatu folder

// service: (kondisi manggil point 2 di repo, panggil repo point 1 )
// => Perlu check yang user masukkan itu dia sebagai is_primary atau gak? jika gak, kalau is_primary nya false maka tidak perlu diubah

// repository:
// 1. create image/save data image ke dalam table campaign_images
// 2. ubah is_primary true ke false (ini kasus kalau kita melakukan upload data image jika is_primary yang sebelumnya true maka kita ubah menjadi false)
