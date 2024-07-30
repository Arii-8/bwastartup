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
	// user memasukkan input (email & password)
	// mapping dari input user ke input struct
	// input struct passing service
	// di service mencari dengan bantuan repository user dengan email x
	// mencocokan password

	// input ditangkap handler
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)

	// check error
	if err != nil {
		errors := helper.FormatValidationError(err) // memanggil format validation erro dengan mengirim kan parameter 'err'
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentoken")

	// menggunakan response dari helper/helper.go
	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
