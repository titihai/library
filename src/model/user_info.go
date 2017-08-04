package model

import "github.com/jinzhu/gorm"
import "time"

// UserInfo 用户信息表
type UserInfo struct {
	gorm.Model
	UserId    int       `gorm:"not null"`                  //用户id
	NickName  string    `gorm:"not null;type:varchar(50)"` // 用户昵称
	BirthDay  time.Time //`gorm:"not null"`                   //用户生日
	Address   string    `gorm:"not null;type:varchar(200)"` //用户地址
	Sex       int       `gorm:"not null"`                   // 0-无 1-男 2-女
	Interest  string    `gorm:"not null;type:varchar(500)"` // 兴趣爱好
	HeadImage string    `gorm:"not null;type:varchar(400)"` //用户头像
}
