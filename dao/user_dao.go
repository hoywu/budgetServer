package dao

import (
	"github.com/hoywu/budgetServer/db"
	"github.com/hoywu/budgetServer/model"
)

func CreateUser(user *model.User) error {
	return db.DB.Create(user).Error
}

func GetUserByID(id uint) (*model.User, error) {
	var user model.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := db.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func UpdateUsername(id uint, username string) error {
	return db.DB.Model(&model.User{}).Where("id = ?", id).Update("username", username).Error
}

func UpdatePassword(id uint, password string) error {
	return db.DB.Model(&model.User{}).Where("id = ?", id).Update("password", password).Error
}

func UpdateNickname(id uint, nickname string) error {
	return db.DB.Model(&model.User{}).Where("id = ?", id).Update("nickname", nickname).Error
}

func UpdateAvatar(id uint, avatar string) error {
	return db.DB.Model(&model.User{}).Where("id = ?", id).Update("avatar", avatar).Error
}
