package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Movie-Reservation-System-API/controllers"
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
	router.POST("/api/v1/auth/signup", controllers.UserRegistration)
	// Auth: admin
	router.POST("/api/v1/auth/admin/signup", controllers.AdminRegistration)

	router.Run()
}
