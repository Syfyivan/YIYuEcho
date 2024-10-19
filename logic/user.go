package logic

import (
	"YiYuEcho/dao/mysql"
	"YiYuEcho/jwt"
	"YiYuEcho/models"
	_ "errors"
	"math/rand"
	"time"
)

// SignUp 注册业务逻辑
func SignUp(p *models.RegisterForm) (error error) {
	//1.判断用户是否存在
	err := mysql.CheckUserExist(int64(p.UserID)) //err表示CheckUserExist函数的返回值
	if err == nil {
		return err //err表示SignUp函数的返回值
	}

	//2.创建用户(生成用户id todo : 用雪花算法）
	////userID, err := mysql.CreateUser(p)
	//userID, err := snowflake.GetID()
	userID := time.Now().UnixNano() + rand.Int63n(1000)
	//if err != nil {
	//	return errors.New("创建用户ID失败")
	//}
	//构造一个user实例
	u := models.User{
		UserID:   userID,
		Phone:    p.Phone,
		Password: p.Password,
	}
	//3.保存到数据库
	return mysql.InsertUser(&u)
}

// Login 登录业务逻辑
func Login(p *models.LoginForm) (user *models.User, err error) {
	user = &models.User{
		Phone:    p.Phone,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {

		return nil, err

	}
	//生成jwt
	accessToken, refreshToken, err := jwt.GenerateToken(user.UserID, user.Phone)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	return
}
