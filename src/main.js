import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import routers from './router'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// editor
import VueQuillEditor from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
// axios
import axios from 'axios'

Vue.use(axios);
Vue.use(VueQuillEditor, /* { default global options } */)
Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(ElementUI);

const router = new VueRouter({
  mode: 'history',
  routes: routers
})

//路由跳转之前
router.beforeEach((to, from, next) => {
  if(!localStorage.token){
    console.log(to.path === '/jayne')
    if (to.path === '/jayne') {
      return next('/login')
    }else if(to.path !== '/' && to.path !== '/login'){
      return next('/')
    }
  }
   next()
})

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
