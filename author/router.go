package author

import (
	"github.com/gin-gonic/gin"
	"kbase/author/controller"
)

func Routes(route *gin.Engine) {

	author := route.Group("/author")
	// 注册
	author.POST("/registered", controller.Registered)
	// 登陆
	author.POST("/login", controller.Login)
}
