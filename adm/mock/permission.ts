import type { MockMethod } from 'vite-plugin-mock'

export type Permission = {
  id: number
  key: string
  pid?: number
  title: string
  icon?: string
  weight?: number
  hidden?: boolean
  external?: boolean
  type: 'menu' | 'action'
}

export type Response<T> = {
  code: number
  data: T
}

export default [
  {
    url: '/api/permission/list',
    method: 'get',
    response: () => {
      return {
        code: 200,
        data: [
          {
            id: 1,
            key: 'dashboard',
            pid: 0,
            title: '仪表盘',
            icon: 'dashboard',
            weight: 1,
            hidden: false,
            external: false,
            type: 'menu'
          },
          {
            id: 2,
            key: 'permission',
            pid: 0,
            title: '系统设置',
            icon: 'lock',
            weight: 2,
            type: 'menu'
          },
          {
            id: 3,
            key: 'permission-list',
            pid: 2,
            title: '权限管理',
            icon: 'lock',
            weight: 1,
            type: 'menu'
          },
          {
            id: 4,
            key: 'role-list',
            pid: 2,
            title: '角色管理',
            icon: 'lock',
            weight: 2,
            type: 'menu'
          }
        ]
      }
    }
  }
] as MockMethod[]
