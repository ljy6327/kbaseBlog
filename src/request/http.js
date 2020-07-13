
// 在http.js中引入axios
import axios from 'axios'; // 引入axios

// http request 拦截器
axios.interceptors.request.use(
    config => {
      if (localStorage.token) { //判断token是否存在
        config.headers.Token = localStorage.token;  //将token设置成请求头
      }
      return config;
    },
    err => {
      return Promise.reject(err);
    }
  );

export function get(url, params){    
    return new Promise((resolve, reject) =>{        
        axios.get(url, {            
            params: params        
        }).then(res => {
            resolve(res.data);
        }).catch(err =>{
            reject(err.data)        
        })    
    });
}

export function post(url, params) {
    return new Promise((resolve, reject) => {
         axios.post(url, params)
        .then(res => {
            resolve(res.data);
        })
        .catch(err =>{
            reject(err.data)
        })
    });
}

export function del(url, params) {
    return new Promise((resolve, reject) => {
         axios.delete(url, {            
            params: params        
        }).then(res => {
            resolve(res.data);
        }).catch(err =>{
            reject(err.data)        
        })    
    });
}

