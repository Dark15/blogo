package article_service

import (
	"blogo/global"
	"blogo/models"
	"gorm.io/gorm"
)

//AddArticle add a single article
func AddArticle(data map[string]interface{}) error {
	article := models.Article{
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		ContentPath:   data["content_path"].(string),
		State:         data["state"].(uint),
		CoverImageURL: data["cover_image_url"].(string),
		Tags:          data["tags"].([]models.Tag),
	}
	if err := global.GORM.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

// DeleteArticle delete a single article
func DeleteArticle(id uint) error {
	if err := global.GORM.Where("id = ?", id).Delete(&(models.Article{})).Error; err != nil {
		return err
	}
	return nil
}

// EditArticle modify a single article
func EditArticle(id uint, data interface{}) error {
	if err := global.GORM.Model(&(models.Article{})).Where(" id  = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

// GetArticle Get a single article based on ID
func GetArticle(id uint) (*models.Article, error) {
	var article models.Article
	err := global.GORM.Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	err = global.GORM.Model(&article).Association("Tags").Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// GetArticles gets a list of articles based on paging constraints
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*models.Article, error) {
	var articles []*models.Article
	err := global.GORM.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetArticleTotal(maps interface{}) (int64, error) {
	var count int64
	if err := global.GORM.Model(&models.Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// ExistArticleByID checks if an article exists based on ID
func ExistArticleByID(id uint) (bool, error) {
	var article models.Article
	err := global.GORM.Select("id").Where("id =?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if article.ID > 0 {
		return true, nil
	}
	return false, nil
}
