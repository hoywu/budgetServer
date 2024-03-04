package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hoywu/budgetServer/config"
	"github.com/hoywu/budgetServer/handler"
	"github.com/hoywu/budgetServer/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	addGlobalMiddlewares(r)

	public := r.Group("/public")
	addPublicRoutes(public)

	private := r.Group("/private")
	private.Use(middleware.AuthMiddleware())
	addPrivateRoutes(private)

	return r
}

func addGlobalMiddlewares(r *gin.Engine) {
	r.Use(middleware.LoggerMiddleware())
	if config.GetPolicyConfig().AllowCORS {
		r.Use(middleware.CORSMiddleware())
	}
}

func addPublicRoutes(r *gin.RouterGroup) {
	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
}

func addPrivateRoutes(r *gin.RouterGroup) {
	us := r.Group("/user")
	{
		us.GET("/info", handler.GetUserInfo)
	}
	ca := r.Group("/category")
	{
		ca.POST("/new", handler.NewCategory)
		ca.POST("/remove", handler.RemoveCategory)
		ca.GET("/getList", handler.GetCategoryList)
	}
	re := r.Group("/record")
	{
		re.POST("/new", handler.NewRecord)
		re.POST("/remove", handler.RemoveRecord)
		re.POST("/update", handler.UpdateRecord)
		re.GET("/getList", handler.GetRecordList)
	}
	bu := r.Group("/budget")
	{
		bu.POST("/new", handler.NewBudget)
		bu.POST("/remove", handler.RemoveBudget)
		bu.POST("/update", handler.UpdateBudget)
		bu.GET("/getList", handler.GetBudgetList)
	}
}
