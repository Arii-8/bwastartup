package main

// import library
import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

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

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// user handler
	userHandler := handler.NewUserHandler(userService)

	// Router
	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)

	router.Run()

	// userRepository.Save(user)

	// Last Episode 6.3 Tutorial (BERHASIL DITEST DI POSTMAN)

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
}

// input
// handler mapping input ke struct
// service mapping ke struct User
// repository save struct User ke db
