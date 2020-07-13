package datebase

import (
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

/*
 MysqlConnectPool
 数据库连接操作库
 基于gorm封装开发
*/
type MysqlConnectPool struct {
}

var instance *MysqlConnectPool
var once sync.Once

var db *gorm.DB
var err error

// Singleton
func GetInstance() *MysqlConnectPool {
	once.Do(func() {
		instance = &MysqlConnectPool{}
	})
	return instance
}

/*
*  init database connect
 */
func (m *MysqlConnectPool) InitDataPool() (issucc bool) {
	db, err = gorm.Open("mysql", "root:password@/kbase?charset=utf8&parseTime=True&loc=Local")

	// 启用Logger，显示详细日志
	db.LogMode(true)
	// 全局禁用表名复数
	db.SingularTable(true)

	if err != nil {
		log.Fatal(err)
		return false
	}

	// defer close database
	//defer db.Close()
	return true
}

/*
*  external connection object
 */
func (m *MysqlConnectPool) GetMysqlDB() (db_con *gorm.DB) {
	return db
}
