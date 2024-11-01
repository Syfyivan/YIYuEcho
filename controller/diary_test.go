package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//type MyLogger struct {
//}
//
//func (m MyLogger) Println(v ...interface{}) {
//	fmt.Println(v)
//}

func TestCreateDiaryHandler(t *testing.T) {
	gin.SetMode(gin.TestMode) //设置Gin模式为测试模式
	r := gin.Default()
	url := "api/v1/diary"

	r.POST("/diary", CreateDiaryHandler) //注册路由

	//r.GET("/log", func(c *gin.Context) {
	//	log.Println("This is a log message")
	//	c.String(200, "Log message printed")
	//})
	//
	////todo
	//// 将日志输出到文件
	//f, err := os.Open("app.log")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//r.Logger.SetOutput(f)

	//r.Run(":8080")

	body := `{
		"content": "这是一个测试日记",
		"title": "测试日记",
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", url, strings.NewReader(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// 判断响应的内容是不是按预期返回了需要登录的错误
	// 1.方法一：判断响应的内容是不是包含指定的字符串
	//assert.Equal(t, w.Body.String(), "需要登录")

	// 2.方法二：将响应的内容反序列化到ResponseData 然后判断字段与预期是否一致
	res := new(ResponseData)
	if err := json.Unmarshal([]byte(w.Body.String()), res); err != nil {

		t.Errorf("Failed to unmarshal response: %v", err)

		//assert.Equal(t, 401, res.Code)
		//assert.Equal(t, res.Message, "需要登录")
	}
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, res.Message, "创建成功")

}
