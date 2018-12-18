package service

import (
	"github.com/cepljxiongjun/hotwind/models"
	"golang.org/x/crypto/bcrypt"
)

// User all func
var User user

type user struct{}

// create new user
func (user) Create(username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u := models.User{
		Username:     username,
		PasswordHash: string(hash),
	}
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (user) GetByID(id uint) (*models.User, error) {
	u := models.User{}
	if err := db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (user) GetByUsername(username string) (*models.User, error) {
	u := models.User{}
	if err := db.Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (user) Login() {

}
