package main

import (
	"fmt"
	"github.com/cepljxiongjun/hotwind/drives"
	"github.com/cepljxiongjun/hotwind/handles"
	"github.com/cepljxiongjun/hotwind/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func main()  {
	connection, err := drives.PostgreSQL()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	pHandler := handles.NewUserRepo(connection)
	r := gin.Default()
	r.POST("/auth", pHandler.Auth)
	r.POST("/user", pHandler.Create)
	r.Use(middleware.JwtMiddleware())
	r.Run(":8080")
}