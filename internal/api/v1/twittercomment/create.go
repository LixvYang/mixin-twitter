package twittercomment

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lixvyang/mixin-twitter/internal/api/v1"
	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
)

func CreateComment(c *gin.Context) {
	var data model.TwitterComment
	if err := c.ShouldBindJSON(&data); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}
	if code := model.CreateTwitterComment(&data); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
