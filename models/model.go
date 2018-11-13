package models

import "time"

// Category defines fields of table `category`
type Category struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at" json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `json:"deleted_at" sql:"index"`
}
