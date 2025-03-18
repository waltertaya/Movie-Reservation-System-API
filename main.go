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

	// Movies
	router.POST("/api/v1/movies", controller.CreateMovie)
	router.GET("/api/v1/movies", controller.GetMovies)
	router.GET("/api/v1/movies/:id", controller.GetMovie)
	router.PUT("/api/v1/movies/:id", controller.UpdateMovie)
	router.DELETE("/api/v1/movies/:id", controller.DeleteMovie)

	// Showtimes
	router.POST("/api/v1/showtimes", controller.CreateShowtime)
	router.GET("/api/v1/showtimes", controller.GetShowtimes)
	router.GET("/api/v1/showtimes/:id", controller.GetShowtime)
	router.PUT("/api/v1/showtimes/:id", controller.UpdateShowtime)
	router.DELETE("/api/v1/showtimes/:id", controller.DeleteShowtime)

	// Reservations
	router.POST("/api/v1/reservations", controller.CreateReservation)
	router.GET("/api/v1/reservations", controller.GetReservations)
	router.DELETE("/api/v1/reservations/:id", controller.CancelReservation)

	router.Run()
}
