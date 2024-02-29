package service

import (
	"github.com/hoywu/budgetServer/dao"
	"github.com/hoywu/budgetServer/dto/request"
	"github.com/hoywu/budgetServer/model"
	"github.com/jinzhu/copier"
)

func NewBudget(uid uint, req *request.BudgetCreateRequest) (*model.Budget, error) {
	budget := &model.Budget{
		UserID: uid,
	}
	copier.Copy(budget, req)
	err := dao.CreateBudget(budget)
	if err != nil {
		return nil, err
	}
	return budget, nil
}

func UpdateBudget(uid uint, req *request.BudgetUpdateRequest) error {
	budget := &model.Budget{
		UserID: uid,
	}
	copier.Copy(budget, req.Data)
	budget.ID = req.ID
	return dao.UpdateBudget(budget)
}

func RemoveBudget(id uint) error {
	return dao.DeleteBudget(id)
}

func GetBudgetList(uid uint) (*[]model.Budget, error) {
	return dao.GetBudgetsByUserID(uid)
}

func IsBudgetIDExist(id uint) bool {
	_, err := dao.GetBudgetByID(id)
	return err == nil
}
