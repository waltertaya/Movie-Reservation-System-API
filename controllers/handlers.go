package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Movie-Reservation-System-API/initialisation"
	"github.com/waltertaya/Movie-Reservation-System-API/models"
)

// Validate using superuser admin api key (hardcoded in .env)
// Only superuser can promote a user to admin

func PromoteUser(ctx *gin.Context) {
	superuser := os.Getenv("SUPER_ADMIN_API_KEY")
	if ctx.GetHeader("x-api-key") != superuser {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	id := ctx.Param("id")
	var user models.User

	result := initialisation.DB.Where("id = ?", id).First(&user)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "User not found",
		})
		return
	}

	user.Role = func(s string) *string { return &s }("ADMIN")
	result = initialisation.DB.Save(&user)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"error": "Error promoting the user",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User promoted to ADMIN",
		"user":    user,
	})
}
