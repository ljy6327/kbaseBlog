package service

import (
	"kbase/common/response"
	"kbase/dao"
	"kbase/model/dto"
	"kbase/model/pojo"
)

func SaveCategory(category pojo.Category) response.Response {
	if category.Author_id == 0 || category.Name == "" {
		return response.IsFalse().SetMsg("参数不合法")
	}

	res := dao.SaveCategory(category)
	return res
}

func DeleteCategory(id uint, authorId uint) response.Response {
	if id == 0 {
		return response.IsFalse().SetMsg("参数不合法")
	}

	res := dao.DeleteCategory(id, authorId)
	return res
}

func GetCategories(param dto.SearchCategory) response.Response {
	res := dao.GetCategories(param)
	return res
}
