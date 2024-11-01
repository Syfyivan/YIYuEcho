package main

import (
	"YiYuEcho/dao/mysql"
	"YiYuEcho/routers"
	"YiYuEcho/settings"
	"fmt"
	_ "gorm.io/gorm/logger"
	"net"
	_ "net"
)

// @title YiYuEcho
// @host localhost:8080
// @BasePath /api/v1

//func initDB() {
//	// TODO: Initialize the database connection
//	var err error
//	//dsn := "root:123456@tcp(127.0.0.1:3306)/yiyuser?charset=utf8mb4&parseTime=True&loc=Local"
//	//db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	db, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/yiyuser?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
//	if err != nil {
//
//		panic("连接数据库失败")
//	}
//	//迁移数据库模式
//	err = db.AutoMigrate(&models.User{})
//	if err != nil {
//		return
//	}
//
//}

// @title YiYuEcho
// @version 1.0
// @description 这是一个基于Gin框架的API接口
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	//todo:加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("加载配置失败，err：%v\n", err)
		return
	}
	if err := mysql.InitDB(settings.Conf.MySQLConfig); err != nil {

		fmt.Printf("初始化数据库失败，err：%v\n", err)
		return
	}

	//GORM会自动管理数据库连接的生命周期，包括连接的创建、重用和关闭

	//todo: redis

	////生成用户id
	////todo: 或者用数据库自增ID？如果你的应用程序需要与数据库交互，你可以使用数据库的自增ID作为唯一ID。在插入新记录时，数据库会自动生成一个唯一的ID

	//注册路由
	port := settings.Conf.Port
	portStr := fmt.Sprintf("%d", port)
	_, err := net.Listen("tcp", ":"+portStr)
	if err != nil {
		fmt.Printf("端口已被占用: %s\n", err)
		return
	}
	fmt.Printf("端口可用: %v\n", port)
	/////////////////////////////
	r := routers.SetUpRouter(settings.Conf.Mode)
	//// 读取YAML文件
	//file, err := os.Open("YiYuEcho/conf/config.yaml")
	////todo
	////todo: 配置文件打不开欸，这个打不开，之前logrus配置文件也打不开
	//if err != nil {
	//	fmt.Println("无法打开配置文件：", err)
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//// 解析YAML文件
	//err = yaml.NewDecoder(file).Decode(&settings.Conf)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("配置信息: %+v\n", settings.Conf)
	//fmt.Println("Mode: %s", settings.Conf.Mode)
	//
	//err2 := r.Run(fmt.Sprintf(":%d\n", settings.Conf.Port))
	//if err2 != nil {
	//
	//	fmt.Printf("启动失败，err：%v\n", err2)
	//	return
	//}
	r.Run(`:8080`)
}

//func main() {
//	initDB()
//	// 检查端口是否可用
//	// 设置端口号
//	port := 8082
//
//	//	// 检查端口是否可用
//	_, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
//	if err != nil {
//		fmt.Printf("端口已被占用： %s", err)
//		return
//	}
//	fmt.Printf("端口可用： %d", port)
//
//	//	// 初始化 Gin 路由
//	r := gin.Default()
//
//	//routers.SetUpRouter(r)	// 注册路由
//	//
//	//	// 启动服务器
//	err = r.Run(fmt.Sprintf(":%d", port))
//	if err != nil {
//		fmt.Printf("启动失败，err：%v", err)
//		return
//	}
//
//	v1 := r.Group("/api/v1")
//	//注册登录业务
//	v1.POST("/login", controller.LoginHandler)
//	v1.POST("/signup", controller.SignUpHandler)
//	v1.GET("/refresh_token", controller.RefreshTokenHandler)
//	//日记业务
//	v1.GET("/diary", controller.DiaryListHandler)
//	v1.POST("/diary", controller.CreateDiaryHandler)
//	v1.GET("/diary/:id", controller.DiaryDetailHandler)
//
//	//	//中间件
//	v1.Use(middlewares.JWTAuthMiddleware()) //应用jwt认证中间件
//	{
//		//todo:panic: handlers are already registered for path '/api/v1/diary'
//		//v1.POST("/diary", controller.CreateDiaryHandler) //创建日记
//		v1.GET("/ping", func(c *gin.Context) {
//
//			c.JSON(200, gin.H{
//				"message": "pong",
//			})
//		})
//	}
//	pprof.Register(r) //注册pprof路由
//	r.NoRoute(func(c *gin.Context) {
//
//		c.JSON(http.StatusNotFound, gin.H{
//			"code":    404,
//			"message": "请求的路径不匹配任何已注册的路由",
//		})
//	})
//	return
//}
