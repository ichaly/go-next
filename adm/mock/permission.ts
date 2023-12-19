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
            icon: 'i-icon-park-outline:system',
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
            icon: 'i-icon-park-outline:personal-privacy',
            name: '/personal',
            title: '个人主页',
            weight: 3,
            type: 'menu',
            hidden: true
          },
          {
            id: 3,
            pid: 0,
            icon: 'i-icon-park-outline:setting-two',
            name: '/system',
            title: '系统设置',
            weight: 2,
            type: 'menu'
          },
          {
            id: 4,
            pid: 3,
            icon: 'i-icon-park-outline:permissions',
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
            icon: 'i-icon-park-outline:audit',
            weight: 2,
            type: 'menu'
          },
          {
            id: 6,
            pid: 3,
            icon: 'i-icon-park-outline:data-user',
            name: '/system/user',
            title: '用户管理',
            weight: 3,
            type: 'menu'
          },
          {
            id: 7,
            pid: 6,
            icon: 'i-icon-park-outline:edit-one',
            name: '/system/user/edit',
            title: '编辑用户',
            weight: 4,
            type: 'menu'
          },
          {
            id: 8,
            pid: 6,
            icon: 'i-icon-park-outline:view-list',
            name: '/system/user/list',
            title: '用户列表',
            weight: 4,
            type: 'menu'
          },
          {
            id: 9,
            pid: 6,
            icon: 'i-icon-park-outline:doc-detail',
            name: '/system/user/:id',
            title: '用户详情',
            weight: 4,
            type: 'menu',
            hidden: true
          },
          {
            id: 10,
            pid: 4,
            icon: 'i-icon-park-outline:view-list',
            name: '/system/permission/list',
            title: '权限列表',
            weight: 4,
            type: 'menu'
          },
          {
            id: 11,
            pid: 4,
            icon: 'i-icon-park-outline:edit-one',
            name: '/system/permission/edit',
            title: '编辑权限',
            weight: 4,
            type: 'menu'
          },
          {
            id: 12,
            pid: 4,
            icon: 'i-icon-park-outline:doc-detail',
            name: '/system/permission/:id',
            title: '权限详情',
            weight: 4,
            type: 'menu',
            hidden: true
          },
          {
            id: 13,
            pid: 5,
            icon: 'i-icon-park-outline:view-list',
            name: '/system/role/list',
            title: '角色列表',
            weight: 4,
            type: 'menu'
          },
          {
            id: 14,
            pid: 5,
            icon: 'i-icon-park-outline:edit-one',
            name: '/system/role/edit',
            title: '编辑角色',
            weight: 4,
            type: 'menu'
          },
          {
            id: 15,
            pid: 5,
            icon: 'i-icon-park-outline:doc-detail',
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
