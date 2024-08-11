package auth

import "github.com/dgrijalva/jwt-go"

// kerangka interface "Service" Auth JWT
type Service interface {
	GenerateToken(userID int) (string, error) // Fungsi untuk memasukkan data kedalam token (string dan error adalah untuk token yang dihasilkan)
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
