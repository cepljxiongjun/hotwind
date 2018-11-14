package service

import (
	"errors"

	"github.com/cepljxiongjun/hotwind/models"
)

// Post wraps post service for easily use.
var Post post

type post struct{}

func (post) Get(id uint) (*models.Post, error) {
	p := models.Post{}
	if db.First(&p, id).RecordNotFound() {
		return nil, errors.New("RecordNotFound")
	}
	return &p, nil
}

func (post) Create(post models.Post) error {
	if err := db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func (post) Update(id uint, param interface{}) error {
	p := models.Post{}
	if err := db.Model(&p).Where("id = ?", id).Update(param).Error; err != nil {
		return err
	}
	return nil
}

func (post) Delete(id uint) error {
	p := models.Post{}
	if err := db.Where("id = ?", id).Delete(&p).Error; err != nil {
		return err
	}
	return nil
}
