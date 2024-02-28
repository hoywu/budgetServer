package response

import "github.com/hoywu/budgetServer/model"

type UserRegisterResp struct {
	Token    string          `json:"token"`
	UserInfo *model.UserInfo `json:"userInfo"`
}

type UserLoginResp struct {
	Token    string          `json:"token"`
	UserInfo *model.UserInfo `json:"userInfo"`
}

type UserGetInfoResp struct {
	UserInfo *model.UserInfo `json:"userInfo"`
}
