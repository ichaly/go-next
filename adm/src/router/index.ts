import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/Login.vue'
import Layout from '@/components/Layout/index.vue'
import Home from '@/views/Home.vue'
import NotFound from '@/views/NotFound.vue'
import Forbidden from '@/views/Forbidden.vue'

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
            title: '首页',
            icon: 'i-ep:home-filled'
          }
        },
        {
          path: '/403',
          component: Forbidden
        },
        {
          path: ':pathMatch(home/.*)',
          redirect: '/403'
        }
      ]
    },
    {
      path: '/404',
      component: () => NotFound
    },
    {
      path: '/:pathMatch(.*)',
      redirect: '/404'
    }
  ]
})

export default router
