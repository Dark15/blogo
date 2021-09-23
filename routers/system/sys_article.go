package system

import (
	"blogo/api"
	"blogo/models/response"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	systemApi := Router.Group("api/system/article")
	{
		systemApi.POST("/listArticles", func(c *gin.Context) {
			api.ApiGroupApp.SystemApiGroup.ListArticles(&response.GinContextE{C: c})
		})
		systemApi.POST("/getArticleDetails")
		systemApi.POST("/searchArticle")
		systemApi.POST("/saveArticle", func(c *gin.Context) {
			api.ApiGroupApp.SystemApiGroup.SaveArticle(&response.GinContextE{C: c})
		})
		systemApi.POST("/publishArticle", func(c *gin.Context) {
			api.ApiGroupApp.SystemApiGroup.PublishArticle(&response.GinContextE{C: c})
		})
		systemApi.POST("/updateArticle")
		systemApi.POST("/deleteArticle", func(c *gin.Context) {
			api.ApiGroupApp.SystemApiGroup.DeleteArticle(&response.GinContextE{C: c})
		})
		systemApi.POST("/recoverArticle")
		systemApi.POST("/listArticleTags")
	}
}
