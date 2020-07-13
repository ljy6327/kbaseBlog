package router

import (
	"github.com/gin-gonic/gin"
	"kbase/article"
	"kbase/author"
	"kbase/category"
	"log"
	"net/http"
)

/**
初始化路由
整合使用GIN路由
*/
func Init() {
	router := gin.Default()
	router.Use(CorsMiddleware())

	author.Routes(router)   //作者
	category.Routes(router) //分类
	article.Routes(router)  //文章

	err := router.Run(":8081")
	if err != nil {
		log.Panic("router error")
	}
}

/*
跨域
*/
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
