package handles

import (
	"github.com/cepljxiongjun/hotwind/models"
	"github.com/cepljxiongjun/hotwind/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Post all func
var Post post

type post struct{}

func (post) All(c *gin.Context) {
	post, err := service.Post.All()
	if err != nil {
		c.JSON(404, gin.H{
			"message": "404 NOT FOUNT",
		})
		return
	}
	c.JSON(200, post)
}

func (post) Get(ctx *gin.Context) {
	id := cast.ToUint(ctx.Param("id"))
	if id < 1 {
		ctx.JSON(404, gin.H{
			"message": "404 NOT FOUNT",
		})
		return
	}
	post, err := service.Post.Get(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "404 NOT FOUNT",
		})
		return
	}
	ctx.JSON(200, post)
}

type postForm struct {
	Title string `form:"title" binding:"required"`
	Body  string `form:"body" binding:"required"`
	CID   uint   `form:"cid" binding:"required"`
}

func (post) Add(ctx *gin.Context) {

	var form postForm
	err := ctx.Bind(&form)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	p := models.Post{
		Title: form.Title,
		Body:  form.Body,
		CID:   form.CID,
	}
	service.Post.Create(p)
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
	return
}

func (post) Update(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))
	var form postForm
	err := c.Bind(&form)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := service.Post.Update(id, form); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
	return
}
func (post) Delete(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))
	if err := service.Post.Delete(id); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
	return
}
