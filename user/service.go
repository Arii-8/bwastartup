package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)   // Service untuk register input
	Login(input LoginInput) (User, error)                 // Service untuk login input
	IsEmailAvailable(input CheckEmailInput) (bool, error) // Service untuk check 'is email available'
	SaveAvatar(ID int, fileLocation string) (User, error) // Service untuk 'Save Avatar'
	GetUserByID(ID int) (User, error)                     // Service untuk mencari user dari db berdasarkan id
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// Fungsi register user
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// Fungsi login user
func (s *service) Login(input LoginInput) (User, error) {

	// input email & password login
	email := input.Email
	password := input.Password

	// check user find by email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	// hash password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	// if return success
	return user, nil
}

// Fungsi check 'is email available'
func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

// Fungsi untuk 'SaveAvatar'
func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	// Dapatkan user berdasarkan ID
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	// Update attribute avatar file name
	user.AvatarFileName = fileLocation

	// Simpan perubahan avatar file name
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

// Fungsi 'GetUserByID'
func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	// jika user tidak ditemukan
	if user.ID == 0 {
		return user, errors.New("no user found on with that id")
	}

	// jika ditemukan/semua berjalan lancar
	return user, nil
}

// mapping struct input ke struct User
// simpan struct User melalui
