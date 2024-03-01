package service

import (
	"github.com/hoywu/budgetServer/consts"
	"github.com/hoywu/budgetServer/dao"
	"github.com/hoywu/budgetServer/model"
	"github.com/jinzhu/copier"
)

func Login(username, password string) (user *model.User, err error) {
	user, err = dao.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		err = consts.ErrPassword
		return nil, err
	}
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
		return nil, err
	}
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
		return nil, err
	}
	return user, nil
}
