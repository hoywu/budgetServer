package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoywu/budgetServer/dto"
	"github.com/hoywu/budgetServer/dto/request"
	"github.com/hoywu/budgetServer/dto/response"
	"github.com/hoywu/budgetServer/model"
	"github.com/hoywu/budgetServer/service"
	"github.com/hoywu/budgetServer/utils"
)

func Login(c *gin.Context) {
	var req request.UserLoginRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	user, err := service.Login(req.Username, utils.HashPassword(req.Password))
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResp(1, "Login failed"))
		return
	}

	token, err := service.IssueToken(user.ID, c.Request.UserAgent())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResp(500, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResp(response.UserLoginResp{
		Token:    token,
		UserInfo: &user.UserInfo,
	}))
}

func Register(c *gin.Context) {
	var req request.UserRegisterRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	if service.IsUserExist(req.Username) {
		c.JSON(http.StatusOK, dto.ErrorResp(1, "User already exists"))
		return
	}

	user, err := service.Register(
		req.Username,
		utils.HashPassword(req.Password),
		&model.UserInfo{
			Nickname: req.Nickname,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResp(500, "Internal Server Error"))
		return
	}

	token, err := service.IssueToken(user.ID, c.Request.UserAgent())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResp(500, "Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResp(response.UserRegisterResp{
		Token:    token,
		UserInfo: &user.UserInfo,
	}))
}

func GetUserInfo(c *gin.Context) {
	user, err := service.GetUserInfo(c.GetUint("uid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResp(500, "Internal Server Error"))
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResp(&response.UserGetInfoResp{
		UserInfo: &user.UserInfo,
	}))
}
