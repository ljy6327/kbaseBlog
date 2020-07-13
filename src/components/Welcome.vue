<template>
    <div id="welcome">
      <div class="atrile">
        <el-timeline>
          <el-timeline-item timestamp="2018/4/12" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/3" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
           <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
           <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2018/4/2" placement="top">
            <el-card>
              <h4>更新 Github 模板</h4>
            </el-card>
          </el-timeline-item>
        </el-timeline>
        <!-- <ul>
          <li  v-for="item in articleList" v-bind:key="item.Id">
            <router-link :to="{path:'/article',query:{articleId:item.Id}}">
              <span class="title">{{ item.Title }}</span>
            </router-link>
          </li>
        </ul> -->
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
.el-timeline{
  margin: 2px;
}
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
  text-align: left;
  width: 670px;
  height: 92%;
  margin: 30px 60px;
  overflow-x:hidden;
  overflow-y:auto;
}

</style>
