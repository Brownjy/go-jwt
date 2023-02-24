package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PassWordCost = 12 //密码加密难度
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Password string
}

// GetUser 查找User
func (u *User) GetUser(userName string) (count int64, err error) {
	err = DB.Model(&User{}).Where("user_name=?", userName).First(&u).Count(&count).Error
	return
}

// CreateUser 创建User
func (u *User) CreateUser() (err error) {
	err = DB.Create(&u).Error
	return
}

// SetPassword 加密密码
func (u *User) SetPassword(password string) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	u.Password = string(bytes)
	return
}

// CheckPassword 校验密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
