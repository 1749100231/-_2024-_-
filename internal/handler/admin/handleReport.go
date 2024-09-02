package admin

import (
	"JH_2024_MJJ/internal/model"
	"JH_2024_MJJ/internal/service"
	"JH_2024_MJJ/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handleReportData struct {
	UserID   int64 `json:"user_id" binding:"required"`
	PostID   int64 `json:"post_id" binding:"required"`
	Approval int   `json:"approval" binding:"required"`
}

func HandleReport(c *gin.Context) {
	var data handleReportData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		utils.Log.Printf("处理失败，参数错误")
		return
	}
	var article *model.Article
	var err error
	if article, err = service.GetArticleByID(data.PostID); err != nil {
		// 文章不存在
		utils.JsonFail(c, 200502, "文章不存在")
		utils.Log.Printf("%d用户不存在", data.PostID)
		return
	} else if article.Author != data.UserID {
		// 用户校验失败

		utils.JsonFail(c, 200502, "用户校验失败")
		utils.Log.Printf("%d用户校验失败", data.UserID)
		return
	} else if article.Status != 0 {
		// 没有被举报
		utils.JsonFail(c, 200502, "处理失败，帖子无需处理")
		utils.Log.Printf("%d处理失败，帖子无需处理", data.PostID)
		return
	}
	//article.Status = data.Approval
	if data.Approval == 1 {
		if err = service.DeleteReport(article); err != nil {
			// 异常
			utils.JsonFail(c, 200502, "处理失败，删除帖子失败")
			utils.Log.Printf("帖子[%d]举报处理失败，删除帖子失败", data.PostID)
			return
		}
	} else if data.Approval == 2 {
		article.Status = data.Approval
		if err = service.HandleReport(article); err != nil {
			// 异常
			utils.JsonFail(c, 200502, "处理失败，更新帖子状态失败")
			utils.Log.Printf("帖子[%d]举报处理失败，更新帖子状态失败", data.PostID)
			return
		}
	}

	utils.Log.Printf("%d举报处理成功，辛苦啦~", data.UserID)
	utils.JsonSuccess(c, nil)
}
