/**
 * @Description //文章持久层
 * @Author jayneLiu
 * @Date 2019/7/20 13:14
 */
package dao

import (
	"kbase/common/datebase"
	"kbase/common/response"
	"kbase/model/dto"
	. "kbase/model/pojo"
	"kbase/model/vo"
)

func SaveArticle(article Article) response.Response {
	db := datebase.GetInstance().GetMysqlDB()
	var articleId = article.ID
	result := db.Save(&article)

	if result.Error != nil {
		if articleId > 0 {
			return response.IsFalse().SetMsg("数据不存在")
		}
		return response.IsFalse().SetMsg(result.Error.Error())
	}
	if articleId > 0 {
		return response.IsFalse().SetMsg("修改成功")
	}
	return response.IsSuccess().SetMsg("添加成功")
}

func DeleteArticle(id uint, authorId uint) response.Response {
	db := datebase.GetInstance().GetMysqlDB()
	result := db.Delete(Article{}, "id = ? and author_id = ?", id, authorId)

	if result.Error != nil {
		return response.IsFalse().SetMsg(result.Error.Error())
	}
	return response.IsSuccess().SetMsg("删除成功")
}

func GetArticles(param dto.SearchArticles) response.Response {
	db := datebase.GetInstance().GetMysqlDB()

	count := 0 //总条数
	articles := make([]vo.Article, 20)

	query := db.Table("article").Where("author_id = ? and category_id = ?", param.AuthorId, param.CategoryId)

	countQuery := query
	// 分页 .Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize)
	query = query.Select("article.id, article.title, article.text").Scan(&articles)

	// 数据总条数
	if err := countQuery.Count(&count).Error; err != nil {
		return response.IsFalse()
	}

	return response.IsSuccess().SetData(articles).SetCount(count)
}

func GetArticle(param dto.SearchArticle) response.Response {
	db := datebase.GetInstance().GetMysqlDB()

	article := vo.Article{}

	query := db.Table("article").Where("id = ?", param.ArticleId)

	// 分页 .Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize)
	query = query.Select("article.id, article.title, article.text").Scan(&article)

	return response.IsSuccess().SetData(article)
}
