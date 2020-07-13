/**
 * @Description //基础查询条件
 * @Author jayneLiu
 * @Date 2019/7/21 0:21
 */
package dto

type SearchModel struct {
	Page     int
	PageSize int
	Keyword  string
}
