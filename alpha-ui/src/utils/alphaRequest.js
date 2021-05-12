import axios from 'axios';
const qs = require('qs');
/**
 * alphaRequest 统一后端请求接口
 */
let baseUrl = 'http://localhost:9090/api/alpha/'
const env = '__pro__';
if (env === '__dev__') {
  baseUrl = 'http://localhost:9090/api/alpha/'
}
// axios.defaults.baseURL = baseUrl
export default {
  // 具体的加密逻辑
  sign(params, token, urlencode) {
    return 'sign(params, token, urlencode)';
  },
  ajax(method, url, params) {
    params = params || {};
    let options = {
      method
    };

    if (method === 'GET') {
      options.url = url + '?' + qs.stringify(params);
      options.headers = {
        'Content-Type': 'application/x-www-form-urlencoded',
        'Access-Control-Allow-Origin': '*'
      }
    } else {
      options.url = url;

      // options.data = qs.stringify(params);
      options.data = params; // 使用json发送
      options.headers = {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*'
      };
    }
    return axios(options)
  },
  /**
   *  GET 请求
   * @param {string} url 接口地址
   * @param {Object} params 参数
   * @return {Object} Promise
   */
  get(url, params) {
    // 重要：跨域问题解决方案https://yq.aliyun.com/articles/705295
    return this.ajax('GET', url, params)
    // this.ajax('get', baseUrl + url, params).then(res => {
    //   console.log(res)
    //   if (res.errno === 0) {
    //     return res.data;
    //   } else {
    //     throw res;
    //   }
    // }).catch((error) => {
    //   console.log(error)
    // });
  },

  /**
   *  POST 请求
   * @param {string} url 接口地址
   * @param {Object} params 参数
   * @return {Object} Promise
   */
  post(url, params) {
    return this.ajax('POST', url, params);
  }
}
