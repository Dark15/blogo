package tag_service

import (
	"blogo/global"
	"blogo/models"
	"gorm.io/gorm"
)

// AddTag Add a Tag
func AddTag(name string) error {
	tag := models.Tag{
		Name: name,
	}
	if err := global.GORM.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

// DeleteTag delete a tag
func DeleteTag(id uint) error {
	if err := global.GORM.Where("id = ?", id).Delete(&models.Tag{}).Error; err != nil {
		return err
	}
	return nil
}

// ExistTagByID determines whether a Tag exists based on the ID
func ExistTagByID(id uint) (bool, error) {
	var tag models.Tag
	err := global.GORM.Select("id").Where("id = ? ", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistTagByName checks if there is a tag with the same name
func ExistTagByName(name string) (bool, error) {
	var tag models.Tag
	err := global.GORM.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetTags gets a list of tags
func GetTags() ([]models.Tag, error) {
	var tags []models.Tag
	err := global.GORM.Select("*").Find(&tags).Error
	if err != nil {
		return tags, err
	} else {
		return nil, err
	}
}
