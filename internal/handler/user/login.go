package user

import (
	"JH_2024_MJJ/internal/service"
	"JH_2024_MJJ/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var data LoginData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		utils.Log.Printf("登录失败，参数错误")
		return
	}

	user, err := service.GetUserByUserName(data.Username)
	if err != nil {
		utils.JsonFail(c, 200502, "用户不存在")
		utils.Log.Printf("%s登录失败，用户不存在", data.Username)
		return
	}

	if user.Password != data.Password {
		utils.JsonFail(c, 200503, "密码错误")
		utils.Log.Printf("%s登录失败，密码错误", data.Username)
		return
	}
	token, errInUpdateToken := service.UpdateToken(user.ID, user.UserType, time.Now().Unix())
	if errInUpdateToken != nil {
		utils.JsonFail(c, 200503, "登录异常")
		utils.Log.Printf("更新token失败")
		return
	}
	utils.Log.Printf("%s登录登录成功，噢耶！", data.Username)
	utils.JsonSuccess(c, gin.H{"token": token})
}
