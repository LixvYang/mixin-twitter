package model

import (
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
	"gorm.io/gorm"
)

type TwitterComment struct {
	gorm.Model
	Tid            int    `json:"tid"`
	Content        string `gorm:"longtext" json:"content"`
	FromUuid       string `gorm:"type:varchar(36);index:comment_from_uid" json:"from_uuid"`
	FromUserName   string `gorm:"type:varchar(36)" json:"from_user_name"`
	FromUserAvatar string `gorm:"type:varchar(255)" json:"from_user_avatar"`
}

func GetTwitterCommentById(id int) (TwitterComment, int) {
	var twittercomment TwitterComment
	if err := db.Where("id = ?", id).Find(&twittercomment).Error; err != nil {
		return twittercomment, errmsg.ERROR
	}
	return twittercomment, errmsg.SUCCSE
}

func CreateTwitterComment(t *TwitterComment) int {
	if err := db.Create(&t).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func DeleteTwitterComment(id int) int {
	if err := db.Where("id = ?", id).Delete(&TwitterComment{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func ListTwitterComment(tid int) ([]TwitterComment, int) {
	var twitterCommentList []TwitterComment
	if err := db.Where("tid = ?", tid).Order("id desc").Find(&twitterCommentList).Error; err != nil {
		return twitterCommentList, errmsg.ERROR
	}
	return twitterCommentList, errmsg.SUCCSE
}
