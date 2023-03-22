package cors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lixvyang/mixin-twitter/internal/utils"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			if utils.Conf.Server.AppMode == "release" {
				c.Header("Access-Control-Allow-Origin", "http://43.143.154.162")
			} else {
				c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
			}
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
