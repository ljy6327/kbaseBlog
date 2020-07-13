<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>

export default {
  name: 'app',
  components: {
  }
}
</script>
<style>
@import './css/global.css'; /*引入公共样式*/
#app {
  color: #52CC76;
  height: 100%;
}
body,html {
  margin: 0px;
  min-width: 1200px;
  height: 100%;
}
.el-row {
  height: 100%;
}
.el-col {
  height: 100%;
}
ul,li{ 
  padding:0;
  margin:0;
  list-style:none
}
a{ 
  text-decoration:none; 
  color: #52CC76;
} 
button{
  border: 0;
  background-color: transparent;
  outline: none;
}
</style>
