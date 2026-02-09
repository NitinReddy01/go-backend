package user

import "github.com/NitinReddy01/go-backend/internal/model"

type User struct {
	model.Base
	Email *string `json:"email,omitempty" db:"email"`
	Phone *string `json:"phone,omitempty" db:"phone"`
	Name  string  `json:"name" db:"name"`
}
