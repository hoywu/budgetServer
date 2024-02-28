package service

import (
	"github.com/hoywu/budgetServer/dao"
	"github.com/hoywu/budgetServer/dto/request"
	"github.com/hoywu/budgetServer/model"
)

func NewCategory(uid uint, req *request.CategoryCreateRequest) (category *model.Category, err error) {
	category = &model.Category{
		UserID: uid,
		Name:   req.Name,
		Icon:   req.Icon,
	}
	err = dao.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func RemoveCategory(uid uint, name string) error {
	return dao.DeleteCategoryByName(uid, name)
}

func IsCategoryExist(uid uint, name string) bool {
	_, err := dao.GetCategoryByName(uid, name)
	if err != nil {
		return false
	}
	return true
}

func IsCategoryIDExist(id uint) bool {
	_, err := dao.GetCategoryByID(id)
	if err != nil {
		return false
	}
	return true
}

func GetCategoryList(uid uint) (*[]model.Category, error) {
	return dao.GetCategoriesByUserID(uid)
}