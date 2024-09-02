package service

import (
	"JH_2024_MJJ/internal/global"
	"JH_2024_MJJ/internal/model"
	"JH_2024_MJJ/pkg/utils"
	"fmt"
)

func GetUserByUserName(username string) (*model.User, error) {
	return d.GetUserByUserName(ctx, username)
}
func GetUserByID(UserID int64) (*model.User, error) {
	return d.GetUserByID(ctx, UserID)
}
func Register(Username string, NickName string, Password string, UserType int, AdminPwd string) error {
	/*
		检查管理员密码
	*/
	if UserType == global.Config.GetInt("userType.admin") &&
		AdminPwd != global.Config.GetString("admin.registerPwd") {
		return &adminPasswordIncorrectError{}
	}

	// 检查账号密码格式
	// 1. 账号只能数字
	for _, char := range Username {
		if char < '0' || char > '9' {
			utils.Log.Printf("%s注册请求不合法\n", Username)
			return &invalidUsernameOrPasswordError{}
		}
	}

	// 2. 密码(md5)32位，数字加小写字母
	if len(Password) != 32 {
		utils.Log.Printf("%s注册请求不合法\n", Username)
		return &invalidUsernameOrPasswordError{}
	}
	for _, char := range Password {
		if !(char >= '0' && char <= '9' || char >= 'a' && char <= 'z') {
			utils.Log.Printf("%s注册请求不合法\n", Username)
			return &invalidUsernameOrPasswordError{}
		}
	}

	// 3. 检查usertype
	if UserType != 1 && UserType != 2 {
		utils.Log.Printf("%s注册请求不合法\n", Username)
		return &invalidUserTypeError{}
	}
	utils.Log.Printf("%s注册成功，噢耶！\n", Username)

	return d.CreateUser(ctx, &model.User{
		Username: Username,
		Nickname: NickName,
		Password: Password,
		UserType: UserType,
	})
}

// 更新Token，返回Token和错误
func UpdateToken(userID int64, userType int, updateTime int64) (string, error) {
	originStr := fmt.Sprintf("%d%d%d", userID, userType, updateTime)
	//	token加密算法，md5凑合一下，嘿嘿~
	md5Str := utils.MD5(originStr)
	newToken := model.TokenTable{
		UserID:     userID,
		UserType:   userType,
		UpdateTime: updateTime,
		Token:      md5Str,
	}
	if _, err := d.GetTokenByID(ctx, userID); err != nil {
		//	创建token
		return md5Str, d.CreateToken(ctx, &newToken)
	} else {
		//	更新token
		return md5Str, d.UpdateToken(ctx, &newToken)
	}
}

func IsLogin(token string) bool {
	_, err := d.GetToken(ctx, token)
	return err == nil
}

func IsAdmin(token string) bool {
	tokenTable, err := d.GetToken(ctx, token)
	if err == nil && tokenTable.UserType == global.Config.GetInt("userType.Admin") {
		return true
	} else {
		return false
	}
}
