package main

import (
	"fmt"

	"github.com/cepljxiongjun/hotwind/handles"
	"github.com/cepljxiongjun/hotwind/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	parseConfig()
	service.Init()
	r := gin.Default()
	r.GET("/posts/:id", handles.Post.Get)
	r.POST("/posts", handles.Post.Add)
	r.PUT("/posts/:id", handles.Post.Update)
	r.DELETE("/posts/:id", handles.Post.Delete)
	r.Run()
}


func parseConfig() {
	viper.SetConfigFile("config/in-local.toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("parse config file fail: %s", err))
	}
}
