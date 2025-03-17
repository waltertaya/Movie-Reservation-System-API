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
	Movie Movie
	Date string
	Time string
	Seats int
	Booked_seats int
}

type Reservation struct {
	gorm.Model
	User User
	Showtimes Showtime
	Seat_number int
	Status string // pending, confirmed, cancelled
}
