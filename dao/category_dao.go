package dao

import (
	"github.com/hoywu/budgetServer/db"
	"github.com/hoywu/budgetServer/model"
)

func CreateCategory(category *model.Category) error {
	return db.DB.Create(category).Error
}

func DeleteCategory(id uint) error {
	return db.DB.Where("id = ?", id).Delete(&model.Category{}).Error
}

func DeleteCategoryByName(uid uint, name string) error {
	return db.DB.Where("user_id = ? AND name = ?", uid, name).Delete(&model.Category{}).Error
}

func GetCategoriesByUserID(uid uint) (*[]model.Category, error) {
	var categories *[]model.Category
	err := db.DB.Where("user_id = ?", uid).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryByName(uid uint, name string) (*model.Category, error) {
	var category model.Category
	if err := db.DB.Where("user_id = ? AND name = ?", uid, name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func GetCategoryByID(id uint) (*model.Category, error) {
	var category model.Category
	if err := db.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
