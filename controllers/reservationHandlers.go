package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Movie-Reservation-System-API/initialisation"
	"github.com/waltertaya/Movie-Reservation-System-API/models"
)

func GetReservations(ctx *gin.Context) {
	
	user_api_key := os.Getenv("USER_API_KEY")

	if ctx.GetHeader("x-api-key") != user_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userID := ctx.MustGet("user_id").(uint)

	var reservations []models.Reservation
	result := initialisation.DB.Where("user_id = ?", userID).Find(&reservations)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "No reservations found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"reservations": reservations,
	})
}

func CreateReservation(ctx *gin.Context) {
	
	user_api_key := os.Getenv("USER_API_KEY")
	if ctx.GetHeader("x-api-key") != user_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var body struct {
		ShowtimeID uint
		Seat_number int
	}
	ctx.Bind(&body)

	// get the user id from the context
	userID := ctx.MustGet("user_id").(uint)

	// check if the showtime exists
	var showtime models.Showtime
	result := initialisation.DB.Where("id = ?", body.ShowtimeID).First(&showtime)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "Showtime not found",
		})
		return
	}

	// check if the seat is available
	if showtime.Seats <= showtime.Booked_seats {
		ctx.JSON(400, gin.H{
			"error": "No available seats",
		})
		return
	}

	// create the reservation
	reservation := models.Reservation{
		UserID:      userID,
		ShowtimeID:  body.ShowtimeID,
		Seat_number: body.Seat_number,
		Status:      "pending",
	}

	result = initialisation.DB.Create(&reservation)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": "Error creating the reservation",
		})
		return
	}

	// update the booked seats
	showtime.Booked_seats++
	result = initialisation.DB.Save(&showtime)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": "Error updating the showtime",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":     "Reservation created",
		"reservation": reservation,
	})
}

func CancelReservation(ctx *gin.Context) {
	user_api_key := os.Getenv("USER_API_KEY")
	if ctx.GetHeader("x-api-key") != user_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	// get the user id from the context
	userID := ctx.MustGet("user_id").(uint)

	id := ctx.Param("id")
	var reservation models.Reservation

	result := initialisation.DB.Where("id = ? AND user_id = ?", id, userID).First(&reservation)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "Reservation not found",
		})
		return
	}

	// check if the reservation is upcoming
	var showtime models.Showtime
	result = initialisation.DB.Where("id = ?", reservation.ShowtimeID).First(&showtime)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "Showtime not found",
		})
		return
	}

	if showtime.Date.Before(showtime.Time) {
		ctx.JSON(400, gin.H{
			"error": "Cannot cancel a past reservation",
		})
		return
	}

	// update the booked seats
	showtime.Booked_seats--
	result = initialisation.DB.Save(&showtime)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": "Error updating the showtime",
		})
		return
	}

	// cancel the reservation
	reservation.Status = "cancelled"
	result = initialisation.DB.Save(&reservation)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": "Error cancelling the reservation",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":     "Reservation cancelled",
		"reservation": reservation,
	})
}
