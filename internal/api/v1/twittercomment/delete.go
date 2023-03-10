package twittercomment

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lixvyang/mixin-twitter/internal/api/v1"
	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
)

type DeleteRequest struct {
	Id  int    `json:"id"`
	Uid string `json:"uid"`
}

func DeleteTwitterComment(c *gin.Context) {
	var r DeleteRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	tc, code := model.GetTwitterCommentById(r.Id)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	if tc.FromUid != r.Uid {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	code = model.DeleteTwitterComment(r.Id)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_DELETE_CATENAME, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
