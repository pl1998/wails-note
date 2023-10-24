// api.js
import axios from 'axios';
import qs from 'qs';
import { ElMessage } from 'element-plus'
 
// 创建axios实例
const instance = axios.create({
  baseURL: 'http://im.pltrue.top', // 替换为你的后端接口地址
  timeout: 10000, // 请求超时时间
});
 
// 请求拦截器
instance.interceptors.request.use(
  config => {
    // 在请求发送之前可以做一些处理，比如添加请求头等
    return config;
  },
  error => {
    console.log(error)
    // 请求错误处理
    ElMessage({
        message: '客户端异常',
        type: 'error',
      })
    return Promise.reject(error);
  }
);
 
// 响应拦截器
instance.interceptors.response.use(
  response => {
    // 在这里可以对响应数据进行处理

    return response.data;
  },
  error => {
    // 响应错误处理
    return Promise.reject(error);
  }
);
 
// 封装get请求
export function apiGet(url, params) {
  return instance.get(url, {
    params,
  });
}
 
// 封装post请求
export function apiPost(url, data) {
  return instance.post(url, qs.stringify(data));
}

// 封装delete请求
export function apiDel(url, id) {
    return instance.delete(url+"/"+id);
}

// 封装put请求
export function apiPut(url, id,data) {
    return instance.put(url+"/"+id,qs.stringify(data));
}


