package dao

import (
	"fmt"

	"github.com/hoywu/budgetServer/db"
	"github.com/hoywu/budgetServer/log"
	"github.com/hoywu/budgetServer/model"
)

func CreateUser(user *model.User) error {
	log.DEBUG("CreateUser: " + user.Username)
	return db.DB.Create(user).Error
}

func GetUserByID(id uint) (*model.User, error) {
	log.DEBUG(fmt.Sprintf("GetUserByID: %d", id))
	var user model.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func GetUserByUsername(username string) (*model.User, error) {
	log.DEBUG("GetUserByUsername: " + username)
	var user model.User
	err := db.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func UpdateUsername(id uint, username string) error {
	log.DEBUG(fmt.Sprintf("UpdateUsername: %d, %s", id, username))
	return db.DB.Model(&model.User{}).Where("id = ?", id).Update("username", username).Error
}

func UpdatePassword(id uint, password string) error {
	log.DEBUG(fmt.Sprintf("UpdatePassword: %d, %s", id, password))
	return db.DB.Model(&model.User{}).Where("id = ?", id).Update("password", password).Error
}

func UpdateNickname(id uint, nickname string) error {
	log.DEBUG(fmt.Sprintf("UpdateNickname: %d, %s", id, nickname))
	return db.DB.Model(&model.User{}).Where("id = ?", id).Update("nickname", nickname).Error
}

func UpdateAvatar(id uint, avatar string) error {
	log.DEBUG(fmt.Sprintf("UpdateAvatar: %d, %s", id, avatar))
	return db.DB.Model(&model.User{}).Where("id = ?", id).Update("avatar", avatar).Error
}
