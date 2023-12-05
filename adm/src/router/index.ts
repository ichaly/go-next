import {createRouter, createWebHistory} from 'vue-router'
import Layout from '@/components/Layout/index.vue'
import Login from '@/views/Login.vue'
import Forbidden from '@/views/Forbidden.vue'
import NotFound from '@/views/NotFound.vue'
import Home from '@/views/Home.vue'

const pages = import.meta.glob('/src/views/modules/**/*.vue')

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

// 动态添加路由
Object.keys(pages).map((path) => {
    const name = path.match(/\/src\/views\/modules\/(.*)\.vue$/)?.[1].toLowerCase()
    name && router?.addRoute('root', {name, path: `/${name}`, component: pages[path]})
})

export default router
