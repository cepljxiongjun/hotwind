package service

import (
	"errors"

	"github.com/cepljxiongjun/hotwind/models"
)

// Post wraps post service for easily use.
var Post post

type post struct{}

func (post) Get(id int) (*models.Post, error) {
	p := models.Post{}
	if db.First(&p, id).RecordNotFound() {
		return nil, errors.New("RecordNotFound")
	}
	return &p, nil
}
