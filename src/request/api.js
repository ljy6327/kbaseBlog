import { get, post, del } from './http'

export const apiLogin = p => post('author/login', p)
// 分类
export const apiGetCategory = p => post('categories', p)
export const apiSaveCategory = p => post('category', p)

// 文章
export const apiSaveArticle = p => post('article', p)
export const apiGetArticles = p => get('articles', p)
export const apiGetArticle = p => get('article', p)
export const apiDeleteArticle = p => del('article', p)
