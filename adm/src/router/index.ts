import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'
import Layout from '@/components/Layout.vue'
import type { App } from 'vue'

const pages = import.meta.glob('/src/views/modules/**/*.vue')

export const views: Record<string, RawRouteComponent> = {}

// 索引组件名称和组件加载路径
Object.keys(pages).map((path) => {
  let name = path.match(/\/src\/views\/modules\/(.*)\.vue$/)?.[1]

  //正则匹配中括号中的文字并使用冒号开头的方式替换
  name = name
    ?.replace(/\[[^\]]+\]/g, (match) => {
      return ':' + match.slice(1, -1)
    })
    .toLowerCase()

  if (name) {
    views[`/${name}`] = pages[path]
  }
})

const callbacks: Function[] = []

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/:pathMatch(.*)*',
      redirect: '/404'
    },
    {
      path: '/404',
      component: () => import('@/views/NotFound.vue')
    },
    {
      path: '/login',
      component: () => import('@/views/Login.vue')
    },
    {
      path: '/',
      name: 'root',
      redirect: (to) => {
        const { menus } = useRootStore()
        if (!menus.length) {
          return '/login'
        }
        const firstMenu = (menus: Item[]): Item => {
          const menu = menus[0]
          if (menu?.children?.length) {
            return firstMenu(menu.children)
          } else {
            return menu
          }
        }
        const { name } = firstMenu(menus)
        return name
      },
      component: Layout,
      meta: {
        name: '/',
        title: '首页'
      },
      children: [
        {
          path: '/403',
          component: () => import('@/views/Forbidden.vue')
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

export function setupRouter(app: App) {
  app.use(router)
}

export function addRouter(route: RouteRecordRaw) {
  callbacks.push(router.addRoute('root', route))
}

export function getRouters() {
  return router.getRoutes()
}
