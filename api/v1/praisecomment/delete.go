package praisecomment

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lixvyang/mixin-twitter/api/v1"
	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
)

type DeleteRequest struct {
	Id  uint `json:"id"`
	Tid uint `json:"tid"`
}

func DeletePraiseComment(c *gin.Context) {
	var data DeleteRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	if code := model.DeletePraiseComment(data.Id, data.Tid); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
