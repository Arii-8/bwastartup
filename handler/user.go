// handler yang khusus menangani urusan yang ada didalam fungsi fungsi folder user
package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service // panggil package 'user' dan panggil struct 'Service'
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// Fungsi untuk register user
func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service

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

	// user formatter
	formatter := user.FormatUser(newUser, "tokentokentoken")

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

	// Memformat data user yang berhasil login, termasuk pembuatan token (dummy token dalam contoh ini)
	formatter := user.FormatUser(loggedinUser, "tokentokentoken")

	// Menggunakan helper untuk membuat respons sukses
	response := helper.APIResponse("Successfully logged in", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// Fungsi baru untuk handler 'check email availability'
func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// Input email di-mapping ke struct input
	// Struct input di-passing ke service
	// Service akan manggil repository - email sudah ada atau belum
	// Repository - db

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
