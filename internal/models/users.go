package models

import (
	"time"
)

type Users struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
