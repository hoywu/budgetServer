package dao

import (
	"github.com/hoywu/budgetServer/db"
	"github.com/hoywu/budgetServer/model"
)

func CreateBudget(budget *model.Budget) error {
	return db.DB.Create(budget).Error
}

func UpdateBudget(budget *model.Budget) error {
	return db.DB.Model(&model.Budget{}).Where("id = ?", budget.ID).Updates(budget).Error
}

func DeleteBudget(id uint) error {
	return db.DB.Where("id = ?", id).Delete(&model.Budget{}).Error
}

func GetBudgetsByUserID(uid uint) (*[]model.Budget, error) {
	var budgets *[]model.Budget
	err := db.DB.Where("user_id = ?", uid).Find(&budgets).Error
	if err != nil {
		return nil, err
	}
	return budgets, nil
}

func GetBudgetByID(id uint) (*model.Budget, error) {
	var budget model.Budget
	if err := db.DB.Where("id = ?", id).First(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}
