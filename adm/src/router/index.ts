import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/Login.vue'
import Layout from '@/components/Layout/index.vue'
import Forbidden from '@/views/Forbidden.vue'
import NotFound from '@/views/NotFound.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: <RouteRecordRaw>[
    {
      path: '/:pathMatch(.*)*',
      redirect: '/404'
    },
    {
      name: 'login',
      path: '/login',
      component: Login
    },
    {
      path: '/',
      name: 'index',
      component: Layout,
      children: [
        {
          path: '403',
          name: '403',
          component: Forbidden
        },
        {
          path: '404',
          name: '404',
          component: NotFound
        },
        {
          path: '',
          name: 'home',
          component: Home,
          meta: {
            title: '首页'
          }
        },
        {
          path: 'dashboard',
          name: 'dashboard',
          component: () => import('@/views/Dashboard.vue'),
          meta: {
            title: '仪表盘'
          }
        }
      ]
    }
  ]
})

//路由发生变化修改页面title
router.beforeEach((to, from, next) => {
  useTitle(to.meta.title)
  next()
})

router.addRoute('index', {
  path: 'table',
  name: 'table',
  component: () => import('@/views/Table.vue'),
  meta: {
    title: '表格'
  }
})

export default router
