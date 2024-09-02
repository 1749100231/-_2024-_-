package student

import (
	"JH_2024_MJJ/internal/model"
	"JH_2024_MJJ/internal/service"
	"JH_2024_MJJ/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DelArticleData struct {
	UserID    int64 `form:"user_id" binding:"required"`
	ArticleID int64 `form:"post_id" binding:"required"`
}

func DelPost(c *gin.Context) {
	var data DelArticleData
	if err := c.ShouldBind(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		utils.Log.Printf("发布失败，参数错误")
		return
	}

	fmt.Println(data)
	var article *model.Article
	var getArticleError error
	if article, getArticleError = service.GetArticleByID(data.ArticleID); getArticleError != nil {
		utils.JsonFail(c, 200502, "获取帖子失败，"+getArticleError.Error())
		utils.Log.Printf("获取帖子失败，" + getArticleError.Error())
		return
	}

	author := article.Author
	if author != data.UserID {
		utils.JsonFail(c, 200502, "删除帖子失败，这是你的帖子吗，你就删？")
		utils.Log.Printf("删除帖子失败，权限不足")
		return
	}

	if err := service.DelArticle(data.ArticleID); err != nil {
		utils.JsonFail(c, 200501, "删除失败")
		utils.Log.Printf("删除帖子失败,因为%s", err.Error())
		return
	}
	utils.Log.Printf("%d删除成功，噢耶！", data.ArticleID)
	utils.JsonSuccess(c, nil)
}
