package models

import (
	"time"
)

type User struct {
	ID        int
	Username  string
	Password  string
	Admin     bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
