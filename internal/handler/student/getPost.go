package student

import (
	"JH_2024_MJJ/internal/service"
	"JH_2024_MJJ/pkg/utils"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	articleList, err := service.GetArticleList()
	if err != nil {
		utils.JsonFail(c, 200501, "获取失败")
		utils.Log.Printf("获取帖子失败,因为%s", err.Error())
		return
	}

	utils.Log.Printf("获取帖子成功，噢耶！")
	utils.JsonSuccess(c, articleList)
}
