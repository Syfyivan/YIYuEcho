package logger

//func GinLogger() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		start := time.Now()
//
//		path := c.Request.URL.Path
//		query := c.Request.URL.RawQuery
//		c.Next()
//
//		cost := time.Since(start)
//
//		logrus.WithFields(logrus.Fields{
//			"status":     c.Writer.Status(),
//			"method":     c.Request.Method,
//			"path":       path,
//			"query":      query,
//			"ip":         c.ClientIP(),
//			"user-agent": c.Request.UserAgent(),
//			"errors":     c.Errors.ByType(gin.ErrorTypePrivate).String(),
//			"cost":       cost,
//		}).Info("request")
//
//		if c.Writer.Status() >= 500 {
//			logrus.WithFields(logrus.Fields{
//				"status":     c.Writer.Status(),
//				"method":     c.Request.Method,
//				"path":       path,
//				"query":      query,
//				"ip":         c.ClientIP(),
//				"user-agent": c.Request.UserAgent(),
//				"errors":     c.Errors.ByType(gin.ErrorTypePrivate).String(),
//				"cost":       cost,
//			}).Error("request")
//		}
//	}
//}

///////////////////////////////////////////

//var logger *zap.Logger
//
//// getLogWriter 获取日志写入器
//func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
//	lumberJackLogger := &lumberjack.Logger{
//		Filename:   filename,
//		MaxSize:    maxSize,
//		MaxBackups: maxBackup,
//		MaxAge:     maxAge,
//	}
//	return zapcore.AddSync(lumberJackLogger)
//}
//
//// getEncoder 获取编码器
//func getEncoder() zapcore.Encoder {
//	encoderConfig := zap.NewProductionEncoderConfig()
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	encoderConfig.TimeKey = "time"
//	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
//	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
//	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
//	return zapcore.NewJSONEncoder(encoderConfig)
//}
//
//// InitLogger 初始化日志 v
//func InitLogger(cfg *settings.LogConfig, mode string) (err error) {
//	//创建日志写入器
//	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
//	encoder := getEncoder()
//	var level zapcore.Level
//	err = level.UnmarshalText([]byte(cfg.Level))
//	if err != nil {
//		return
//	}
//	var core zapcore.Core
//	if mode == "dev" {
//		// 开发模式，创建一个多输出核心，将日志同时输出到文件和终端。
//		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
//		core = zapcore.NewTee(
//			zapcore.NewCore(encoder, writeSyncer, level),
//			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
//		)
//	} else {
//		//否则，只将日志输出到文件
//		core = zapcore.NewCore(encoder, writeSyncer, level)
//	}
//	//使用zap.New函数创建一个日志记录器（logger），并添加一个调用者信息字段，以便在日志中记录生成日志的代码文件名和行号
//	logger = zap.New(core, zap.AddCaller())
//	//将新创建的日志记录器设置为全局日志记录器，并输出一条初始化成功的日志信息。
//	zap.ReplaceGlobals(logger)
//	zap.L().Info("Logger init success")
//	return
//}
//
//// GinLogger 接收gin框架默认的日志
//func GinLogger() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// 开始时间
//		startTime := time.Now()
//		// 获取请求信息
//		path := c.Request.URL.Path
//		query := c.Request.URL.RawQuery
//		// 处理请求(Next方法会执行后续的中间件和处理器函数)
//		c.Next()
//
//		// 计算耗时
//		cost := time.Since(startTime)
//		logger.Info(path,
//			//记录请求的详细信息:状态码、请求方法、路径、查询参数、客户端IP、User-Agent、错误信息以及处理耗时
//			zap.Int("status", c.Writer.Status()),
//			zap.String("method", c.Request.Method),
//			zap.String("path", path),
//			zap.String("query", query),
//			zap.String("ip", c.ClientIP()),
//			zap.String("user-agent", c.Request.UserAgent()),
//			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
//			zap.Duration("cost", cost),
//		)
//	}
//}
//
//// GinRecovery recover掉项目异常，并使用zap记录相关日志
//func GinRecovery(stack bool) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		defer func() {
//			if err := recover(); err != nil {
//				// 检查错误类型是否为net.OpError，如果是，则检查错误信息是否包含"broken pipe"或"connection reset by peer"，如果是，则将brokenPipe设置为true
//				var brokenPipe bool
//				if ne, ok := err.(*net.OpError); ok {
//					if se, ok := ne.Err.(*os.SyscallError); ok {
//						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
//							brokenPipe = true
//						}
//					}
//				}
//				httpRequest, _ := httputil.DumpRequest(c.Request, false)
//				if brokenPipe {
//					logger.Error(c.Request.URL.Path,
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//					)
//					// If the connection is dead, we can't write a status to it.
//					c.Error(err.(error)) // nolint: errcheck
//					c.Abort()
//					return
//				}
//				if stack {
//					logger.Error("[Recovery from panic]",
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//						zap.String("stack", string(debug.Stack())),
//					)
//				} else {
//					logger.Error("[Recovery from panic]",
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//					)
//				}
//				c.AbortWithStatus(http.StatusInternalServerError)
//			}
//		}()
//		c.Next()
//	}
//
//}
