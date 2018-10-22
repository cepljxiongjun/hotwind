package middleware

import (
	"github.com/cepljxiongjun/hotwind/helper"
	"github.com/gin-gonic/gin"
	"strings"
)

func JwtMiddleware() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		jwt := ctx.GetHeader("Authorization")
		authHeaderParts := strings.Split(jwt, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			ctx.AbortWithStatus(401)
		}
		if err := helper.JwtDecode(authHeaderParts[1]); err != nil {
			ctx.AbortWithError(401, err)
			return
		}
		ctx.Next()
	}
}