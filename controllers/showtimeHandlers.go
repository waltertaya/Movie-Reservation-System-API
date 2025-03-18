package controllers

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Movie-Reservation-System-API/initialisation"
	"github.com/waltertaya/Movie-Reservation-System-API/models"
)

func CreateShowtime(ctx *gin.Context) {
	admin_api_key := os.Getenv("ADMIN_API_KEY")

	if ctx.GetHeader("x-api-key") != admin_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var body struct {
		MovieID      uint
		Date         time.Time
		Time         time.Time
		Seats        int
		Booked_seats int
	}

	ctx.Bind(&body)

	showtime := models.Showtime{
		MovieID:      body.MovieID,
		Date:         body.Date,
		Time:         body.Time,
		Seats:        body.Seats,
		Booked_seats: body.Booked_seats,
	}

	// check if movie exists
	var movie models.Movie
	result := initialisation.DB.Where("id = ?", body.MovieID).First(&movie)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "Movie not found",
		})
		return
	}

	result = initialisation.DB.Create(&showtime)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": "Error creating the showtime",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":  "Showtime created",
		"showtime": showtime,
	})
}

func GetShowtimes(ctx *gin.Context) {
	admin_api_key := os.Getenv("ADMIN_API_KEY")

	if ctx.GetHeader("x-api-key") != admin_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var showtimes []models.Showtime
	result := initialisation.DB.Find(&showtimes)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": "Error fetching the showtimes",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"showtimes": showtimes,
	})
}

func GetShowtime(ctx *gin.Context) {
	admin_api_key := os.Getenv("ADMIN_API_KEY")

	if ctx.GetHeader("x-api-key") != admin_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	id := ctx.Param("id")
	var showtime models.Showtime

	result := initialisation.DB.Where("id = ?", id).First(&showtime)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "Showtime not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"showtime": showtime,
	})
}

func UpdateShowtime(ctx *gin.Context) {
	admin_api_key := os.Getenv("ADMIN_API_KEY")

	if ctx.GetHeader("x-api-key") != admin_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	id := ctx.Param("id")
	var showtime models.Showtime

	result := initialisation.DB.Where("id = ?", id).First(&showtime)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "Showtime not found",
		})
		return
	}

	var body struct {
		MovieID      uint
		Date         time.Time
		Time         time.Time
		Seats        int
		Booked_seats int
	}

	ctx.Bind(&body)

	showtime.MovieID = body.MovieID
	showtime.Date = body.Date
	showtime.Time = body.Time
	showtime.Seats = body.Seats
	showtime.Booked_seats = body.Booked_seats

	result = initialisation.DB.Save(&showtime)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": "Error updating the showtime",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":  "Showtime updated",
		"showtime": showtime,
	})
}

func DeleteShowtime(ctx *gin.Context) {
	admin_api_key := os.Getenv("ADMIN_API_KEY")

	if ctx.GetHeader("x-api-key") != admin_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	id := ctx.Param("id")
	var showtime models.Showtime

	result := initialisation.DB.Where("id = ?", id).First(&showtime)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "Showtime not found",
		})
		return
	}

	result = initialisation.DB.Delete(&showtime)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": "Error deleting the showtime",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Showtime deleted",
	})
}
