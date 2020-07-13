package controller

import (
	"github.com/gin-gonic/gin"
	"kbase/author/service"
	"kbase/model/pojo"
)

func Registered(c *gin.Context) {
	var author pojo.Author
	_ = c.ShouldBind(&author)
	result := service.Registered(author)
	c.JSON(200, result)
}

func Login(c *gin.Context) {
	var author pojo.Author
	_ = c.ShouldBind(&author)
	result := service.Login(author.Email, author.Password)
	c.JSON(200, result)
}
