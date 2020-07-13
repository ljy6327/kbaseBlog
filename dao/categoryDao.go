/**
 * @Description //分类持久层
 * @Author jayneLiu
 * @Date 2019/7/20 13:12
 */
package dao

import (
	"kbase/common/datebase"
	"kbase/common/response"
	"kbase/model/dto"
	. "kbase/model/pojo"
	"kbase/model/vo"
)

func SaveCategory(category Category) response.Response {
	db := datebase.GetInstance().GetMysqlDB()
	var categoryId = category.ID
	result := db.Save(&category)

	if result.Error != nil {
		if categoryId > 0 {
			return response.IsFalse().SetMsg("数据不存在")
		}
		return response.IsFalse().SetMsg(result.Error.Error())
	}
	if categoryId > 0 {
		return response.IsFalse().SetMsg("修改成功")
	}
	return response.IsSuccess().SetMsg("添加成功")
}

func DeleteCategory(id uint, authorId uint) response.Response {
	db := datebase.GetInstance().GetMysqlDB()
	result := db.Delete(Category{}, "id = ? and author_id = ?", id, authorId)

	if result.Error != nil {
		return response.IsFalse().SetMsg(result.Error.Error())
	}

	return response.IsSuccess().SetMsg("删除成功")
}

func GetCategories(param dto.SearchCategory) response.Response {
	db := datebase.GetInstance().GetMysqlDB()

	count := 0 //总条数
	categories := make([]vo.Category, 20)

	query := db.Table("category").Where("author_id = ?", param.AuthorId)

	countQuery := query
	// 分页
	query = query.Select("category.id, category.name").Scan(&categories)

	// 数据总条数
	if err := countQuery.Count(&count).Error; err != nil {
		return response.IsFalse()
	}

	return response.IsSuccess().SetData(categories).SetCount(count)
}
