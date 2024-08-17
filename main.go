package main

// import library
import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local" // menghubungkan ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})                                  // menghubungkan ke gorm

	// check jika error
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)                        // user repository
	userService := user.NewService(userRepository)                  // user service
	authService := auth.NewService()                                // user generate token 'auth service'
	userHandler := handler.NewUserHandler(userService, authService) // user handler

	// Router
	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()

	// userRepository.Save(user)

	// Last Episode 10.2 Tutorial (BERHASIL DITEST DI POSTMAN)

	/*
	 * CLUE BLUEPRINT
	 *
	 * input dari user
	 * handler, mpping input dari user -> struct User
	 * service : melakukan mapping dari struct ke struct User
	 * repository
	 * db
	 *
	 */

	// input dari user
	// handler, mapping input dari user -> struct input
	// service : melakukan mapping dari struct input ke struct User
	// repository
	// db

	// Langkah-langkah middleware (BLUEPRINT)
	// langkah 1: ambil nilai header 'Authorization: Bearer tokentokentoken'
	// langkah 2: dari header Authorization, kita ambil nilai tokennya saja
	// Langkah 3: kita ambil validasi token
	// Langkah 4: Jika valid, ambil nilai user_id
	// Langkah 5: ambil user dari db berdasarkan user_id lewat service (membuat fungsi service)
	// Langkah 6: jika user ada, kita set context isinya user
}

// Fungsi middleware
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Contoh hasil: Bearer token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// ambil user id
		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}

// input
// handler mapping input ke struct
// service mapping ke struct User
// repository save struct User ke db
