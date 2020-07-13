package dto

type SearchArticles struct {
	SearchModel
	AuthorId   uint
	CategoryId uint
}

type SearchArticle struct {
	SearchModel
	ArticleId uint
}
