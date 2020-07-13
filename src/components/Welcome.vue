<template>
    <div id="welcome">
      <div class="atrile">
        <ul>
          <li  v-for="item in articleList" v-bind:key="item.Id">
            <router-link :to="{path:'/article',query:{articleId:item.Id}}">
              <span class="title">{{ item.Title }}</span>
            </router-link>
          </li>
        </ul>
      </div>
    </div>
</template>
<script>
import { apiGetArticles } from '@/request/api';
export default {
   data() {
      return {
        articleList:[{Id:'1',Title:'测试'}]
      }
   },
   created(){
    // this.init();
   },
   methods:{
     init(){
        apiGetArticles({
            AuthorId: "6"
         }).then(res=>{
            this.articleList = res.Data
         })
     }
   }
}
</script>
<style>
#welcome{
  width: 100%;
  height: 100%;
  background-image: url(../assets/background3.jpg);
  background-size: 120%;
  background-position-x: -60%;
  overflow:scroll;
  overflow-x:hidden;
}
#welcome .atrile::-webkit-scrollbar{
  width: 0px;
}
#welcome .atrile{
  width: 650px;
  height: 100%;
  margin:0 auto;
  color: #000;
  overflow-x:hidden;
  overflow-y:auto;
}
#welcome .atrile ul li{
  text-align: left;
  padding: 40px 60px 0px 60px;
}
#welcome .atrile .title{
  font-size: 25px;
  font-weight: 800;
  color: #fff;
  font-family: KaiTi;
}
#welcome .atrile .title:hover{
  color: #52CC76;
}
#welcome .atrile .content{
  padding-left: 5px;
  font-size: 20px;
  color: rgb(80, 78, 78);
  white-space: nowrap;
}


</style>
