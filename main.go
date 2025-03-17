package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Movie-Reservation-System-API/initialisation"
)

func init() {
	initialisation.LoadEnv()
}


func main() {

	router := gin.Default()

	fmt.Println("Starting the server at http://localhost:3000")

	router.GET("/", func (ctx *gin.Context) {
		ctx.JSON(
			201, gin.H{
				"message": "Successfully reached the route",
			})
	})

	router.Run()
}
