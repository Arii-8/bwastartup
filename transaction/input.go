package transaction

import "bwastartup/user"

type GetTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

// struct CreateTransactionInput
type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User       user.User
}

// struct TransactionNotificationInput
type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
