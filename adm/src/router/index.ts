import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'
import Layout from '@/components/Layout/index.vue'
import Login from '@/views/Login.vue'
import Forbidden from '@/views/Forbidden.vue'
import NotFound from '@/views/NotFound.vue'
import Home from '@/views/Home.vue'
import type { App } from 'vue'

const views = import.meta.glob('/src/views/modules/**/*.vue')

export const pages: typeof views = {}

const callbacks: Function [] = []

// 索引组件名称和组件加载路径
Object.keys(views).map((path) => {
  const name = path.match(/\/src\/views\/modules\/(.*)\.vue$/)?.[1].toLowerCase()
  if (name) {
    pages[name] = views[path]
  }
  // name && router?.addRoute('root', { name, path: `/${name}`, component: views[path] })
})

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/:pathMatch(.*)*',
      redirect: '/404'
    },
    {
      path: '/login',
      component: Login
    },
    {
      path: '/',
      name: 'root',
      component: Layout,
      children: [
        {
          path: '/403',
          component: Forbidden
        },
        {
          path: '/404',
          component: NotFound
        },
        {
          path: '',
          name: 'home',
          component: Home,
          meta: {
            icon: '',
            title: '首页',
            weight: 100
          }
        }
      ]
    }
  ]
})

// 设置页面标题
router.beforeEach((to, from, next) => {
  useTitle(to.meta.title)
  next()
})

export function resetRouter() {
  while (callbacks.length) {
    callbacks.pop()?.()
  }
}

export function addRouter(route: RouteRecordRaw) {
  callbacks.push(router.addRoute('root', route))
}

export function setupRouter(app: App) {
  app.use(router)
}