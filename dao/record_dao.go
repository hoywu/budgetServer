package dao

import (
	"github.com/hoywu/budgetServer/db"
	"github.com/hoywu/budgetServer/model"
)

func CreateRecord(record *model.Record) error {
	return db.DB.Create(record).Error
}

func UpdateRecord(record *model.Record) error {
	return db.DB.Model(&model.Record{}).Where("id = ?", record.ID).Omit("created_at, updated_at").Save(record).Error
}

func DeleteRecord(id uint) error {
	return db.DB.Where("id = ?", id).Delete(&model.Record{}).Error
}

func GetRecordsByUserID(uid uint, limit int) (*[]model.Record, error) {
	var records *[]model.Record
	err := db.DB.Where("user_id = ?", uid).Limit(limit).Order("time desc").Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

func GetRecordByID(id uint) (*model.Record, error) {
	var record model.Record
	if err := db.DB.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}
