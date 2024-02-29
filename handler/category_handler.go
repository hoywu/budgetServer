package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoywu/budgetServer/dto"
	"github.com/hoywu/budgetServer/dto/request"
	"github.com/hoywu/budgetServer/dto/response"
	"github.com/hoywu/budgetServer/service"
	"github.com/jinzhu/copier"
)

func NewCategory(c *gin.Context) {
	var req request.CategoryCreateRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResp(400, "Category name is required"))
		return
	}

	if service.IsCategoryExist(c.GetUint("uid"), req.Name) {
		c.JSON(http.StatusBadRequest, dto.ErrorResp(400, "Category already exists"))
		return
	}

	category, err := service.NewCategory(c.GetUint("uid"), &req)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(1, "Create category failed"))
		return
	}

	categoryDto := &response.CategoryDTO{}
	copier.Copy(categoryDto, category)

	c.JSON(http.StatusOK, dto.SuccessResp(
		response.CategoryCreateResp{Category: categoryDto},
	))
}

func RemoveCategory(c *gin.Context) {
	var req request.CategoryRemoveRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResp(400, "Category name is required"))
		return
	}

	err := service.RemoveCategory(c.GetUint("uid"), req.Name)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(1, "Remove category failed"))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResp())
}

func GetCategoryList(c *gin.Context) {
	categoryList, err := service.GetCategoryList(c.GetUint("uid"))
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(1, "Get category list failed"))
		return
	}

	cs := make([]response.CategoryDTO, 0, len(*categoryList))
	for _, c := range *categoryList {
		categoryDto := response.CategoryDTO{}
		copier.Copy(&categoryDto, &c)
		cs = append(cs, categoryDto)
	}
	c.JSON(http.StatusOK, dto.SuccessResp(
		response.CategoryListResp{Categories: &cs},
	))
}
