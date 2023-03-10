package model

import (
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
	"gorm.io/gorm"
)

type PraiseTwitter struct {
	gorm.Model
	Tid      uint   `json:"tid" comment:"twitter_id"`
	ToUuid   string `gorm:"type:varchar(36);index:cid_uid_praise" json:"to_uuid"`
	FromUuid string `gorm:"type:varchar(36);index:cid_uid_praise" json:"from_uuid"`
}

func CreatePraiseTwitter(p *PraiseTwitter) int {
	if err := db.Create(&p).Error; err != nil {
		return errmsg.ERROR
	}
	db.Model(&Twitter{}).Where("id = ?", p.Tid).UpdateColumn("praise_num", gorm.Expr("praise_num + ?", 1))
	return errmsg.SUCCSE
}

func DeletePraiseTwitter(id uint, tid uint) int {
	if err := db.Where("id = ?", id).Delete(&PraiseTwitter{}).Error; err != nil {
		return errmsg.ERROR
	}
	db.Model(&Twitter{}).Where("id = ?", tid).UpdateColumn("praise_num", gorm.Expr("praise_num - ?", 1))
	return errmsg.SUCCSE
}
