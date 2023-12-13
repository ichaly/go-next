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
            name: 'dashboard',
            title: '仪表盘',
            weight: 1,
            hidden: false,
            external: false,
            type: 'menu'
          },
          {
            id: 2,
            pid: 0,
            icon: 'i-icon-park-outline:setting-two',
            name: 'setting',
            title: '系统设置',
            weight: 2,
            type: 'menu'
          },
          {
            id: 3,
            pid: 2,
            icon: 'i-icon-park-outline:permissions',
            name: 'permission',
            title: '权限管理',
            weight: 1,
            type: 'menu'
          },
          {
            id: 4,
            name: 'RoleList',
            pid: 2,
            title: '角色管理',
            icon: 'i-icon-park-outline:audit',
            weight: 2,
            type: 'menu',
            hidden: true
          },
          {
            id: 5,
            pid: 2,
            icon: 'i-icon-park-outline:data-user',
            name: 'userlist',
            title: '用户管理',
            weight: 3,
            type: 'menu'
          },
          {
            id: 6,
            pid: 5,
            icon: 'i-icon-park-outline:add-user',
            name: 'adduser',
            title: '添加用户',
            weight: 4,
            type: 'menu'
          },
          {
            id: 7,
            pid: 0,
            icon: 'i-icon-park-outline:personal-privacy',
            name: 'personal',
            title: '个人主页',
            weight: 3,
            type: 'menu'
          },
          {
            id: 8,
            pid: 0,
            icon: 'i-icon-park-outline:table-file',
            name: 'table',
            title: '基础表格',
            weight: 3,
            type: 'menu'
          }
        ]
      }
    }
  }
] as MockMethod[]
