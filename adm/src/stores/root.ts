import { defineStore } from 'pinia'
import { getPermission } from '@/apis/permission'
import { addRouter, pages, resetRouter } from '@/router'

function formatMenu(items: Permission[]) {
  let temp: Record<number, Menu> = {}
  let tree: Menu[] = []
  for (const i of items) {
    i.name = i.name.toLowerCase()
    //如果不是路由，或者没有对应的页面，则跳过
    if (i.type !== 'menu' || !pages[i.name]) {
      continue
    }
    //添加动态路由到名称为root到根路由下保证每个页面都会使用Layout组件装饰
    addRouter({
      path: `/${i.name}`,
      component: pages[i.name],
      meta: {
        title: i.title,
        icon: i.icon
      }
    })
    //如果隐藏则不添加到菜单
    if (i.hidden) {
      continue
    }
    temp[i.id] = {
      ...i,
      children: temp[i.id]?.children ?? []
    }
    const item = temp[i.id]
    if (item.pid === 0) {
      tree.push(item)
    } else {
      if (!temp[i.pid]) {
        temp[i.pid] = {
          name: '',
          children: []
        }
      }
      temp[i.pid].children?.push(item)
    }
  }
  return tree
}

export const useRootStore = defineStore('root', () => {
  const [isCollapse, toggleCollapse] = useToggle(false)
  const { isFullscreen, toggle: toggleFullscreen } = useFullscreen()
  const menus: Ref<Menu[]> = ref([])

  const loadMenus = () => {
    getPermission().then((res) => {
      //移除之前的路由
      resetRouter()
      //更新新菜单
      menus.value = formatMenu(res.data ?? [])
    })
  }

  //先自动加载一次数据
  loadMenus()
  return {
    menus,
    loadMenus,
    isCollapse,
    toggleCollapse,
    isFullscreen,
    toggleFullscreen
  }
})
