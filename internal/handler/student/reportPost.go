package student

import (
	"JH_2024_MJJ/internal/service"
	"JH_2024_MJJ/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ReportData struct {
	UserID int64  `json:"user_id" binding:"required"`
	PostID int64  `json:"post_id" binding:"required"`
	Reason string `json:"reason" binding:"required"`
}

func ReportPost(c *gin.Context) {
	var data ReportData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		utils.Log.Printf("举报失败，参数错误")
		return
	}
	//	获取文章信息
	_, err = service.GetArticleByID(data.PostID)
	if err != nil {
		utils.JsonFail(c, 200501, "文章不存在")
		utils.Log.Printf("获取被举报文章[%d]信息失败，因为%s", data.PostID, err.Error())
		return
	}
	////	 校验UserID
	//if data.UserID != reportedArticle.Author {
	//	utils.JsonFail(c, 200501, "用户校验失败")
	//	utils.Log.Printf("用户校验失败")
	//	return
	//}
	err = service.ReportArticle(data.PostID, data.Reason, data.UserID)
	if err != nil {
		utils.JsonFail(c, 200501, "举报失败")
		utils.Log.Printf("举报文章[%d]失败，因为%s", data.PostID, err.Error())
		return
	}
	utils.JsonSuccess(c, nil)
	utils.Log.Printf("举报成功，加速审核中~")
	return
}
