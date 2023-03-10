package model

import (
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
	"gorm.io/gorm"
)

type TwitterComment struct {
	gorm.Model
	Content        string `gorm:"longtext" json:"content"`
	FromUid        string `gorm:"type:varchar(36);index:comment_from_uid" json:"from_uid"`
	FromUserName   string `gorm:"type:varchar(36)" json:"from_user_name"`
	FromUserAvatar string `gorm:"type:varchar(255)" json:"from_user_avatar"`
	PraiseNum      int    `gorm:"type:int(8); default 0" json:"praise_num"`
	ToUid          string `gorm:"type:varchar(36);default null" json:"to_uid"`
	ToUserName     string `gorm:"type:varchar(36);default null" json:"to_user_name"`
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
