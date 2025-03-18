package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Movie-Reservation-System-API/initialisation"
	"github.com/waltertaya/Movie-Reservation-System-API/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func UserRegistration(ctx *gin.Context) {
	var body struct {
		First_name string
		Last_name string
		Email string
		Password string
		Phone string
	}

	ctx.Bind(&body)

	user := models.User{
		First_name: &body.First_name,
		Last_name: &body.Last_name,
		Email: &body.Email,
		Password: func(s string) *string { return &s }(HashPassword(body.Password)),
		Role: func(s string) *string { return &s }("USER"),
		Phone: &body.Phone,
	}

	user_exist := initialisation.DB.Where("email = ?", body.Email).First(&user)

	if user_exist.RowsAffected > 0 {
		ctx.JSON(500, gin.H{
			"error": "User already exists",
		})
		return
	}

	result := initialisation.DB.Create(&user)
	if result.Error != nil {
		log.Fatal("Error creating the user, ", result.Error)
		ctx.JSON(500, gin.H{
			"error": "Error creating the user",
		})
		return
	}
	ctx.JSON(201, gin.H{
		"message": "User creaed successfully",
		"user": user,
	})
}

func UserLogin(ctx *gin.Context) {
	var body struct {
		Email string
		Password string
	}

	ctx.Bind(&body)

	user := models.User{
		Email: &body.Email,
	}

	result := initialisation.DB.Where("email = ?", body.Email).First(&user)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "User not found",
		})
		return
	}

	if !ComparePassword(*user.Password, body.Password) {
		ctx.JSON(401, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User logged in successfully",
		"user": user,
	})
}
