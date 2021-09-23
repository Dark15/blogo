package response

import "blogo/models"

type ListArticleResponse struct {
	Articles   []*models.Article `json:"articles,string"`
	TotalCount int               `json:"total_count,sting"`
}
