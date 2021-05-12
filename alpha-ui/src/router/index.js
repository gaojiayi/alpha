
import Vue from 'vue'
import VueRouter from 'vue-router'
import ViewUI from 'view-design'

Vue.use(ViewUI)

const routers = [
  {
    // 后期可以展示一个大的LOGO
    path: '/',
    meta: {
      title: ''
    },
    component: (resolve) => require(['@/components/HelloWorld.vue'], resolve)
  },
  {
    path: '/img_identify',
    meta: {
      title: ''
    },
    component: (resolve) => require(['@/views/image/img_identify.vue'], resolve)
  },
  {
    path: '/img_detect',
    meta: {
      title: ''
    },
    component: (resolve) => require(['@/views/image/img_detect.vue'], resolve)
  },
  {
    path: '/image_api',
    meta: {
      title: ''
    },
    component: (resolve) => require(['@/views/image/api.vue'], resolve)
  },
  { // 这里面的component作为children的父组件，内部需要添加router-view
    path: '/about', component: (resolve) => require(['@/views/about/index.vue'], resolve),
    children: [
      {
        path: 'comments',
        meta: {
          title: ''
        },
        component: (resolve) => require(['@/views/about/message.vue'], resolve)
      },
      {
        path: 'community',
        meta: {
          title: ''
        },
        component: (resolve) => require(['@/views/about/dev_community.vue'], resolve)
      }

    ]
  },
  { // 这里面的component作为children的父组件，内部需要添加router-view
    path: '/download', component: (resolve) => require(['@/views/download/index.vue'], resolve),
    children: [
      {
        path: 'resource',
        meta: {
          title: ''
        },
        component: (resolve) => require(['@/views/download/resource.vue'], resolve)
      }
    ]
  }
]
const RouterConfig = {
  routes: routers
}
const router = new VueRouter(RouterConfig)
export default router
