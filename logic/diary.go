package logic

import (
	"YiYuEcho/dao/mysql"
	"YiYuEcho/models"
	"github.com/google/uuid"
)

// CrateDiary 创建日记
func CreateDiary(diary *models.Diary) (err error) {
	//生成日记id
	dairyID := uuid.New().ID()
	diary.DiaryID = dairyID

	//创建日记 保存到数据库
	if err := mysql.CreateDiary(diary); err != nil {
		//todo: 日志
		//zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}
	//todo ： redis储存日记信息
	return
}

// GetPostByID 根据id查询详情
func GetDiaryByID(diaryID int64) (data *models.ApiDiaryDetail, err error) {
	diary, err := mysql.GetDiaryByID(diaryID)
	if err != nil {

		//todo: 日志
		//zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return nil, err
	}

	//根据作者id查询作者信息
	user, err := mysql.GetUserById(diary.AuthorID)
	if err != nil {

		//todo: 日志
		//zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return nil, err

	}
	//接口数据拼接
	data = &models.ApiDiaryDetail{Diary: diary, AuthorPhoneNumber: user.Phone}
	return data, nil
}

// GetDiaryList 获取日记列表
func GetDiaryList(page, size int64) ([]*models.ApiDiaryDetail, error) {
	diaryList, err := mysql.GetDiaryList(page, size)
	if err != nil {

		//todo: 日志
		//zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return nil, err
	}
	data := make([]*models.ApiDiaryDetail, 0, len(diaryList))

	for _, diary := range diaryList {
		//根据作者id查询作者信息
		user, err := mysql.GetUserById(diary.AuthorID)
		if err != nil {
			//todo: 日志
			//zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
			continue
		}
		//接口数据拼接
		postDetail := &models.ApiDiaryDetail{
			Diary:             diary,
			AuthorPhoneNumber: user.Phone,
		}
		data = append(data, postDetail)
	}
	return data, nil
}

//todo
//PostSearch 搜索业务
