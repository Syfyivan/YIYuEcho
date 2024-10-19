package controller

import (
	"YiYuEcho/logic"
	"YiYuEcho/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 创建日记
// @Description 创建一篇新的日记
// @Tags diary
// @Accept json
// @Produce json
// @Param diary body models.Diary true "日记内容"
// @Success 200 {object} models.Diary "创建成功的日记"
// @Failure 400 {object} string "请求参数有误"
// @Failure 401 {object} string "获取用户ID失败"
// @Failure 500 {object} string "创建日记失败"
// @Router /diaries [post]

// CreateDiaryHandler 创建日记
func CreateDiaryHandler(c *gin.Context) {
	//获取参数并校验参数
	var diary models.Diary
	if err := c.ShouldBindJSON(&diary); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"status":  0,
			"message": "请求参数有误",
			"error":   err.Error()})

		return
	}
	// 获取作者id，当前请求的UserID(从c取到当前发请求的用户ID)
	userID, err := getCurrentUserID(c)
	if err != nil {
		//todo 日志
		//zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "status": 0, "message": "获取用户ID失败", "error": err.Error()})
		return
	}
	diary.AuthorID = userID
	//创建日记
	if err := logic.CreateDiary(&diary); err != nil {
		//todo
		//zap.zap.L().Error("logic.CreatePost failed", zap.Error(err))

		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "status": 0, "message": "创建日记失败", "error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "status": 1, "message": "创建日记成功", "data": diary})
}

// @Summary 获取日记列表
// @Description 分页获取日记列表
// @Tags diary
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} []models.Diary "日记列表"
// @Failure 500 {object} string "获取日记列表失败"
// @Router /diaries [get]

// DiaryListHandler 分页获取日记列表
func DiaryListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	// 获取日记列表
	data, err := logic.GetDiaryList(page, size) //排序在mysql.GetDiaryList奥
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "status": 0, "message": "获取日记列表失败", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "status": 1, "message": "获取日记列表成功", "data": data})

}

// @Summary 获取日记详情
// @Description 根据日记ID获取日记详情
// @Tags diary
// @Accept json
// @Produce json
// @Param id path int true "日记ID"
// @Success 200 {object} models.Diary "日记详情"
// @Failure 400 {object} string "请求参数有误"
// @Failure 500 {object} string "获取日记详情失败"
// @Router /diaries/{id} [get]

// DiaryDetailHandler 获取日记详情
func DiaryDetailHandler(c *gin.Context) {
	// 获取日记id(从url中获取)
	diaryIDStr := c.Param("id")
	diaryID, err := strconv.ParseInt(diaryIDStr, 10, 64) //将一个字符串类型的变量postIdStr转换为int64，并将结果存储在变量postId中。如果转换过程中出现错误，错误信息将被存储在变量err中
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "status": 0, "message": "请求参数有误", "error": err.Error()})
		return
	}
	// 根据id从数据库中查询日记数据
	diary, err := logic.GetDiaryByID(diaryID)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "status": 0, "message": "获取日记详情失败", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "status": 1, "message": "获取日记详情成功", "data": diary})
}

//todo : 搜索业务
