package model

import "github.com/jinzhu/gorm"

// User 数据库用户表
type User struct {
	gorm.Model
	Account  string `gorm:"not null;type:varchar(50)"`  // 用户账号名
	Password string `gorm:"not null;type:varchar(50)"`  // 用户密码（加盐后加密）
	Salt     string `gorm:"not null;type:varchar(50)"`  // 盐
	Phone    string `gorm:"not null;type:varchar(50)"`  // 手机号
	Email    string `gorm:"not null;type:varchar(100)"` // 用户邮箱
}
