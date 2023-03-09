package twitter

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	v1 "github.com/lixvyang/mixin-twitter/internal/api/v1"
	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
)

type ListTwitterResp struct {
	List         []model.Twitter `json:"list"`
	PrePageToken string          `json:"pre_page_token"`
}

var (
	defaultPageSize int64 = 10
	defaultCursor   int64 = 10
)

func ListTwitter(c *gin.Context) {
	defaultCursor = model.CheckTwitterLength()
	pageToken := c.Query("page_token")
	page := Token(pageToken).Decode()
	var (
		cursor   int64 = defaultCursor
		pageSize int64 = defaultPageSize
	)

	if pageToken != "" {
		// 解析分页
		if page.NextTimeAtUTC > time.Now().Unix() || time.Now().Unix()-page.NextTimeAtUTC > int64(time.Hour)*24 {
			v1.SendResponse(c, errmsg.ERROR, "bad page token")
			return
		}

		// invaild
		if page.PreID <= 0 || page.NextTimeAtUTC == 0 || page.NextTimeAtUTC > time.Now().Unix() || page.PageSize <= 0 {
			v1.SendResponse(c, errmsg.ERROR, "bad page_token")
			return
		}
		fmt.Println(page)
		cursor = page.PreID
		pageSize = page.PageSize
	}

	// 查询数据库
	twittersList, err := model.ListTwitters(cursor, pageSize+1)
	if err != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, "bad page_token")
		return
	}

	var (
		hasPrePage   bool
		prePageToken string
	)
	if len(twittersList) > int(pageSize) {
		hasPrePage = true
	}

	// if has pre page
	if hasPrePage {
		prePageInfo := Page{
			PreID:         int64(twittersList[len(twittersList)-1].ID),
			NextTimeAtUTC: time.Now().Unix(),
			PageSize:      int64(pageSize),
		}
		prePageToken = string(prePageInfo.Encode())
		v1.SendResponse(c, errmsg.SUCCSE, ListTwitterResp{
			PrePageToken: prePageToken,
			List:         twittersList[:len(twittersList)-1],
		})
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, ListTwitterResp{
		PrePageToken: prePageToken,
		List:         twittersList,
	})

	return
}
