package routers

import (
	"YiYuEcho/controller"
	"YiYuEcho/middlewares"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	"log"
	"os"

	_ "YiYuEcho/docs"
)

// SetUpRouter 设置路由
func SetUpRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //设置成发布模式
	}
	//if mode == gin.DebugMode {
	//	gin.SetMode(gin.DebugMode) //设置成调试模式(用于开发阶段，输出详细的日志信息，便于调试)
	//}

	// 创建一个日志文件
	file, err := os.Create("gin.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 创建一个日志记录器
	logger := log.New(file, "[GIN]", log.LstdFlags)
	r := gin.Default()
	r.Use(gin.LoggerWithWriter(logger.Writer()))

	//todo

	//注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //swaggerFiles.Handler是Swagger UI的文件处理程序，用于提供Swagger UI所需的静态文件

	v1 := r.Group("/api/v1")
	//注册登录业务
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/signup", controller.SignUpHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)
	//日记业务
	v1.GET("/diary", controller.DiaryListHandler)
	v1.POST("/diary", controller.CreateDiaryHandler)
	v1.GET("/diary/:id", controller.DiaryDetailHandler)

	//中间件
	v1.Use(middlewares.JWTAuthMiddleware()) //应用jwt认证中间件
	{
		//todo:panic: handlers are already registered for path '/api/v1/diary'
		//v1.POST("/diary", controller.CreateDiaryHandler) //创建日记
		v1.GET("/ping", func(c *gin.Context) {

			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	pprof.Register(r) //注册pprof路由
	r.NoRoute(func(c *gin.Context) {

		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "请求的路径不匹配任何已注册的路由",
		})
	})
	return r
}
