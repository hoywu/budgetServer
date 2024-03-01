package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoywu/budgetServer/dto"
	"github.com/hoywu/budgetServer/dto/request"
	"github.com/hoywu/budgetServer/dto/response"
	"github.com/hoywu/budgetServer/log"
	"github.com/hoywu/budgetServer/service"
	"github.com/hoywu/budgetServer/utils"
)

func NewRecord(c *gin.Context) {
	var req request.RecordCreateRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	if !service.IsCategoryIDExist(req.CategoryID) {
		c.JSON(http.StatusBadRequest, dto.ErrorResp(400, "Category not exist"))
		return
	}
	if !utils.IsTimeValid(req.Time) {
		c.JSON(http.StatusBadRequest, dto.ErrorResp(400, "Time format invalid"))
		return
	}

	record, err := service.NewRecord(c.GetUint("uid"), &req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(500, "Create record failed"))
		log.ERROR("Create record failed: [UID %d] %v", c.GetUint("uid"), err)
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResp(&response.RecordCreateResp{
		Record: record,
	}))
}

func UpdateRecord(c *gin.Context) {
	var req request.RecordUpdateRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	if !service.IsRecordIDExist(req.ID) {
		c.JSON(http.StatusBadRequest, dto.ErrorResp(400, "Record not exist"))
		return
	}
	if !service.IsCategoryIDExist(req.Data.CategoryID) {
		c.JSON(http.StatusBadRequest, dto.ErrorResp(400, "Category not exist"))
		return
	}
	if !utils.IsTimeValid(req.Data.Time) {
		c.JSON(http.StatusBadRequest, dto.ErrorResp(400, "Time format invalid"))
		return
	}

	err := service.UpdateRecord(c.GetUint("uid"), &req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(500, "Update record failed"))
		log.ERROR("Update record failed: [UID %d] [RID %d] %v", c.GetUint("uid"), req.ID, err)
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResp())
}

func RemoveRecord(c *gin.Context) {
	var req request.RecordRemoveRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	err := service.RemoveRecord(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(500, "Remove record failed"))
		log.ERROR("Remove record failed: [UID %d] [RID %d] %v", c.GetUint("uid"), req.ID, err)
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResp())
}

func GetRecordList(c *gin.Context) {
	limit := c.DefaultQuery("limit", "999999999")
	limitInt, err := strconv.Atoi(limit)
	records, err := service.GetRecordList(c.GetUint("uid"), limitInt)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(500, "Get record list failed"))
		log.ERROR("Get record list failed: [UID %d] %v", c.GetUint("uid"), err)
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResp(&response.RecordGetListResp{
		Records: records,
	}))
}
