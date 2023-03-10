package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter(signal chan os.Signal) {
	var h r
	gin.SetMode(viper.GetString("server.AppMode"))
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("api/v1")
	h.HandleOauthRouter(api)
	h.HandlePraiseCommentRouter(api)
	h.HandlePraiseTwitterRouter(api)
	h.HandleTwitterRouter(api)
	h.HandleTwitterCommentRouter(api)
	h.HandleUserRouter(api)

	r.Run(viper.GetString("server.HttpPort"))
}
