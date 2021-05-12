// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import VueRouter from 'vue-router'
import Router from './router/index'
// import ViewUI from 'view-design'
import 'view-design/dist/styles/iview.css'
import store from './store';
import 'ant-design-vue/dist/antd.css';
import { Upload, Modal, Icon, card, Comment, Avatar, Button, Form, Input, List, Tooltip ,Spin} from 'ant-design-vue'
Vue.component(Upload.name, Upload)
Vue.component(Modal.name, Modal)
Vue.component(Icon.name, Icon)
Vue.component(Icon.name, Icon)
Vue.component(card.name, card)
Vue.component(card.Meta.name, card.Meta)
Vue.component(Form.Item.name, Form.Item)
Vue.component(Comment.name, Comment)
Vue.component(Avatar.name, Avatar)
Vue.component(Button.name, Button)
Vue.component(Input.TextArea.name, Input.TextArea)
Vue.component(List.name, List)
Vue.component(Tooltip.name, Tooltip)
Vue.component(Spin.name, Spin)
// eslint-disable-next-line standard/object-curly-even-spacing
// import { Layout, Footer, Content, MenuItem} from 'view-design'
// Vue.component('Layout', Layout)
// Vue.component('Footer', Footer)
// Vue.component('Content', Content)
// Vue.component('MenuItem', MenuItem)
Vue.config.productionTip = false

// Vue.use(ViewUI)
// 路由配置
Vue.use(VueRouter)
// const RouterConfig = {
//   routes: Routers
// }
// const router = new VueRouter(RouterConfig)
/* eslint-disable no-new */
// new Vue({
//   el: '#app',
//   router,
//   components: { App },
//   template: '<App/>'
// })

new Vue({
  el: '#app',
  router: Router,
  store,
  render: h => h(App)
})
