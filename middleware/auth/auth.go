package auth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/0014526/gin-demo/util/util"
)

// auth middleware(权限验证middle ware)
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		code:=http.StatusOK
		var data interface{}
		token:=context.Query("token")
		if token == "" {
			code = http.StatusBadRequest
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = http.StatusGatewayTimeout
				default:
					code = http.StatusUnauthorized
				}
			}
		}

		if code != http.StatusOK {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "success",
				"data": data,
			})

			context.Abort()
			return
		}

		context.Next()
	}
}
