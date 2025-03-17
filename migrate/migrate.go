package main

import (
	"github.com/waltertaya/Movie-Reservation-System-API/initialisation"
	"github.com/waltertaya/Movie-Reservation-System-API/models"
)

func init() {
	initialisation.LoadEnv()
	initialisation.ConnectDB()
}

func main() {
	initialisation.DB.AutoMigrate(&models.User{})
	initialisation.DB.AutoMigrate(&models.Movie{})
	initialisation.DB.AutoMigrate(&models.Showtime{})
	initialisation.DB.AutoMigrate(&models.Reservation{})
}
