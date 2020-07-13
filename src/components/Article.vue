<template>
  <div id="article">
    <el-row v-if="deleteFlag" style="height:30px">
      <el-button v-on:click="editorArticle" type="primary" icon="el-icon-edit" circle></el-button>
      <el-button v-on:click="deleteArticle" type="danger" icon="el-icon-delete" circle></el-button>
    </el-row>
    <div class="title">{{ article.Title}}</div>
    <div class="ql-snow">
      <div class="ql-editor">
        <div class="content" v-html="article.Text"></div>
      </div>
    </div>
  </div>
</template>
<script>
import { apiGetArticle, apiDeleteArticle } from "@/request/api"; // 导入我们的api接口
export default {
  data() {
    return {
      articleId: "",
      categoryId:"",
      article: {},
      deleteFlag: false
    };
  },
  created() {
    this.articleId = this.$route.query.articleId;
    this.getArticle();

    if (localStorage.token) {
      this.deleteFlag = true;
    }
  },
  watch: {
    $route() {
      //监听路由是否变化
      this.articleId = this.$route.query.articleId;
      this.getArticle(); //重新加载数据
    }
  },
  methods: {
    getArticle() {
      apiGetArticle({
        ArticleId: this.articleId 
      }).then(res => {
        this.article = res.Data;
        this.categoryId = res.Data.CategoryId;
      });
    },
    editorArticle() {
      this.$router.push({
        path: "/category/write",
        query: { articleId: this.articleId, categoryId: this.categoryId}
      });
    },
    deleteArticle(){
      apiDeleteArticle({
          id: this.articleId 
      }).then(res => {
          this.$message.success(res.Msg);
      });
    }
  }
};
</script>
<style>
#article {
  width: 100%;
  height: 100%;
  overflow: scroll;
  overflow-x: hidden;
}
#article .title {
  font-size: 50px;
  color: #000;
  padding: 30px 50px;
}
#article .content {
  font-size: 18px;
  color: #000;
  padding: 0px 50px;
  text-align: left;
}
</style>
