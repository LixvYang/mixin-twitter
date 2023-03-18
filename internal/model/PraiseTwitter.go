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

func DeletePraiseTwitter(tid uint, from_uuid string) int {
	if err := db.Where("tid = ? AND from_uuid = ?", tid, from_uuid).Delete(&PraiseTwitter{}).Error; err != nil {
		return errmsg.ERROR
	}
	db.Model(&Twitter{}).Where("id = ?", tid).UpdateColumn("praise_num", gorm.Expr("praise_num - ?", 1))
	return errmsg.SUCCSE
}

// uid from_uuid
func CheckIfPraise(uid string, tid uint) int {
	var praiseTwitter PraiseTwitter
	if err := db.Where("from_uuid = ? AND tid = ?", uid, tid).Last(&praiseTwitter).Error; err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
