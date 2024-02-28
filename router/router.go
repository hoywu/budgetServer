package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hoywu/budgetServer/handler"
	"github.com/hoywu/budgetServer/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// public routes
	public := r.Group("/public")
	{
		public.POST("/login", handler.Login)
		public.POST("/register", handler.Register)
	}

	// private routes
	private := r.Group("/private")
	private.Use(middleware.AuthMiddleware())
	us := private.Group("/user")
	{
		us.GET("/info", handler.GetUserInfo)
	}
	ca := private.Group("/category")
	{
		ca.POST("/new", handler.NewCategory)
		ca.POST("/remove", handler.RemoveCategory)
		ca.GET("/getList", handler.GetCategoryList)
	}
	re := private.Group("/record")
	{
		re.POST("/new", handler.NewRecord)
		re.POST("/remove", handler.RemoveRecord)
		re.POST("/update", handler.UpdateRecord)
		re.GET("/getList", handler.GetRecordList)
	}

	return r
}
