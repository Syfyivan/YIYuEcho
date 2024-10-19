package middlewares

import (
	"YiYuEcho/controller"
	"YiYuEcho/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWTAuthMiddleware 基于jwt的认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {

			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请求头中缺少Authorization字段"})
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请求头中Authorization字段格式有误"})
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们将其发送给ExtractClaims方法进行验证
		mc, err := jwt.ParseToken(parts[1])

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "无效的Token"})

			c.Abort()

			return
		}
		//将当前请求的userid信息保存在请求的上下文上
		c.Set(controller.ContextUserIDKey, mc.UserID)

		c.Next()

	}

}
