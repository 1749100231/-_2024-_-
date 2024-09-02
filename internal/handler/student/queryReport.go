package student

import (
	"JH_2024_MJJ/internal/service"
	"JH_2024_MJJ/pkg/utils"
	"github.com/gin-gonic/gin"
)

type QueryReportData struct {
	UserID int64 `form:"user_id" binding:"required"`
}

func QueryReport(c *gin.Context) {
	var data QueryReportData
	if err := c.ShouldBind(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		utils.Log.Printf("查询失败，参数错误")
		return
	}
	if _, err := service.GetUserByID(data.UserID); err != nil {
		utils.JsonFail(c, 200502, "用户不存在")
		utils.Log.Printf("%d用户不存在", data.UserID)
		return
	}
	err, queryData := service.QueryReport(data.UserID)
	if err != nil {
		utils.JsonFail(c, 200502, "查询失败，"+err.Error())
		utils.Log.Printf("查询失败，" + err.Error())
		return
	}
	utils.Log.Printf("%d查询成功，噢耶！", data.UserID)
	utils.JsonSuccess(c, queryData)
}
