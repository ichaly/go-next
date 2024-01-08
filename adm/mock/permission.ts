import type { MockMethod } from 'vite-plugin-mock'

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
            pid: 0,
            icon: 'i-ri:apps-line',
            name: '/dashboard',
            title: '仪表盘',
            weight: 1,
            hidden: false,
            external: false,
            type: 'menu',
            default: true
          },
          {
            id: 2,
            pid: 0,
            icon: 'i-ri:file-user-line',
            name: '/personal',
            title: '个人主页',
            weight: 3,
            type: 'menu',
            hidden: true
          },
          {
            id: 3,
            pid: 0,
            icon: 'i-ri:settings-4-line',
            name: '/system',
            title: '系统设置',
            weight: 2,
            type: 'menu'
          },
          {
            id: 4,
            pid: 3,
            icon: 'i-ri:shield-keyhole-line',
            name: '/system/permission',
            title: '权限管理',
            weight: 1,
            type: 'menu'
          },
          {
            id: 5,
            name: '/system/role',
            pid: 3,
            title: '角色管理',
            icon: 'i-ri:t-shirt-2-line',
            weight: 2,
            type: 'menu'
          },
          {
            id: 6,
            pid: 3,
            icon: 'i-ri:contacts-book-3-line',
            name: '/system/user',
            title: '用户管理',
            weight: 3,
            type: 'menu'
          },
          {
            id: 7,
            pid: 6,
            icon: 'i-ri:file-edit-line',
            name: '/system/user/edit',
            title: '编辑用户',
            weight: 4,
            type: 'menu'
          },
          {
            id: 8,
            pid: 6,
            icon: 'i-ri:file-list-line',
            name: '/system/user/list',
            title: '用户列表',
            weight: 4,
            type: 'menu'
          },
          {
            id: 9,
            pid: 6,
            icon: 'i-ri:file-info-line',
            name: '/system/user/:id',
            title: '用户详情',
            weight: 4,
            type: 'menu',
            hidden: true
          },
          {
            id: 10,
            pid: 4,
            icon: 'i-ri:file-list-line',
            name: '/system/permission/list',
            title: '权限列表',
            weight: 4,
            type: 'menu'
          },
          {
            id: 11,
            pid: 4,
            icon: 'i-ri:file-edit-line',
            name: '/system/permission/edit',
            title: '编辑权限',
            weight: 4,
            type: 'menu'
          },
          {
            id: 12,
            pid: 4,
            icon: 'i-ri:file-info-line',
            name: '/system/permission/:id',
            title: '权限详情',
            weight: 4,
            type: 'menu',
            hidden: true
          },
          {
            id: 13,
            pid: 5,
            icon: 'i-ri:file-list-line',
            name: '/system/role/list',
            title: '角色列表',
            weight: 4,
            type: 'menu'
          },
          {
            id: 14,
            pid: 5,
            icon: 'i-ri:file-edit-line',
            name: '/system/role/edit',
            title: '编辑角色',
            weight: 4,
            type: 'menu'
          },
          {
            id: 15,
            pid: 5,
            icon: 'i-ri:file-info-line',
            name: '/system/role/:id',
            title: '角色详情',
            weight: 4,
            type: 'menu',
            hidden: true
          }
        ]
      }
    }
  }
] as MockMethod[]
