package controller

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
///**
// * @Description 封装响应
// */
//
//type ResponseData struct {
//	Code    int         `json:"code"`
//	Message string      `json:"message"`
//	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
//}
//
//// ResponseError
//func ResponseError(c *gin.Context, code int, message string) {
//	c.JSON(http.StatusOK, ResponseData{Code: code, Message: message})
//}
