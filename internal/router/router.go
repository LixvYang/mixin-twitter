package router

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/lixvyang/mixin-twitter/internal/utils"
	"github.com/lixvyang/mixin-twitter/internal/utils/cors"
)

var store = cookie.NewStore([]byte("secret"))

func InitRouter(signal chan os.Signal) {
	var h r
	gin.SetMode(utils.Conf.Server.AppMode)
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	r.Use(gin.Logger(), gin.Recovery(), cors.Cors())
	r.Use(sessions.Sessions("_mixin_twitter", store))

	api := r.Group("api/v1")
	
	h.HandleOauthRouter(api)
	h.HandlePraiseCommentRouter(api)
	h.HandlePraiseTwitterRouter(api)
	h.HandleTwitterRouter(api)
	h.HandleTwitterCommentRouter(api)
	h.HandleUserRouter(api)

	r.Run(utils.Conf.Server.HttpPort)
}
