package handles

import (
	"github.com/cepljxiongjun/hotwind/service"
	"github.com/gin-gonic/gin"
)

func GetPostByID(ctx *gin.Context) {
	post, err := service.Post.Get(1)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "404 NOT FOUNT",
		})
		return
	}
	ctx.JSON(200, post)
}
