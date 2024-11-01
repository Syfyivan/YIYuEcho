package controller

import (
	"github.com/gin-gonic/gin"
)

/**
* @Description 封装响应
 */

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

// ResponseError 封装错误响应
func ResponseError(c *gin.Context, code int) {
	rd := &ResponseData{
		Code:    code,
		Message: "error",
		Data:    nil,
	}
	c.JSON(200, rd)
}

// ResponseErrorWithMsg 封装错误响应
func ResponseErrorWithMsg(c *gin.Context, code int, msg string) {
	rd := &ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
	c.JSON(200, rd)
}

// ResponseSuccess 封装成功响应
func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	c.JSON(200, rd)
}
