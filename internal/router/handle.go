package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lixvyang/mixin-twitter/internal/api/v1/oauth"
	"github.com/lixvyang/mixin-twitter/internal/api/v1/praisecomment"
	praisetwitter "github.com/lixvyang/mixin-twitter/internal/api/v1/praisecomment"
	"github.com/lixvyang/mixin-twitter/internal/api/v1/twitter"
	"github.com/lixvyang/mixin-twitter/internal/api/v1/twittercomment"
	"github.com/lixvyang/mixin-twitter/internal/api/v1/user"
)

type r struct {
}

func (*r) HandleTwitterRouter(c *gin.RouterGroup) {
	c.GET("/twitter/list", twitter.ListTwitter)
	c.POST("/twitter/create", twitter.CreateTwitter)
}

func (*r) HandleUserRouter(c *gin.RouterGroup) {
	c.POST("/user/create", user.CreateUser)
}

func (*r) HandleOauthRouter(c *gin.RouterGroup) {
	c.GET("/oauth/redirect", oauth.MixinOauth)
}

func (*r) HandlePraiseCommentRouter(c *gin.RouterGroup) {
	c.POST("/praisecomment/create", praisecomment.CreateParaseComment)
	c.DELETE("/praisecomment/delete", praisecomment.DeleteParaseComment)
}

func (*r) HandlePraiseTwitterRouter(c *gin.RouterGroup) {
	c.POST("/praisetwitter/create", praisetwitter.CreateParaseComment)
	c.DELETE("/praisetwitter/create", praisetwitter.CreateParaseComment)
}

func (*r) HandleTwitterCommentRouter(c *gin.RouterGroup) {
	c.POST("/twittercomment/create", twittercomment.CreateComment)
	c.DELETE("/twittercomment/create", twittercomment.DeleteTwitterComment)
}
