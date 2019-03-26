package main

import (
	"fmt"

	"github.com/cepljxiongjun/hotwind/middleware"

	"github.com/cepljxiongjun/hotwind/handles"
	"github.com/cepljxiongjun/hotwind/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	parseConfig()
	service.Init()
	r := gin.Default()
	r.POST("/login", handles.User.Login)
	api := r.Group("/api").Use(middleware.JwtMiddleware())
	{
		api.GET("/posts", handles.Post.All)
		api.GET("/posts/:id", handles.Post.Get)
		api.POST("/posts", handles.Post.Add)
		api.PUT("/posts/:id", handles.Post.Update)
		api.DELETE("/posts/:id", handles.Post.Delete)
		api.GET("/categories", handles.Post.Get)
		api.GET("/categories/:id", handles.Post.Get)
		api.POST("/categories", handles.Post.Add)
		api.PUT("/categories/:id", handles.Post.Update)
		api.DELETE("/categories/:id", handles.Post.Delete)
		api.POST("/users", handles.User.Create)
	}
	r.Run(viper.GetString("port.default"))
}

func parseConfig() {
	viper.SetConfigFile("config/in-local.toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("parse config file fail: %s", err))
	}
}
