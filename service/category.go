package service

import (
	"errors"

	"github.com/cepljxiongjun/hotwind/models"
)

// Cat wraps post service for easily use.
var Cat category

type category struct{}

func (category) Get(id uint) (*models.Category, error) {
	p := models.Category{}
	if db.First(&p, id).RecordNotFound() {
		return nil, errors.New("RecordNotFound")
	}
	return &p, nil
}

func (category) Create(cat models.Category) error {
	if err := db.Create(&cat).Error; err != nil {
		return err
	}
	return nil
}

func (category) Update(id uint, param interface{}) error {
	p := models.Category{}
	if err := db.Model(&p).Where("id = ?", id).Update(param).Error; err != nil {
		return err
	}
	return nil
}

func (category) Delete(id uint) error {
	p := models.Category{}
	if err := db.Where("id = ?", id).Delete(&p).Error; err != nil {
		return err
	}
	return nil
}
