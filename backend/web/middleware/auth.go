package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleWare(needLogin bool) gin.HandlerFunc {

	return func(c *gin.Context) {
		// TODO
		c.Next()
	}
}
