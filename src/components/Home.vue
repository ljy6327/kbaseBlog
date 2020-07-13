<template>
    <el-row>
        <el-col :span="7">
            <div id="home" ref="home">
                <div class="flag">
                    <img src="../assets/chifan.png"/>
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
    methods: {
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
    background-color: #52CC76;
    border-radius: 10px;
    width: 0px;
}
#home .flag{
    padding-top: 180px;
}
#home .flag img{
    width:90px;
}
#home .comment{
    color: #000;
    margin-top: 80px;
}
</style>
