package controller

import (
	_ "YiYuEcho/dao/mysql"
	"YiYuEcho/jwt"
	"gorm.io/gorm"
	"net/http"
	"strings"

	"YiYuEcho/logic"
	"YiYuEcho/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

// SignUpHandler 注册业务
// @Summary 注册接口
// @Description 用户注册接口
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param data body models.RegisterForm true "用户注册参数"
// @Success 200 {object} models.User
func SignUpHandler(c *gin.Context) {
	// 1.从前端获取请求参数，用fo接收
	var fo *models.RegisterForm
	//2.检验数据有效性（暂时不用
	//if err := c.ShouldBindJSON(&fo); err != nil {
	//	c.JSON(400, gin.H{"error": err.Error()})
	//	return
	//}
	//…………………
	fmt.Printf("fo: %v\n", fo)
	// 3.调用logic层方法，完成业务处理-注册用户
	if err := logic.SignUp(fo); err != nil {
		//todo: 日志
		//zap.L().Error("logic.SignUp failed", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"status": 0, "code": 404, "mag": "用户不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "code": 500, "mag": "服务器内部错误"})
		return
	}
	//返回响应

	c.JSON(http.StatusOK, gin.H{"status": 1, "code": 200, "mag": "注册成功"})

}

// LoginHandler 登录业务
// @Summary 登录接口
// @Description 用户登录接口
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param data body models.LoginForm true "用户登录参数"
// @Success 200 {object} models.User
func LoginHandler(c *gin.Context) {
	// 1.获取请求参数
	var lo *models.LoginForm
	if err := c.ShouldBindJSON(&lo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "code": 400, "mag": "请求参数错误"})
		return
	}
	// 2.调用logic层方法，完成业务处理-登录
	user, err := logic.Login(lo)
	//todo: 日志
	if err != nil {
		//zap.L().Error("logic.SignUp failed", zap.Error(err))
		if err == gorm.ErrRecordNotFound {

			c.JSON(http.StatusNotFound, gin.H{"status": 0, "code": 404, "mag": "用户不存在"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "code": 500, "mag": "服务器内部错误"})
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "code": 200, "mag": "登录成功", "data": user})
}

// RefreshTokenHandler 刷新token
// @Summary 刷新token接口
// @Description 刷新用户token
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param data body models.RefreshTokenForm true "刷新token参数"
// @Success 200 {object} models.Token
func RefreshTokenHandler(c *gin.Context) {
	refreshToken := c.Query("refresh_token")
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URL
	// 这里假设Token放在Header的 Authorization 中，并使用 Bearer 开头
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {

		c.JSON(http.StatusUnauthorized, gin.H{"status": 0, "code": 401, "mag": "请求未携带token，或者token格式错误"})
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 0, "code": 401, "mag": "请求未携带token，或者token格式错误"})
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], refreshToken)
	//todo: 日志

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"status": 0, "code": 401, "mag": "token刷新失败"})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":        1,
		"code":          200,
		"mag":           "刷新成功",
		"access_token":  aToken,
		"refresh_token": rToken})
}
