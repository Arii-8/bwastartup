package user

import "time"

// struct User (entity) dalam bahasa pemrograman lain adalah ibaratkan sebuah model
type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
