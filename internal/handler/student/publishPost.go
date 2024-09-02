package student

import (
	"JH_2024_MJJ/internal/service"
	"JH_2024_MJJ/pkg/utils"
	"github.com/gin-gonic/gin"
)

type PublishData struct {
	Content string `json:"content" binding:"required"`
	Author  int64  `json:"user_id" binding:"required"`
}

func PublishPost(c *gin.Context) {
	var data PublishData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		utils.Log.Printf("发布失败，参数错误")
		return
	}
	if err := service.PublishPost(data.Content, data.Author); err != nil {
		utils.JsonFail(c, 200502, "发布失败，"+err.Error())
		utils.Log.Printf("发布失败，" + err.Error())
		return
	}
	utils.Log.Printf("%d发布成功，噢耶！", data.Author)
	utils.JsonSuccess(c, nil)
}
