package praisetwitter

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lixvyang/mixin-twitter/internal/api/v1"
	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
)

type CheckPraiseTwitterReq struct {
	Uid string   `json:"uid"`
	Tid uint `json:"tid"`
}

func CheckPraiseTwitter(c *gin.Context) {
	var r CheckPraiseTwitterReq

	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}
	code := model.CheckIfPraise(r.Uid, r.Tid)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, nil)
	return
}
