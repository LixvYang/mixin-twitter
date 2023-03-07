package model

import (
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
	"gorm.io/gorm"
)

type PraiseComment struct {
	gorm.Model
	Cid uint
	ToUuid    string `gorm:"type:varchar(36);index:cid_uid_praise" json:"to_uuid"`
	FromUuid  string `gorm:"type:varchar(36);index:cid_uid_praise" json:"from_uuid"`
}

func CreatePraise(p *PraiseComment) int {
	if err := db.Create(&p).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func DeletePraise(id int) int {
	if err := db.Where("id = ?", id).Delete(&PraiseComment{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
