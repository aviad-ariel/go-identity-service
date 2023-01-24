package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"stupix/helpers"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		tokenArray := strings.Split(tokenString, " ")
		if tokenArray[0] != "Bearer" {
			context.JSON(401, gin.H{"error": "Invalid token"})
			context.Abort()
			return
		}
		err := helpers.ValidateToken(tokenArray[1])
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
