package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controller "github.com/waltertaya/Movie-Reservation-System-API/controllers"
	"github.com/waltertaya/Movie-Reservation-System-API/initialisation"
)

func init() {
	initialisation.LoadEnv()
	initialisation.ConnectDB()
}


func main() {

	router := gin.Default()

	fmt.Println("Starting the server at http://localhost:3000")

	// Auth: user
	router.POST("/api/v1/auth/signup", controller.UserRegistration)
	// Auth: login
	router.POST("/api/v1/auth/login", controller.UserLogin)

	// promote: USER to ADMIN
	router.PUT("/api/v1/auth/promote/:id", controller.PromoteUser)

	router.Run()
}
