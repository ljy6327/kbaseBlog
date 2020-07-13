<template>
   <el-row>
      <el-col :span="7">
         <div id="blog">
            <div class="category" ref='category'>
               <div class="add-category" @click="createCategory()">
                  <i class="el-icon-circle-plus-outline"></i>
                  <span>新建分类</span>
               </div>
               <ul>
                  <li v-for="(item, index) in categoryList" v-bind:key="item.Id"  v-on:click="openArticleList(index,item.Id)" v-bind:class="{selected:index==current}"> 
                     <span v-on:dblclick="deleteCategory(item.Id)">{{ item.Name}}</span>
                  </li>
               </ul>
            </div>
            <div class="article" v-if="articleShow">
               <router-link :to="{path:'/category/write',query:{categoryId:categoryId}}">
                  <div class="add-article" > 
                     <div class="article-title">
                        <i class="el-icon-circle-plus-outline"></i>
                        新建文章
                        </div>
                     <div class="article-line"></div>
                  </div>
               </router-link>
               <ul>
                  <li v-for="item in articleList" v-bind:key="item.Id"> 
                     <router-link :to="{path:'/category/article',query:{articleId:item.Id}}">
                        <div class="article-title" v-on:dblclick="editorArticle(item.Id)">{{ item.Title }}</div>
                     </router-link>
                     <div class="article-line"></div>
                  </li>
               </ul>
            </div>
         </div>
      </el-col>
      <el-col :span="17">
         <router-view></router-view>
      </el-col>
   </el-row>
