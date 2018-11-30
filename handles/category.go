package handles

import (
	"github.com/cepljxiongjun/hotwind/models"
	"github.com/cepljxiongjun/hotwind/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Category all func
var Category category

type category struct {
}

func (category) Get(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))
	post, err := service.Cat.Get(id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "404 NOT FOUNT",
		})
		return
	}
	c.JSON(200, post)
}

type categoryForm struct {
	Name string `form:"name" binding:"required"`
}

func (category) Add(c *gin.Context) {
	var form categoryForm
	err := c.Bind(&form)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	p := models.Category{
		Name: form.Name,
	}
	service.Cat.Create(p)
	c.JSON(200, gin.H{
		"message": "ok",
	})
	return
}

func (category) Update(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))
	var form categoryForm
	err := c.Bind(&form)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := service.Cat.Update(id, form); err != nil {
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

func (category) Delete(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))
	if err := service.Cat.Delete(id); err != nil {
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
