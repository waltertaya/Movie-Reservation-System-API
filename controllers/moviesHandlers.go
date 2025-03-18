package controllers

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Movie-Reservation-System-API/initialisation"
	"github.com/waltertaya/Movie-Reservation-System-API/models"
)

func CreateMovie(ctx *gin.Context) {

	admin_api_key := os.Getenv("ADMIN_API_KEY")

	if ctx.GetHeader("x-api-key") != admin_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var body struct {
		Title string
		Description string
		Genre string
		Poster_url string
	}

	ctx.Bind(&body)

	movie := models.Movie{
		Title: body.Title,
		Description: body.Description,
		Genre: body.Genre,
		Poster_url: body.Poster_url,
	}

	result := initialisation.DB.Create(&movie)
	if result.Error != nil {
		log.Fatal("Error creating the movie, ", result.Error)
		ctx.JSON(500, gin.H{
			"error": "Error creating the movie",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Movie created",
		"movie": movie,
	})
}

func GetMovies(ctx *gin.Context) {

	admin_api_key := os.Getenv("ADMIN_API_KEY")

	if ctx.GetHeader("x-api-key") != admin_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var movies []models.Movie
	result := initialisation.DB.Find(&movies)
	if result.Error != nil {
		log.Fatal("Error fetching the movies, ", result.Error)
		ctx.JSON(500, gin.H{
			"error": "Error fetching the movies",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"movies": movies,
	})
}

func GetMovie(ctx *gin.Context) {
	
	admin_api_key := os.Getenv("ADMIN_API_KEY")

	if ctx.GetHeader("x-api-key") != admin_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	id := ctx.Param("id")
	var movie models.Movie

	result := initialisation.DB.Where("id = ?", id).First(&movie)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "Movie not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"movie": movie,
	})
}

func UpdateMovie(ctx *gin.Context) {
	
	admin_api_key := os.Getenv("ADMIN_API_KEY")

	if ctx.GetHeader("x-api-key") != admin_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	id := ctx.Param("id")
	var movie models.Movie

	result := initialisation.DB.Where("id = ?", id).First(&movie)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "Movie not found",
		})
		return
	}

	var body struct {
		Title string
		Description string
		Genre string
		Poster_url string
	}

	ctx.Bind(&body)

	movie.Title = body.Title
	movie.Description = body.Description
	movie.Genre = body.Genre
	movie.Poster_url = body.Poster_url

	result = initialisation.DB.Save(&movie)
	if result.Error != nil {
		log.Fatal("Error updating the movie, ", result.Error)
		ctx.JSON(500, gin.H{
			"error": "Error updating the movie",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Movie updated",
		"movie": movie,
	})
}

func DeleteMovie(ctx *gin.Context) {
	
	admin_api_key := os.Getenv("ADMIN_API_KEY")

	if ctx.GetHeader("x-api-key") != admin_api_key {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	id := ctx.Param("id")
	var movie models.Movie

	result := initialisation.DB.Where("id = ?", id).First(&movie)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "Movie not found",
		})
		return
	}

	result = initialisation.DB.Delete(&movie)
	if result.Error != nil {
		log.Fatal("Error deleting the movie, ", result.Error)
		ctx.JSON(500, gin.H{
			"error": "Error deleting the movie",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Movie deleted",
	})
}
