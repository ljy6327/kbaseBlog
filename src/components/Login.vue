<template>
  <div>
    <el-input v-model="email" placeholder="请输入邮箱"></el-input>
    <el-input v-model="password" placeholder="请输入密码"></el-input>
    <el-button type="primary" @click="login">登录</el-button>
  </div>
</template>

<script>
  import { apiLogin } from '@/request/api';// 导入我们的api接口
  export default {
    data () {
      return {
        email:'',
        password:''
      }
    },
    methods: {
      login(){
         apiLogin({                    
          email: this.email,                    
          password: this.password               
        }).then(res => {
          if(res.Success){      
            localStorage.setItem('token',res.Data.token); 
            localStorage.setItem('authorId',res.Data.authorId); 
            this.$router.push({ path: '/category'})
          }         
        })    
      }
    }
  }
</script>

<style>

</style>