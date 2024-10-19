package mysql

import (
	"YiYuEcho/models"
	"errors"
	"gorm.io/gorm"
)

// CreateTag 创建标签
func CreateTag(tag *models.Tag) (err error) {
	result := db.Create(tag)

	if result.Error != nil {

		return result.Error
	}

	return nil
}

// DeleteTag 删除标签
func DeleteTag(id int64) (err error) {
	result := db.Where("id = ?", id).Delete(&models.Tag{})

	if result.Error != nil {
		return result.Error
	}

	return nil

}

// UpdateTag 更新标签
func UpdateTag(tag *models.Tag) (err error) {
	result := db.Save(tag)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetTagByID 根据ID获取标签
func GetTagByID(id int64) (tag *models.Tag, err error) {
	tag = new(models.Tag)
	result := db.First(tag, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("标签不存在")

		}
		return nil, result.Error

	}

	return
}

// GetTagList 获取标签列表
func GetTagList(page, size int) (tags []models.Tag, err error) {
	offset := (page - 1) * size
	result := db.Order("created_at desc").Offset(offset).Limit(size).Find(&tags)
	if result.Error != nil {

		return nil, result.Error
	}
	return
}

// GetTagListByKeywords 根据关键词获取标签列表
func GetTagListByKeywords(p *models.ParamTagList) (tags []models.Tag, err error) {
	offset := (p.Page - 1) * p.Size
	result := db.Where("name like ?", "%"+p.Search+"%").Order("created_at desc").Offset(offset).Limit(p.Size).Find(&tags)
	if result.Error != nil {

		return nil, result.Error
	}
	return
}
