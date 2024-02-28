package dao

import (
	"time"

	"github.com/hoywu/budgetServer/db"
	"github.com/hoywu/budgetServer/model"
)

func CreateToken(token *model.Token) error {
	return db.DB.Create(token).Error
}

func GetToken(token string) (*model.Token, error) {
	var t model.Token
	if err := db.DB.Where("token = ?", token).First(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func GetUserTokens(uid uint) ([]*model.Token, error) {
	var tokens []*model.Token
	err := db.DB.Where("user_id = ?", uid).Find(&tokens).Error
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func RefreshToken(token string) error {
	newExpireTime := time.Now().Add(time.Hour * 24 * 7).Unix()
	return db.DB.Model(&model.Token{}).Where("token = ?", token).Update("expire_time", newExpireTime).Error
}

func DeleteToken(token string) error {
	return db.DB.Where("token = ?", token).Delete(&model.Token{}).Error
}

func DeleteTokensByUserID(uid uint) error {
	return db.DB.Where("user_id = ?", uid).Delete(&model.Token{}).Error
}
