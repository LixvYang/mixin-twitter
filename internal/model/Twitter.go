package model

import (
	"context"
	"fmt"

	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
	"gorm.io/gorm"
)

type Twitter struct {
	gorm.Model
	Uid        string `gorm:"type:varchar(36);index:comment_from_uid" json:"from_uid"`
	Content    string `gorm:"longtext" json:"content"`
	UserName   string `gorm:"type:varchar(36)" json:"user_name"`
	UserAvatar string `gorm:"type:varchar(255)" json:"user_avatar"`
	PraiseNum  int    `gorm:"type:int(8); default 0" json:"praise_num"`
}

func CreateTwitter(t *Twitter) int {
	if err := db.Create(&t).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func DeleteTwitter(id int) int {
	if err := db.Where("id = ?", id).Delete(&Twitter{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func CheckTwitterLength() int64 {
	var count int64
	err := db.Model(&Twitter{}).Where("id >= ?", 0).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

func ListTwitters(preId int64, pageSize int64) ([]Twitter, int) {
	var twitters []Twitter
	err := db.Debug().WithContext(context.Background()).Where("id <= ?", preId).Order("id desc").Limit(int(pageSize)).Find(&twitters).Error
	if err != nil {
		fmt.Println(222)
		return twitters, errmsg.ERROR
	}
	return twitters, errmsg.SUCCSE
}
