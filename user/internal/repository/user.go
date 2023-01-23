package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"user/internal/service"
)

type User struct {
	UserId         int64  `gorm:"primarykey"`
	Username       string `gorm:"unique"`
	Nickname       string
	PasswordDigest string
	Model
}

const (
	PasswordCost = 12 // 密码加密难度
)

// CheckUserExist	检查用户是否存在
func (user *User) CheckUserExist(req *service.DouyinUserLoginRequest) bool {
	if err := DB.Where("username = ?", req.GetUsername()).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// ShowUserInfo	获取用户信息
func (user *User) ShowUserInfo(req *service.DouyinUserLoginRequest) error {
	if exist := user.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("User Not Exist")
}

// UserCreate 创建新用户
func (*User) UserCreate(req *service.DouyinUserRegisterRequest) (user User, err error) {
	var count int64
	DB.Where("username = ?", req.GetUsername()).Count(&count)
	if count != 0 {
		return User{}, errors.New("Username Exist")
	}
	user = User{
		Username: req.GetUsername(),
		Nickname: req.GetUsername(),
	}
	err = user.SetPassword(req.Password)
	err = DB.Create(&user).Error
	return user, err

}

// SetPassword 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 密码校验
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func (*User) UserShow(req *service.DouyinUserRequest) (user User, err error) {
	err = DB.First(&user, req.GetUserId()).Error
	return user, err
}

func BuildUser(item User) *service.User {
	return &service.User{
		Id:            item.UserId,
		Name:          item.Username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
}
