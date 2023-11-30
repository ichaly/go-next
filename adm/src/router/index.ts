import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/Login.vue'
import Layout from '@/components/Layout/index.vue'
import Home from '@/views/Home.vue'
import NotFound from '@/views/NotFound.vue'
import Forbidden from '@/views/Forbidden.vue'

const rootRoutePage404: RouteRecordRaw[] = [
  {
    path: '/404',
    component: () => NotFound
  },
  {
    path: '/:pathMatch(.*)',
    redirect: '/404'
  }
  // {
  //   path: '/:pathMatch(index/.*)',
  //   redirect: '/index/404'
  // }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      name: 'login',
      path: '/login',
      component: Login,
      meta: {}
    },
    {
      path: '/',
      component: Layout,
      redirect: '/home',
      children: [
        {
          name: 'home',
          path: '/home',
          component: Home,
          meta: {
            title: 'title',
            icon: 'i-ep:setting'
          }
        },
        {
          name: '401',
          path: '/401',
          component: Forbidden,
          meta: {
            title: 'title',
            icon: 'i-ep:setting'
          }
        }
      ]
    },
    ...rootRoutePage404
  ]
})

export default router
