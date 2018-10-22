package helper

import (
	"errors"
	"github.com/cepljxiongjun/hotwind/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

func EncodePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

func CheckPassword(encodePW, passwordOK string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOK))
	if err != nil {
		return false
	} else {
		return true
	}
}

func JwtEncode(user *models.User)  (string, error) {
	mySigningKey := []byte("hotwind")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}
	id := strconv.FormatInt(user.ID, 10)
	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    user.Username,
			Id: id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return t, nil

}

func JwtDecode(jwtToken string) error {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("hotwind"), nil
	})

	if token.Valid {
		return nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return errors.New("Token is expired.")
		} else {
			return errors.New("Couldn't handle this token:")
		}
	} else {
		return errors.New("Couldn't handle this token:")
	}
}