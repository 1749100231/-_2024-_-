package router

import (
	"JH_2024_MJJ/internal/handler/admin"
	"JH_2024_MJJ/internal/handler/student"
	"JH_2024_MJJ/internal/handler/user"
	"JH_2024_MJJ/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"
	api := r.Group(pre)
	{
		userGroup := api.Group("/user")
		{
			/***USER
			POST 	login
			POST 	reg
			*/

			userGroup.POST("register", user.Register)
			userGroup.POST("login", user.Login)
		}
		stuGroup := api.Group("/student").Use(middleware.IsLogin)
		{
			//	/***Student
			//	POST 	post	发布
			//	GET 	post	获取
			//	DELETE	post	删除
			//	PUT		post	修改
			//	POST	report-post	举报
			//	GET		report-post	查看审批
			//	*/
			//
			stuGroup.POST("post", student.PublishPost)

			// CROS
			stuGroup.GET("post", middleware.Corss, student.GetPost)

			stuGroup.DELETE("post", student.DelPost)

			stuGroup.POST("report-post", student.ReportPost)

			stuGroup.PUT("post", student.UpdatePost)

			// CROS
			stuGroup.GET("report-post", middleware.Corss, student.QueryReport)

		}
		adminGroup := api.Group("/admin").Use(middleware.IsLogin, middleware.IsAdmin)
		{
			//	/***Admin
			//	GET 	report 获取所有未审批的被举报帖子
			//	POST 	report 审核所有未审批的被举报帖子
			//	*/
			adminGroup.GET("report", admin.QueryUnhandledReport)
			adminGroup.POST("report", admin.HandleReport)
		}
	}

}
