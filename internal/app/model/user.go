package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation"
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

// Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.PurePassword, validation.By(requiredIf(u.Password == "")), validation.Length(6, 100)),
	)
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
