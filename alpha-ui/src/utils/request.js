// 建一个单独axios实例用来做特殊请求
import axios from 'axios';
// 创建一个token实例
// axios config http://www.axios-js.com/zh-cn/docs/
const service = axios.create({
  baseURL: process.env.baseURL,
  timeout: 5000 // request timeout
});
// TODO 请求拦截器  后续做一些token检验

// TODO 响应拦截器  可以对指定错误码做一些信息提示  以及登出操作
export default service;
