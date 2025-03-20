// handler yang khusus menangani urusan yang ada didalam fungsi fungsi folder user
package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service // panggil package 'user' dan panggil interface 'Service'
	authService auth.Service // panggil package 'auth' dan panggil interface 'Service'
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// Fungsi untuk register user
func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input) // check error jika ada form input data yang tidak dimasukkan
	if err != nil {

		errors := helper.FormatValidationError(err) // memanggil format validation erro dengan mengirim kan parameter 'err'
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// register user (memasukkan user baru)
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Token & user formatter
	token, err := h.authService.GenerateToken(newUser.ID) // memasukkan new user id/id user baru untuk digenerate token
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, token)

	// menggunakan response dari helper/helper.go
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// Fungsi Login endpoint
func (h *userHandler) Login(c *gin.Context) {
	// Menerima input dari user (email & password)
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)

	// Memeriksa apakah terjadi error saat parsing input JSON
	if err != nil {
		// Memformat error validasi ke dalam bentuk yang mudah dipahami
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		// Membuat respons dengan pesan "Login failed" dan detail error
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Mencoba melakukan login dengan input yang diberikan
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		// Jika terjadi error saat login, mengirim respons error
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Memformat data user yang berhasil login, termasuk pembuatan token
	token, err := h.authService.GenerateToken(loggedinUser.ID) // memasukkan new user id/id user baru untuk digenerate token
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, token)

	// Menggunakan helper untuk membuat respons sukses
	response := helper.APIResponse("Successfully logged in", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// Fungsi baru untuk handler 'check email availability'
func (h *userHandler) CheckEmailAvailability(c *gin.Context) {

	// Ada input email dari user
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)

	// Memeriksa apakah terjadi error saat parsing input JSON
	if err != nil {
		// Memformat error validasi ke dalam bentuk yang mudah dipahami
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		// Membuat respons dengan pesan "email checking failed" dan detail error
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Panggil 'IsEmailAvailable' dari service
	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if IsEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

// Fungsi handler Microuser 'UploadAvatar'
func (h *userHandler) UploadAvatar(c *gin.Context) {
	// Input dari user
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// id user tuh harusnya dapet dari jwt
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	// Simpan gambarnya di folder "images/" berdasarkan id di filename
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Di service kita panggil repo
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	// JWT (sementara hardcore, seakan akan user yang login ID = 1)
	// Repo ambil data user yang ID = 1

	// Repo update data user simpan lokasi file
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

// Fungsi FetchUser
func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	formatter := user.FormatUser(currentUser, "")
	response := helper.APIResponse("Successfuly fetsh user data", http.StatusOK, "success", formatter)

	// response
	c.JSON(http.StatusOK, response)
}
