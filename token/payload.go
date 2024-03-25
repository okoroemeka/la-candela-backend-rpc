package token

import (
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	Username  string    `json:"username"`
	ID        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
