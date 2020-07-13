<template>
    <div id="welcome">
      <div class="atrile">
        <ul>
          <li  v-for="item in articleList" v-bind:key="item.Id">
            <router-link :to="{path:'/article',query:{articleId:item.Id}}">
              <span class="title">{{ item.Title }}</span>
              <div class="content">{{ item.Text}}</div>
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
        articleList:[]
      }
   },
   created(){
    this.init();
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
::-webkit-scrollbar {
        display: none;
}
#welcome{
  width: 100%;
  height: 100%;
  background-image: url(../assets/background3.jpg);
  overflow:scroll;
  overflow-x:hidden;
}
#welcome::-webkit-scrollbar{
  background-color: #FFA500;
  border-radius: 10px;
  width: 0px;
}
#welcome .atrile{
  width: 786px;
  height: 100%;
  margin:0 auto;
  background: hsla(0,0%,100%,.55) border-box;
  color: #000;
  overflow-x:hidden;
  overflow-y:auto;
}
#welcome .atrile ul li{
  text-align: left;
  padding: 40px 60px 0px 60px;
}
#welcome .atrile .title{
  font-size: 30px;
  font-weight: 800;
  color:black;
}
#welcome .atrile .content{
  width: 600px;
  padding-left: 5px;
  font-size: 20px;
  color: rgb(80, 78, 78);
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}


</style>
