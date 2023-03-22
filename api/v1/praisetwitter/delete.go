package praisetwitter

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lixvyang/mixin-twitter/api/v1"
	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
)

type DeleteRequest struct {
	Tid     uint   `json:"tid"`
	FromUid string `json:"from_uid"`
}

func DeletePraiseTwitter(c *gin.Context) {
	var data DeleteRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	if code := model.DeletePraiseTwitter(data.Tid, data.FromUid); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
