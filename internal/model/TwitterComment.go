package model

import "gorm.io/gorm"

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
