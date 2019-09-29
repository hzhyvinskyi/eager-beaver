package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID				string
	FirstName		string
	LastName		string
	Email			string
	Password		string
	PurePassword	string
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

// BeforeCreate ...
func (u *User) BeforeCreate() error {
	if len(u.PurePassword) > 0 {
		enc, err := encryptString(u.PurePassword)
		if err != nil {
			return err
		}

		u.Password = enc
	}

	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
