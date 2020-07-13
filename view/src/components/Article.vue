<template>
    <div id="article">
        <div class="title">{{ article.Title}}</div>
        <div class="content" v-html="article.Text"></div>
    </div>
</template>
<script>
import { apiGetArticle } from '@/request/api';// 导入我们的api接口
export default {
    data() {
        return {
            articleId:'',
            article:{}
        }
    },
    created(){
        this.articleId = this.$route.query.articleId
        this.getArticle()
    },
    watch: {
        $route(){ //监听路由是否变化
            this.articleId = this.$route.query.articleId
            this.getArticle();//重新加载数据
        }
    },
    methods: {
        getArticle(){
            apiGetArticle({
                ArticleId: this.$route.query.articleId
            }).then(res=>{
                this.article = res.Data
            })
        }    
    }
}
</script>
<style>
#article{
    width: 100%;
    height: 100%;
    overflow:scroll;
    overflow-x:hidden;
}
#article .title{
    font-size: 50px;
    color: #000;
    padding: 30px 50px;
}
#article .content{
    font-size: 18px;
    color: #000;
    padding: 0px 50px;
    text-align: left;
}
</style>
