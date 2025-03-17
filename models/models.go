package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Email string
	Password string
	Role string // admin, user
}

type Movie struct {
	gorm.Model
	Title string
	Description string
	Genre string
	Poster_url string
}

type Showtime struct {
	gorm.Model
	MovieID uint
	Movie Movie `gorm:"foreignKey:MovieID"`
	Date string
	Time string
	Seats int
	Booked_seats int
}

type Reservation struct {
	gorm.Model
	UserID    uint    // Foreign key
	User      User    `gorm:"foreignKey:UserID"`
	ShowtimeID uint   // Foreign key
	Showtime  Showtime `gorm:"foreignKey:ShowtimeID"`
	Seat_number int
	Status string // pending, confirmed, cancelled
}
