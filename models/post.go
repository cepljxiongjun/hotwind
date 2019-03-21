package models

import (
	"time"
)

// Post defines fields of table  `posts`
type Post struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CID       uint       `json:"cid"`
	CreatedAt time.Time  `json:"created_at" json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	//Category  Category   `json:"category" gorm:"-"`
}
