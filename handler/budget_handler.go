package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoywu/budgetServer/dto"
	"github.com/hoywu/budgetServer/dto/request"
	"github.com/hoywu/budgetServer/dto/response"
	"github.com/hoywu/budgetServer/log"
	"github.com/hoywu/budgetServer/service"
	"github.com/jinzhu/copier"
)

func NewBudget(c *gin.Context) {
	var req request.BudgetCreateRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	budget, err := service.NewBudget(c.GetUint("uid"), &req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(500, "Create budget failed"))
		log.ERROR("Create budget failed: [UID %v] %v", c.GetUint("uid"), err)
		return
	}

	budgetDto := &response.BudgetDTO{}
	copier.Copy(budgetDto, budget)

	c.JSON(http.StatusOK, dto.SuccessResp(
		response.BudgetCreateResp{Budget: budgetDto},
	))
}

func UpdateBudget(c *gin.Context) {
	var req request.BudgetUpdateRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	err := service.UpdateBudget(c.GetUint("uid"), &req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(500, "Update budget failed"))
		log.ERROR("Update budget failed: [UID %v] [BID %v] %v", c.GetUint("uid"), req.ID, err)
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResp())
}

func RemoveBudget(c *gin.Context) {
	var req request.BudgetRemoveRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	err := service.RemoveBudget(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(500, "Remove budget failed"))
		log.ERROR("Remove budget failed: [UID %v] [BID %v] %v", c.GetUint("uid"), req.ID, err)
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResp())
}

func GetBudgetList(c *gin.Context) {
	budgets, err := service.GetBudgetList(c.GetUint("uid"))
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(500, "Get budget list failed"))
		log.ERROR("Get budget list failed: [UID %v] %v", c.GetUint("uid"), err)
		return
	}

	bs := make([]response.BudgetDTO, 0, len(*budgets))
	for _, b := range *budgets {
		budgetDto := response.BudgetDTO{}
		copier.Copy(&budgetDto, &b)
		bs = append(bs, budgetDto)
	}
	c.JSON(http.StatusOK, dto.SuccessResp(
		response.BudgetListResp{Budgets: &bs},
	))
}
