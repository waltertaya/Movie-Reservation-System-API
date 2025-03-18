package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	First_name *string `json:"first_name" validate:"required,min=2,max=100"`
	Last_name *string `json:"last_name" validate:"required,min=2,max=100"`
	Email *string `json:"email" validate:"email,required"`
	Password *string `json:"Password" validate:"required,min=6"`
	Role *string `json:"role" validate:"required,eq=ADMIN|eq=USER"`// admin, user         
    Phone *string `json:"phone" validate:"required"`
}

type Movie struct {
	gorm.Model
	Title string `json:"title" validate:"required,min=2,max=100"`
	Description string `json:"description" validate:"required,min=2,max=100"`
	Genre string `json:"genre" validate:"required,min=2,max=100"`
	Poster_url string `json:"poster_url" validate:"required,min=2,max=100"`
}

type Showtime struct {
	gorm.Model
	MovieID uint    // Foreign key
	Movie   Movie   `gorm:"foreignKey:MovieID"`
	Date    time.Time `json:"date" validate:"required"`
	Time    time.Time `json:"time" validate:"required"`
	Seats   int    `json:"seats" validate:"required"`
	Booked_seats int `json:"booked_seats" validate:"required"`
}

type Reservation struct {
	gorm.Model
	UserID uint    // Foreign key
	User   User   `gorm:"foreignKey:UserID"`
	ShowtimeID uint   // Foreign key
	Showtime Showtime `gorm:"foreignKey:ShowtimeID"`
	Seat_number int `json:"seat_number" validate:"required"`
	Status string `json:"status" validate:"required,eq=pending|eq=confirmed|eq=cancelled"` // pending, confirmed, cancelled
}
