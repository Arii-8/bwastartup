package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// parameter di uri
// tangkap parameter mapping ke input struct
// panggil service, input struct sebagai parameter
// service, berbekal campaign id bisa panggil repository
// repository mencari data transaction suatu campaign

type transactionHandler struct {
	service transaction.Service
}

// function NewTransactionHandler
func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

// function GetCampaignTransactions
func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign's transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

// function GetUserTransactions
func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		// check jika gagal
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// jika berhasil
	response := helper.APIResponse("User's transactions", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

// GetUserTransactions
// Handler
// Ambil nilai user dari jwt/middleware
// Service
// Repository => ambil data transactions (preload campaign)
