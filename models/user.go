package models

import "time"

// User tables users
type User struct {
	ID           int64      `json:"id"`
	Username     string     `json:"username" sql:"unique_index"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"password_hash"`
	CreatedAt    time.Time  `json:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index"`
}
