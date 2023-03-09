package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lixvyang/mixin-twitter/internal/api/v1/twitter"
	"github.com/spf13/viper"
)

type R struct{}

func InitRouter(signal chan os.Signal) {
	var h R
	gin.SetMode(viper.GetString("server.AppMode"))
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("api/v1")

	h.HandleTwitter(api)

	r.Run(viper.GetString("server.HttpPort"))
}

func (*R) HandleTwitter(c *gin.RouterGroup) {
	c.GET("/twitter/list", twitter.ListTwitter)
	c.POST("/twitter/create", twitter.CreateTwitter)
}
