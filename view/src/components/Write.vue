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
  import { apiSaveArticle } from '@/request/api'
  export default {
    data () {
      return {
        title:'',
        content: '',
        editorOption: {}
        
      }
    },
    // 如果需要手动控制数据同步，父组件需要显式地处理changed事件
    methods: {
      save(){
        apiSaveArticle({
          category: parseInt(this.$route.query.categoryId),
          title: this.title,
          text: this.content
        }).then(res =>{
          console.log(res)
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