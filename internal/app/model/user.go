package model

import "time"

// User ...
type User struct {
	ID			string
	FirstName	string
	LastName	string
	Email		string
	Password	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
