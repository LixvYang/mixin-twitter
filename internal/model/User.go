package model

import (
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	FullName  string `gorm:"type:varchar(50);not null" json:"full_name"`
	Uid       string `gorm:"type:varchar(36);index;" json:"uid"`
	AvatarUrl string `gorm:"type:varchar(255);not null" json:"avatar_url"`
	SessionId string `gorm:"type:varchar(50);" json:"session_id"`

	gorm.Model
}

func CreateUser(u *User) int {
	if err := db.Create(&u).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// Delete user
func DeleteUser(uid string) int {
	if err := db.Where("uid = ?", uid).Delete(User{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// CheckUser 查询用户是否存在
func CheckUser(user_id string) int {
	var user User
	db.Model(&User{}).Where("mixin_uuid = ?", user_id).Last(&user)
	if user.Uid == "" {
		return errmsg.ERROR //1001
	}
	return errmsg.SUCCSE
}
