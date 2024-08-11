package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// kerangka interface "Service" Auth JWT
type Service interface {
	GenerateToken(userID int) (string, error)       // Fungsi untuk memasukkan data kedalam token (string dan error adalah untuk token yang dihasilkan)
	ValidateToken(token string) (*jwt.Token, error) // Fungsi untuk memvalidasi token
}

// struct JWT Service
type jwtService struct {
}

// Secret key (seharusnya jangan disimpan di git karena bersifat sangat rahasia)
var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")

// Fungsi 'NewService' untuk user_id yang me-generate token
func NewService() *jwtService {
	return &jwtService{}
}

// Fungsi untuk "GenerateToken"
func (s *jwtService) GenerateToken(userID int) (string, error) {
	// data yang akan dimasukkan adalah id dari user "user_id"
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY) // verify signature
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

// Fungsi 'ValidateToken' untuk melakukan validasi token
func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		// check jika error
		if !ok {
			return nil, errors.New("invalid token")
		}

		// jika berhasil
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
