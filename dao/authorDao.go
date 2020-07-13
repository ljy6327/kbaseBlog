/**
 * @Description //作者持久层
 * @Author jayneLiu
 * @Date 2019/7/20 13:12
 */
package dao

import (
	"github.com/jinzhu/gorm"
	"kbase/common/datebase"
	"kbase/model/pojo"
)

func Registered(author pojo.Author) (result *gorm.DB) {
	db := datebase.GetInstance().GetMysqlDB()
	result = db.Create(&author)
	return
}

func FindAuthorByEmail(email string) (result pojo.Author) {
	db := datebase.GetInstance().GetMysqlDB()
	db.Find(&result, "email = ?", email)
	return
}
