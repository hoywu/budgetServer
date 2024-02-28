package service

import (
	"github.com/hoywu/budgetServer/dao"
	"github.com/hoywu/budgetServer/dto/request"
	"github.com/hoywu/budgetServer/model"
	"github.com/jinzhu/copier"
)

func NewRecord(uid uint, req *request.RecordCreateRequest) (record *model.Record, err error) {
	record = &model.Record{
		UserID: uid,
	}
	copier.Copy(record, req)
	err = dao.CreateRecord(record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func UpdateRecord(uid uint, req *request.RecordUpdateRequest) error {
	record := &model.Record{
		UserID: uid,
	}
	copier.Copy(record, req.Data)
	record.ID = req.ID
	return dao.UpdateRecord(record)
}

func RemoveRecord(id uint) error {
	return dao.DeleteRecord(id)
}

func GetRecordList(uid uint, limit int) (*[]model.Record, error) {
	return dao.GetRecordsByUserID(uid, limit)
}

func IsRecordIDExist(id uint) bool {
	_, err := dao.GetRecordByID(id)
	return err == nil
}
