package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	v1 "github.com/lixvyang/mixin-twitter/internal/api/v1"
	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
)

type GetUserReq struct {
	Uid string `json:"uid"`
}

func GetUserInfoByUserId(c *gin.Context) {
	fmt.Println("来请求了")
	var r GetUserReq
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	data, code := model.GetUser(r.Uid)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, data)
}
