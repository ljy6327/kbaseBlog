package category

import (
	"github.com/gin-gonic/gin"
	"kbase/category/controller"
	"kbase/common/request"
)

func Routes(route *gin.Engine) {

	category := route
	// 添加，编辑分类
	category.POST("/category", request.Validate(), controller.SaveCategory)
	// 删除分类
	category.DELETE("/category", request.Validate(), controller.DeleteCategory)
	// 分类列表
	category.POST("/categories", request.Validate(), controller.GetCategories)

}