</template>
<script>
import { apiGetCategory, apiSaveCategory, apiGetArticles } from '@/request/api';// 导入我们的api接口
export default {
   data() {
      return {
         current:'0',
         blog:false,
         articleShow:false, //文章列表显隐
         categoryList:[],
         articleList:[],
         categoryId:""
      }
   },
   created(){
      this.getCategories()
   },
   methods: {
      getCategories(){
         apiGetCategory({}).then(res=>{
            this.categoryList = res.Data
         })
      },
      createCategory(){
         this.$prompt('请输入分类名', '添加分类', {
            confirmButtonText: '确定',
            cancelButtonText: '取消'
         }).then(({ value }) => {
            apiSaveCategory({
               name: value
            }).then(res=>{
               if(res.Success){
                  this.$message({
                     type: 'success',
                     message: '添加分类成功'
                  });
                  this.getCategories()
               }
               
            })
         }).catch(() => {});
        
      },
      getArticles(){
         apiGetArticles({
            AuthorId: localStorage.getItem("authorId"),
            CategoryId: this.categoryId
         }).then(res=>{
            this.articleList = res.Data
         })
      },
      openArticleList(index,categoryId){
         // 选择变色
         this.current = index;
         // 设置样式
         this.$refs.category.style.width = '42%'
         this.$refs.category.style.float = 'left'

         this.articleShow = true;
         this.categoryId = categoryId;
         this.getArticles()
      },
   }
}
</script>
<style>
#blog{
    width: 100%;
    height: 100%;
    overflow:scroll;
    overflow-x:hidden;
}
#blog::-webkit-scrollbar{
    background-color: #FFA500;
    border-radius: 10px;
    width: 0px;
}
#blog .category .add-category{
   height: 60px;
   border-radius: 10px;
   cursor: pointer;
   font-size: 22px;
   line-height: 60px;
   padding-left: 20px;
   cursor: pointer;
   color: #ffffff;
   margin: 1px 0px;
   animation: categoryBreath 3s infinite;
}
@keyframes categoryBreath {
  0% { background: #0ABFBC;}
  50% {background: rgb(14, 211, 208);}
  100% {background: #0ABFBC;}
}
#blog .category .add-category:hover{
   background:#FFA500;
}
#blog .category ul li{
   height: 60px;
   border-radius: 10px;
   color: #ffffff;
   font-size: 22px;
   line-height: 60px;
   padding-left: 20px;
   cursor: pointer;
}
#blog .category ul li:nth-child(1){
   background-image: linear-gradient(to right,#ee9ca7 0%, #ffffff 100%);
}
#blog .category ul li:nth-child(10n+1){
   background-image: linear-gradient(to right,#ee9ca7 0%, #ffffff 100%);
}
#blog .category ul li:nth-child(2){
   background-image: linear-gradient(to right,#FBD786 0%, #ffffff 100%);
}
#blog .category ul li:nth-child(10n+2){
   background-image: linear-gradient(to right,#FBD786 0%, #ffffff 100%);
}
#blog .category ul li:nth-child(3){
   background-image: linear-gradient(to right,#2980B9 0%,#6DD5FA 50%,#ffffff  100%);
}
#blog .category ul li:nth-child(10n+3){
   background-image: linear-gradient(to right,#2980B9 0%,#6DD5FA 50%,#ffffff  100%);
}
#blog .category ul li:nth-child(4){
   background-image: linear-gradient(to right,#4AC29A 0%,#ffffff 100%);
}
#blog .category ul li:nth-child(10n+4){
   background-image: linear-gradient(to right,#4AC29A 0%,#ffffff 100%);
}
#blog .category ul li:nth-child(5){
   background-image: linear-gradient(to right,#ADA996 0%,#DBDBDB 50%,#ffffff 100%);
}
#blog .category ul li:nth-child(10n+5){
   background-image: linear-gradient(to right,#ADA996 0%,#DBDBDB 50%,#ffffff 100%);
}
#blog .category ul li:nth-child(6){
   background-image: linear-gradient(to right,#f4c4f3 0%,#ffffff 100%);
}
#blog .category ul li:nth-child(10n+6){
   background-image: linear-gradient(to right,#f4c4f3 0%,#ffffff 100%);
}
#blog .category ul li:nth-child(7){
   background-image: linear-gradient(to right,#f7797d 0%,#ffffff 100%);
}
#blog .category ul li:nth-child(10n+7){
   background-image: linear-gradient(to right,#f7797d 0%,#ffffff 100%);
}
#blog .category ul li:nth-child(8){
   background-image: linear-gradient(to right,#f5af19 0%,#ffffff 100%);
}
#blog .category ul li:nth-child(10n+8){
   background-image: linear-gradient(to right,#f5af19 0%,#ffffff 100%);
}
#blog .category ul li:nth-child(9){
   background-image: linear-gradient(to right,#00c3ff 0%,#ffffff 100%);
}
#blog .category ul li:nth-child(10n+9){
   background-image: linear-gradient(to right,#00c3ff 0%,#ffffff 100%);
}
#blog .category ul li:nth-child(10){
   background-image: linear-gradient(to right,#636FA4 0%,#ffffff 100%);
}
#blog .category ul li:nth-child(10n+10){
   background-image: linear-gradient(to right,#636FA4 0%,#ffffff 100%);
}
#blog .category ul li:hover{
   background: #FFA500;
}
#blog .article ul li{
   height: 50px;
   border-radius: 10px;
   color: #000;
   font-size: 22px;
   line-height: 50px;
}
#blog .article .add-article{
   width: 58%;
   margin-left: 42%;
   margin-top: 1px;
   height:60px;
   border-radius: 10px;
   line-height: 60px;
   cursor: pointer;
   color: #ffffff;
   font-size: 22px;
   animation: categoryArticle 3s infinite;
   text-align: center;
}
@keyframes categoryArticle {
  0% { background: #0ABFBC;}
  50% {background: rgb(14, 211, 208);}
  100% {background: #0ABFBC;}
}
#blog .article ul li .article-title{
   padding: 0px 30px;
   white-space:nowrap;
   overflow:hidden;
   text-overflow : ellipsis;
   cursor: pointer;
   color: #000;
}
#blog .article ul li .article-title:hover{
   color: #FFA500;
}
#blog .article ul li .article-line{
   position: relative;
   left: 47%;
   width: 50%;
   border: 1px dashed #B5B5B5;
}
.selected{
   background:#FFA500 !important
}
</style>
