<template>
    <el-row>
        <el-col :span="7">
            <div id="home" ref="home">
                <div class="title">
                    <div class="tab">
                        <ul>
                            <li>
                                <span class="font23">Jayne's blog</span>
                            </li>
                        </ul>
                    </div>
                </div> 
                <div class="flag">
                    <img src="../assets/flag.png"/>
                </div> 
                <div class="comment">
                    <!-- <button class="title font23" @click="openCommentAdd">留言<i class="el-icon-d-arrow-right"></i></button> -->
                    QQ : 632786948<br>
                    E-MAIL : ljy6327@vip.qq.com
                </div>
            </div>
        </el-col>
        <el-col :span="17">
            <router-view></router-view>
        </el-col>
    </el-row>
</template>
<script>
export default {
    data() {
        return {

        }
    },
    mounted(){

    },
    created(){
        
    },
    watch:{
       
    },
    activated() {

    },
    methods: {
        //更多tab
        moreTab(home,tab,comment){
            this.ishome = home;
            this.istab = tab;
            this.iscomment = comment;
        },
        //打开评论详情
        openCommentDetail() {
            this.$alert('这是一段内容', '标题名称', {
            confirmButtonText: '确定'
            });
        },
        //添加评论
        openCommentAdd() {
            this.$prompt('请输入邮箱', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消'
            }).then(({ value }) => {
            this.$message({
                type: 'success',
                message: '你的邮箱是: ' + value
            });
            }).catch(() => {
            this.$message({
                type: 'info',
                message: '取消输入'
            });       
            });
      }
    }
}
</script>
<style scoped>
.el-col {
    text-align: center;
}
#home{
    width: 100%;
    height: 100%;
    overflow:scroll;
    overflow-x:hidden;
}
#home::-webkit-scrollbar{
    background-color: #FFA500;
    border-radius: 10px;
    width: 0px;
}
#home .title:after{
    content: '';
    display: block;
    clear: both;
}
#home .title .tab{
    float: left;
    padding: 20px;
}

#home .title .tab ul li{
    float: left;
}
#home .title .tab ul li a{
    display: block;
    font-weight: 500;
    letter-spacing: 1px;
    padding: 5px 5px 5px;
    color: #000;
}
#home .title .tab ul li a:hover{
    color: #FFA500;
}
#home .title .more{
    padding: 23px;
    padding-left: 0px;
    float: right;
}
#home .title .more i{
    color: 	#B5B5B5;
    cursor: pointer;
}
#home .title .more i:hover{
    color: 	#FFA500;
}
#home .flag{
    padding-top: 50px;
    padding-left:40%;
}
#home .flag img{
    width:35px;
}
#home .comment{
    color: #000;
    margin-top: 50px;
}
#home .comment .title{
    cursor: pointer;
}
#home .moreTab ul li{
    margin-top: 40px;
}
#home .moreTab li a{
    color: 	#000;
}
#home .moreTab li a:hover{
    color: 	#FFA500;
}
</style>
