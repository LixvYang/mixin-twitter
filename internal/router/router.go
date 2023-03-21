package router

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/lixvyang/mixin-twitter/internal/utils/cors"
	"github.com/spf13/viper"
)

// func createMyRender() multitemplate.Renderer {
// 	p := multitemplate.NewRenderer()
// 	p.AddFromFiles("front", "./build/index.html")
// 	return p
// }

var store = cookie.NewStore([]byte("secret"))

func InitRouter(signal chan os.Signal) {
	var h r
	gin.SetMode(viper.GetString("server.AppMode"))
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	// r.HTMLRender = createMyRender()
	r.Use(gin.Logger(), gin.Recovery(), cors.Cors())
	r.Use(sessions.Sessions("_mixin_twitter", store))
	// r.Static("/static", "./dist/static")
	// r.StaticFile("/favicon.ico", "./build/favicon.ico")

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(200, "front", nil)
	// })

	api := r.Group("api/v1")
	h.HandleOauthRouter(api)
	h.HandlePraiseCommentRouter(api)
	h.HandlePraiseTwitterRouter(api)
	h.HandleTwitterRouter(api)
	h.HandleTwitterCommentRouter(api)
	h.HandleUserRouter(api)

	// r.Run(viper.GetString("server.HttpPort"))
	r.Run(":3000")
}
