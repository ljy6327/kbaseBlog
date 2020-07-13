package service

import (
	"kbase/common/response"
	"kbase/dao"
	"kbase/model/dto"
	"kbase/model/pojo"
)

func SaveArticle(article pojo.Article) response.Response {
	if article.Author_id == 0 || article.Category_id == 0 || article.Title == "" || article.Text == "" {
		return response.IsFalse().SetMsg("参数不合法")
	}
	res := dao.SaveArticle(article)
	return res
}

func DeleteArticle(id uint, authorId uint) response.Response {
	if id == 0 {
		return response.IsFalse().SetMsg("参数不合法")
	}

	res := dao.DeleteArticle(id, authorId)
	return res
}

func GetArticles(param dto.SearchArticles) response.Response {
	res := dao.GetArticles(param)
	return res
}

func GetArticle(param dto.SearchArticle) response.Response {
	res := dao.GetArticle(param)
	return res
}
