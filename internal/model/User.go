package model

import (
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	FullName       string `gorm:"type:varchar(50);not null" json:"full_name"`
	Uid            string `gorm:"type:varchar(36);index;" json:"identity_number"`
	IdentityNumber string `gorm:"type:varchar(36);index;" json:"uid"`
	AvatarUrl      string `gorm:"type:varchar(255);not null" json:"avatar_url"`
	SessionId      string `gorm:"type:varchar(50);" json:"session_id"`

	gorm.Model
}

func GetUser(uid string) (*User, int) {
	var user *User
	if err := db.Where("uid = ?", uid).First(&user).Error; err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
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
	db.Model(&User{}).Where("uid = ?", user_id).Last(&user)
	if user.Uid == "" {
		return errmsg.ERROR //1001
	}
	return errmsg.SUCCSE
}

// EditUser 编辑用户信息
func UpdateUser(uid string, data *User) int {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return errmsg.ERROR
	}

	// 锁住指定 id 的 User 记录
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("uid = ?", uid).Error; err != nil {
		tx.Rollback()
		return errmsg.ERROR
	}

	var maps = make(map[string]interface{})
	maps["full_name"] = data.FullName
	maps["avatar_url"] = data.AvatarUrl
	maps["session_id"] = data.SessionId
	if err := db.Model(&User{}).Where("uid = ? ", uid).Updates(maps).Error; err != nil {
		return errmsg.ERROR
	}
	if err := tx.Commit().Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
