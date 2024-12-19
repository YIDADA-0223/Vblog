import { createRouter, createWebHistory } from 'vue-router'
import app from '@/stores/app'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      // 全局唯一名称
      name: 'LoginPage',
      // 组件,没有嵌套就是最终页面,会把这个页面的代码加载过来
      // 惰性加载，使用函数()=> 返回 Component，等使用的时候在加载这个组件
      component: () => import('@/views/login/LoginPage.vue')
    },
    {
      path: '/backend',
      // 全局唯一名称
      name: 'BackendLayout',
      // 组件,没有嵌套就是最终页面
      component: () => import('@/views/backend/BackendLayout.vue'),
      redirect: { name: 'BackendBlogList' },
      beforeEnter: () => {
        // 怎么确认用户当前有没有登录喃?
        // 如果中断直接返回你要去向的页面
        if (!app.value.token) {
          return { name: 'LoginPage' }
        } else {
          // return true
        }
      },

      // redirect: { name: 'BackendLayout' },
      // beforeEnter: () => {
      //   // 怎么确认用户当前有没有登录
      //   // 如果中断直接返回登录页
      //   if (!app.value.token) {
      //     // console.log()
      //     // console.log(app.value)
      //     return { name: 'LoginPage' }
      //   }
      // },
      children: [
        {
          // /backend/vblogs
          path: 'blogs/list',
          name: 'BackendBlogList',
          component: () => import('@/views/backend/blogs/ListPage.vue')
        },
        {
          path: 'blogs/edit',
          name: 'BackendBlogEdit',
          component: () => import('@/views/backend/blogs/EditPage.vue')
        },
        {
          path: 'comments/list',
          name: 'BackendCommentsList',
          component: () => import('@/views/backend/comment/ListPage.vue')
        },
        {
          path: 'tags/list',
          name: 'BackendTagList',
          component: () => import('@/views/backend/tag/ListPage.vue')
        }
      ]
    },
    {
      path: '/frontend',
      // 全局唯一名称
      name: 'FrontendLayout',
      // 组件,没有嵌套就是最终页面
      component: () => import('@/views/frontend/FrontendLayout.vue'),
      redirect: { name: 'FrontendBlogList' },
      children: [
        {
          path: 'blogs/list',
          name: 'FrontendBlogList',
          component: () => import('@/views/frontend/blogs/ListPage.vue')
        },
        {
          path: 'blogs/detail',
          name: 'FrontendBlogDetail',
          component: () => import('@/views/frontend/blogs/DetailPage.vue')
        }
      ]
    },
    // 匹配前面所有没有被名字的路由, 都指向404页面 /*
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('@/views/common/NotFound.vue')
    }
  ]
})

export default router
