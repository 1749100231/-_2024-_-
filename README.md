# 精弘2024大作业

## 1、MGG的小b话

这里是MGG，七日速成Go，写的可能存在这那的问题，也可能存在前后风格不一致的情况，这也正好反应了本人对Go以及后端架构理解的发展过程，同时也让本开发蒟蒻和狮山代码Maker意识到了项目架构和代码规范的重要意义，对于本次大作业，也是我第一个Go的项目，如果哪里写的不好请随时来~~喷我~~指导我，靴靴。

## 2、接口介绍可以看<a href="https://github.com/1749100231/JH_2024BigHomework/tree/main/ApiDoc" title="API文档">API文档</a>，并导入APIFox以便于测试

## 3、简述鉴权的实现

### 我针对鉴权的实现主要思路是：

#### 1）在登录的时候后端生成一个Token返回到前端,并将token以及UserType、UserID、UpdateTime保存到表TokenTable中，用于后续的校验。

```go
func Login(c *gin.Context) {
	//...
	token, errInUpdateToken := service.UpdateToken(user.ID, user.UserType, time.Now().Unix())
	if errInUpdateToken != nil {
		utils.JsonFail(c, 200503, "登录异常")
		utils.Log.Printf("更新token失败")
		return
	}
	utils.Log.Printf("%s登录登录成功，噢耶！", data.Username)
	utils.JsonSuccess(c, gin.H{"token": token})
}

```



```go
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
        //  save 就行，不要管我了
		//	创建token
		return md5Str, d.CreateToken(ctx, &newToken)
	} else {
		//	更新token
		return md5Str, d.UpdateToken(ctx, &newToken)
	}
}
```



#### 3）前端在后续的请求中在请求头中携带这个token，在后端路由组Student接收到这个请求时，会在中间件IsLogin中判断token是否有效。在Admin路由组则在中间件IsAdmin中去判断数据库中token对应用户的权限是否为管理员。

```go
//Service层
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

```

```go
//Midware层
func IsLogin(c *gin.Context) {
	isLogin := service.IsLogin(c.GetHeader("token"))
	if !isLogin {
		utils.JsonResponse(c, 200, 200404, "登录过期", nil)
		c.Abort()
	}
}

func IsAdmin(c *gin.Context) {
	isAdmin := service.IsAdmin(c.GetHeader("token"))
	if !isAdmin {
		utils.JsonResponse(c, 200, 200404, "权限不足", nil)
		c.Abort()
	}
}

```

