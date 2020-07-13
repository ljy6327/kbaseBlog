package controller

import (
	"github.com/gin-gonic/gin"
	"kbase/category/service"
	"kbase/model/dto"
	"kbase/model/pojo"
	"strconv"
)

/**
添加,编辑分类
*/
func SaveCategory(c *gin.Context) {
	// 绑定数据
	var category pojo.Category
	_ = c.ShouldBind(&category)

	getAuthor, ok := c.Get("author")
	if !ok {
		c.JSON(401, "请先登陆")
	}
	category.Author_id = getAuthor.(dto.Author).Id

	result := service.SaveCategory(category)
	c.JSON(200, result)
}

/**
删除分类
*/
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.ParseUint(c.DefaultQuery("id", "0"), 10, 0)
	getAuthor, ok := c.Get("author")
	if !ok {
		c.JSON(401, "请先登陆")
	}
	authorId := getAuthor.(dto.Author).Id

	result := service.DeleteCategory(uint(id), authorId)
	c.JSON(200, result)
}

/**
分类列表
*/
func GetCategories(c *gin.Context) {
	// 绑定数据
	var searchCategory dto.SearchCategory
	_ = c.ShouldBind(&searchCategory)

	getAuthor, ok := c.Get("author")
	if !ok {
		c.JSON(401, "请先登陆")
	}
	searchCategory.AuthorId = getAuthor.(dto.Author).Id

	result := service.GetCategories(searchCategory)
	c.JSON(200, result)
}
