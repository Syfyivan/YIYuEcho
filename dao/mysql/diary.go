package mysql

import (
	"YiYuEcho/models"
	"errors"
	"gorm.io/gorm"
)

// CreateDiary 创建日记
func CreateDiary(diary *models.Diary) (err error) {
	//sqlStr := `insert into post(diary_id, tag, content, author_id)values(?, ?, ?, ?)`
	//_, err = db.Exec(sqlStr, diary.DiaryID, diary.Tag, diary.Content, diary.AuthorID)
	result := db.Create(diary)
	if result.Error != nil {
		//todo : 日志
		//zap.L().Error("db.Create(post) failed", zap.Error(result.Error))
		err = errors.New("插入数据失败")
		return
	}
	return nil
}

// todo
// GetDiaryByTag 根据Tag查询日记
func GetDiaryByTag(tag string) (diaryList []*models.Diary, err error) {
	return
}

// GetDiaryByID
func GetDiaryByID(diaryID int64) (diary *models.Diary, err error) {
	diary = new(models.Diary)
	result := db.First(diary, diaryID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("日记不存在")

		}

		return nil, result.Error

	}
	return diary, nil
}

// GetDiaryListByKeywords
func GetDiaryListByKeywords(p *models.ParamDiaryList, userID int) (diaryList []*models.Diary, err error) {
	offset := (p.Page - 1) * p.Size
	result := db.Where("user_id = ? AND (title LIKE ? OR content LIKE ?) AND is_deleted = ?",
		userID, "%"+p.Search+"%", "%"+p.Search+"%", false).Order(p.Order).Offset(offset).Limit(p.Size).Find(&diaryList)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}

// GetDiaryList  获取帖子列表
func GetDiaryList(page, size int64) (diaries []*models.Diary, err error) {
	err = db.Table("diary").Limit(int(size)).Offset(int((page - 1) * size)).Order("created_at").Find(&diaries).Error
	return
}

// AddTagToDiary 给日记添加标签
func AddTagToDiary(diaryID int64, tag string) (err error) {
	//result := db.Model(&models.Diary{}).Where("diary_id = ?", diaryID).Association("Tags").Append(&models.Tag{Tag: tag})
	//if result.Error != nil {
	//	return
	//}
	err = db.Model(&models.Diary{}).Where("diary_id = ?", diaryID).Association("Tags").Append(&models.Tag{Tag: tag})
	if err != nil {
		return err
	}
	return nil
}

// RemoveTagFromDiary 从日记中移除标签
func RemoveTagFromDiary(diaryID int64, tag string) (err error) {
	err = db.Model(&models.Diary{}).Where("diary_id = ?", diaryID).Association("Tags").Delete(&models.Tag{Tag: tag})
	if err != nil {
		return err
	}
	return nil
}

// UpdateTagOfDiary 更新日记的标签
func UpdateTagOfDiary(diaryID int64, tag string) (err error) {
	err = db.Model(&models.Diary{}).Where("diary_id = ?", diaryID).Association("Tags").Replace(&models.Tag{Tag: tag})
	if err != nil {
		return err
	}
	return nil
}

// DeleteAllDiariesWithTag 删除所有带有指定标签的日记
func DeleteAllDiariesWithTag(tag string) (err error) {
	err = db.Where("tag = ?", tag).Delete(&models.Diary{}).Error
	if err != nil {
		return err
	}
	return nil
}
