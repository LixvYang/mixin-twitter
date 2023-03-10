package model

import (
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
	"gorm.io/gorm"
)

type PraiseComment struct {
	gorm.Model
	Tid      uint
	ToUuid   string `gorm:"type:varchar(36);index:cid_uid_praise" json:"to_uuid"`
	FromUuid string `gorm:"type:varchar(36);index:cid_uid_praise" json:"from_uuid"`
}

func CreatePraiseComment(p *PraiseComment) int {
	if err := db.Create(&p).Error; err != nil {
		return errmsg.ERROR
	}
	db.Model(&TwitterComment{}).Where("id = ?", p.Tid).UpdateColumn("praise_num", gorm.Expr("praise_num + ?", 1))
	return errmsg.SUCCSE
}

func DeletePraiseComment(id uint, tid uint) int {
	if err := db.Where("id = ?", id).Delete(&PraiseComment{}).Error; err != nil {
		return errmsg.ERROR
	}
	db.Model(&TwitterComment{}).Where("id = ?", tid).UpdateColumn("praise_num", gorm.Expr("praise_num - ?", 1))
	return errmsg.SUCCSE
}
