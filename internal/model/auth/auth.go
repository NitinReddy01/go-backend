package auth

import (
	"time"

	"github.com/NitinReddy01/go-backend/internal/model"
)

type AuthIdentity struct {
	model.Base

	UserID   string `json:"userId" db:"user_id"`
	Identity string `json:"identity" db:"identity"`

	Username *string `json:"username,omitempty" db:"username"`
	Password *string `json:"-" db:"password"` // dont expose in JSON

	Email *string `json:"email,omitempty" db:"email"`
	Phone *string `json:"phone,omitempty" db:"phone"`
}

type AuthSession struct {
	model.Base

	UserID     string `json:"userId" db:"user_id"`
	IdentityID string `json:"identityId" db:"identity_id"`

	ExpiresAt time.Time `json:"expiresAt" db:"expires_at"`
	IsActive  bool      `json:"isActive" db:"is_active"`
}
