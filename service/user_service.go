package service

import (
	"github.com/hoywu/budgetServer/consts"
	"github.com/hoywu/budgetServer/dao"
	"github.com/hoywu/budgetServer/log"
	"github.com/hoywu/budgetServer/model"
	"github.com/jinzhu/copier"
)

func Login(username, password string) (user *model.User, err error) {
	user, err = dao.GetUserByUsername(username)
	if err != nil {
		log.INFO("Login failed: [" + username + "] " + err.Error())
		return nil, err
	}
	if user.Password != password {
		log.INFO("Login password error: " + username)
		err = consts.ErrPassword
		return nil, err
	}
	log.INFO("Login success: " + username)
	return user, nil
}

func Register(username, password string, req *model.UserInfo) (user *model.User, err error) {
	user = &model.User{
		Username: username,
		Password: password,
	}
	copier.Copy(user, req)
	err = dao.CreateUser(user)
	if err != nil {
		log.INFO("Register failed: " + username + " " + err.Error())
		return nil, err
	}
	log.INFO("Register success: " + username)
	return user, nil
}

func IsUserExist(username string) bool {
	_, err := dao.GetUserByUsername(username)
	if err != nil {
		return false
	}
	return true
}

func GetUserInfo(uid uint) (user *model.User, err error) {
	user, err = dao.GetUserByID(uid)
	if err != nil {
		log.ERROR("Get user info failed: " + err.Error())
		return nil, err
	}
	return user, nil
}
