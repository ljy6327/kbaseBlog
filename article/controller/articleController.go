package controller

import (
	"github.com/gin-gonic/gin"
	"kbase/article/service"
	"kbase/model/dto"
	"kbase/model/pojo"
)

func SaveArticle(c *gin.Context) {
	// 绑定数据
	var article pojo.Article
	_ = c.ShouldBind(&article)

	getAuthor, ok := c.Get("author")
	if !ok {
		c.JSON(401, "请先登陆")
	}
	article.Author_id = getAuthor.(dto.Author).Id

	result := service.SaveArticle(article)
	c.JSON(200, result)
}

/**
分类列表
*/
func GetArticles(c *gin.Context) {
	// 绑定数据
	var searchArticles dto.SearchArticles
	_ = c.ShouldBind(&searchArticles)

	result := service.GetArticles(searchArticles)
	c.JSON(200, result)
}

/**
文章
*/
func GetArticle(c *gin.Context) {
	// 绑定数据
	var searchArticle dto.SearchArticle
	_ = c.ShouldBind(&searchArticle)

	result := service.GetArticle(searchArticle)
	c.JSON(200, result)
}
