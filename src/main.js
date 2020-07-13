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
  if(localStorage.token && to.path === '/jayne'){
    return next('/category')
  }
  if(!localStorage.token){
    if (to.path === '/jayne') {
      return next('/login')
    }else if(to.path === '/article'){
      return next()
    }else if(to.path !== '/' && to.path !== '/login'){
      return next('/')
    }
  }
   next()
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
