package request

type AddArticleForm struct {
	Title      string     `form:"title" json:"title" binding:"require"`
	CoverImage CoverImage `json:"coverImage" form:"coverImage"`
	Desc       string     `json:"desc" form:"desc"`
	Content    string     `json:"content" form:"content"`
	Tags       []string   `json:"tags" form:"tags"`
}

type PublishArticleForm struct {
	ArticleId uint `json:"articleId,string" valid:"require"`
}

type DeleteArticleForm struct {
	ArticleId uint `json:"articleId,string" valid:"require"`
}

type ListArticlesForm struct {
	Offset int `json:"offset,string" valid:"require"`
	Size   int `json:"size,string" valid:"require"`
}

type SaveArticleForm struct {
	Title      string     `json:"title" valid:"require"`
	CoverImage CoverImage `json:"cover_image" valid:"require"`
	Desc       string     `json:"desc" valid:"require"`
	Content    string     `json:"content" valid:"require"`
	Tags       []string   `json:"tags" valid:"require"`
}
