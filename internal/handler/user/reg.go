package user

import (
	"JH_2024_MJJ/internal/service"
	"JH_2024_MJJ/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type RegisterData struct {
	Username string `json:"username" binding:"required"`  //用户名
	NickName string `json:"name" binding:"required"`      //昵称
	Password string `json:"password" binding:"required"`  //密码
	UserType int    `json:"user_type" binding:"required"` //用户类型 1 学生 2 管理员
	AdminPwd string `json:"type_pwd" binding:""`          //类型密码，学生注册则留空
}

func Register(c *gin.Context) {
	var data RegisterData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误，前端在干什么！乱传参数！")
		return
	}

	_, err := service.GetUserByUserName(data.Username)
	if err == nil {
		log.Printf("%s注册失败，因为%s\n", data.Username, err.Error())
		utils.JsonFail(c, 200503, "该用户名已注册，换个更好听的名字吧~")
		return
	}
	err = service.Register(data.Username, data.NickName, data.Password, data.UserType, data.AdminPwd)
	if err != nil {
		log.Printf("%s注册失败，因为%s\n", data.Username, err.Error())
		utils.JsonFail(c, 200504, "靠！注册失败，为什么！... 哦！原来是因为"+err.Error())
		return
	}
	utils.JsonSuccess(c, nil)

}
