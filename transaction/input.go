package transaction

import "bwastartup/user"

type GetTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
