package service

import (
	"time"

	"github.com/hoywu/budgetServer/dao"
	"github.com/hoywu/budgetServer/log"
	"github.com/hoywu/budgetServer/model"
	"github.com/hoywu/budgetServer/utils"
)

func IssueToken(uid uint, userAgent string) (token string, err error) {
	log.INFO("IssueToken: uid=%d, userAgent=%s", uid, userAgent)
	token = utils.GenerateToken()
	err = dao.CreateToken(&model.Token{
		UserID:     uid,
		Token:      token,
		UserAgent:  userAgent,
		ExpireTime: time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	return
}

func RefreshToken(token string) error {
	return dao.RefreshToken(token)
}

func IsTokenValid(token string) (isValid bool, uid uint) {
	t, err := dao.GetToken(token)
	if err != nil {
		return false, 0
	}
	if t.ExpireTime < time.Now().Unix() {
		return false, 0
	}
	return true, t.UserID
}
