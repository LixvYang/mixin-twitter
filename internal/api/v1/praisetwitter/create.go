package praisetwitter

import (
	"fmt"

	"github.com/gin-gonic/gin"
	v1 "github.com/lixvyang/mixin-twitter/internal/api/v1"
	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
)

func CreatePraiseTwitter(c *gin.Context) {
	var data model.PraiseTwitter
	if err := c.ShouldBindJSON(&data); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}
	fmt.Println(data)

	if code := model.CreatePraiseTwitter(&data); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
