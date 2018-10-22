package models

type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
