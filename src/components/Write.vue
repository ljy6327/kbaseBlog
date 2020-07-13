<template>
    <div>
        <div class="font18">
            <el-input v-model="title" placeholder="请输入题目"></el-input>
            <el-button type="primary" @click="save">保存</el-button>
        </div>
        <quill-editor v-model="content"
            ref="myQuillEditor"
            :options="editorOption"
            >
        </quill-editor>
    </div>
</template>

<script>
  import { apiSaveArticle, apiGetArticle } from '@/request/api'
  export default {
    data () {
      return {
        title:'',
        content: '',
        editorOption: {}
        
      }
    },
    created(){
        this.getArticle()
    },
    methods: {
      getArticle(){
        apiGetArticle({
            ArticleId: this.$route.query.articleId
        }).then(res=>{
            this.title = res.Data.Title
            this.content = res.Data.Text
        })
      },  
      save(){
        apiSaveArticle({
          category: parseInt(this.$route.query.categoryId),
          ID: parseInt(this.$route.query.articleId),
          title: this.title,
          text: this.content
        }).then(res =>{
          this.$message.success(res.Msg);
        })
      }
    },
    computed: {
      editor() {
        return this.$refs.myQuillEditor.quill
      }
    }
  }
</script>