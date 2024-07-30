package user

// Register input
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

// Login input
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// check email input
type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

/* Input struct yang berarti akan diolah oleh handler */
