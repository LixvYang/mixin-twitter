package twittercomment

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/lixvyang/mixin-twitter/internal/api/v1"
	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
)

type ListTwitterCommentResp struct {
	List []model.TwitterComment `json:"list"`
}

func ListTwitterComment(c *gin.Context) {
	tid := c.Query("tid")
	r, err := strconv.Atoi(tid)
	if err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	data, code := model.ListTwitterComment(r)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, ListTwitterCommentResp{
		List: data,
	})
	return
}
