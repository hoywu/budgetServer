package response

import "github.com/hoywu/budgetServer/model"

type RecordCreateResp struct {
	Record *model.Record
}

type RecordGetListResp struct {
	Records *[]model.Record
}
