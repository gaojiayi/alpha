const path = require("path")
// 获取绝对路径函数
const resolve = dir => {
  return path.join(__dirname, dir)
}
const BASE_URL = process.env.NODE_ENV === 'production' ? '/' : '/'
module.exports = {
  base_url: BASE_URL,
  lintOnSave: true,
  // 打包时不生成.map文件
  productionSourceMap: false,
  // 这里写你调用接口的基础路径，来解决跨域，如果设置了代理，那你本地开发环境的axios的baseUrl要写为 '' ，即空字符串
  // devServer: {
  //   proxy: 'localhost:3000'
  // }
  chainWebPack:config => {
config.require.alias.set('@', resolve('src')).set('_c', resolve('src/components'))
  }
}


