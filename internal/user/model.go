package user

import (
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Username  string    `json:"username" db:"username"`
	Bio       string    `json:"bio" db:"bio"`
	Avatar    string    `json:"avatar" db:"avatar"`
	Phone     string    `json:"phone" db:"phone"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password -`
	Status    string    `json:"status" db:"status"`
}
