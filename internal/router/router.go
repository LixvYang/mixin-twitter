package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lixvyang/mixin-twitter/internal/utils/cors"
	"github.com/spf13/viper"
)

func InitRouter(signal chan os.Signal) {
	var h r
	gin.SetMode(viper.GetString("server.AppMode"))
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	r.Use(gin.Logger(), gin.Recovery(), cors.Cors())

	// http:127.0.0.1:8080/api/v1//user/create
	api := r.Group("api/v1")
	h.HandleOauthRouter(api)
	h.HandlePraiseCommentRouter(api)
	h.HandlePraiseTwitterRouter(api)
	h.HandleTwitterRouter(api)
	h.HandleTwitterCommentRouter(api)
	h.HandleUserRouter(api)

	r.Run(viper.GetString("server.HttpPort"))
}
