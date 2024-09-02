package student

import (
	"JH_2024_MJJ/internal/service"
	"JH_2024_MJJ/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UpdataData struct {
	UserID  int64  `json:"user_id" binding:"required"`
	PostID  int64  `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func UpdatePost(c *gin.Context) {
	var data UpdataData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		utils.Log.Printf("修改失败，参数错误")
		return
	}
	//	获取文章信息
	updatedArticle, err := service.GetArticleByID(data.PostID)
	if err != nil {
		utils.JsonFail(c, 200501, "获取文章信息失败")
		utils.Log.Printf("获取文章[%d]信息失败，因为%s", data.PostID, err.Error())
		return
	}
	//	 校验UserID
	if data.UserID != updatedArticle.Author {
		utils.JsonFail(c, 200501, "用户校验失败")
		utils.Log.Printf("用户校验失败")
		return
	}
	err = service.UpdateArticle(data.PostID, data.Content)
	if err != nil {
		utils.JsonFail(c, 200501, "更新失败")
		utils.Log.Printf("更新文章[%d]失败，因为%s", data.PostID, err.Error())
		return
	}
	utils.JsonSuccess(c, nil)
	utils.Log.Printf("更新成功")
	return
}
