package handles

import (
	"encoding/json"
	"github.com/cepljxiongjun/hotwind/drives"
	"github.com/cepljxiongjun/hotwind/helper"
	"github.com/cepljxiongjun/hotwind/models"
	"github.com/cepljxiongjun/hotwind/repository"
	"github.com/cepljxiongjun/hotwind/repository/user"
	"github.com/gin-gonic/gin"
)

func NewUserRepo(db *drives.DB) *User {
	return &User{
		repo: user.NewPgUserRepo(db.SQL),
	}
}

type User struct {
	repo repository.UserRepo
}

func (u *User) Create(ctx *gin.Context)  {
	users := models.User{}
	json.NewDecoder(ctx.Request.Body).Decode(&users)
	newID, err := u.repo.Create(ctx.Request.Context(), &users)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err,
		})
		return
	}
	users.ID = newID
	ctx.JSON(200,  gin.H{
		"message":users,
	})
}

func (u *User) Fetch(ctx *gin.Context) {
	username := ctx.Param("username")
	users, err := u.repo.Fetch(ctx.Request.Context(), username)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(200,  users)
}

func (u *User) Auth(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	users, err := u.repo.Fetch(ctx.Request.Context(), username)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	if helper.CheckPassword(users.PasswordHash, password) == false {
		ctx.JSON(401, gin.H{
			"message": "auth failed",
		})
		return
	}
	token , _ := helper.JwtEncode(users)
	ctx.JSON(200, gin.H{
		"token": token,
	})
	return
}


