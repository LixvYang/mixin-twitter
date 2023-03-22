package twitter

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lixvyang/mixin-twitter/api/v1"
	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
)

type DelTwitterReq struct {
	Id int `json:"id"`
}

func DeleteTwitter(c *gin.Context) {
	var data DelTwitterReq
	if err := c.ShouldBindJSON(&data); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	if code := model.DeleteTwitter(data.Id); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
