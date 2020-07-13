package article

import (
	"github.com/gin-gonic/gin"
	"kbase/article/controller"
	"kbase/common/request"
)

func Routes(route *gin.Engine) {

	article := route
	// 新增，修改文章
	article.POST("/article", request.Validate(), controller.SaveArticle)
	// 获取文章列表
	article.GET("/articles", controller.GetArticles)
	// 获取文章
	article.GET("/article", controller.GetArticle)
}
