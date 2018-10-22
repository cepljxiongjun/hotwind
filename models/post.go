package models

type Post struct {
	ID    int64 `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	CreatedAt string`json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
