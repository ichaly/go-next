import { defineStore } from 'pinia'
import { getPermission } from '@/apis/permission'
import { addRouter, resetRouter, views } from '@/router'

function formatMenu(items: Item[]) {
  const dict: Record<number, Partial<Item>> = {}
  const tree: Item[] = []
  for (const i of items) {
    i.name = i.name.toLowerCase()
    //如果不是路由，或者没有对应的页面，则跳过
    if (i.type !== 'menu' || !views[i.name]) {
      continue
    }
    //添加动态路由到名称为root到根路由下保证每个页面都会使用Layout组件装饰
    addRouter({
      path: `${i.name}`,
      component: views[i.name],
      meta: {
        icon: i.icon,
        title: i.title,
        items: []
      }
    })
    //如果隐藏则不添加到菜单
    if (i.hidden) {
      continue
    }
    dict[i.id] = { ...i, children: dict[i.id]?.children ?? [] }
    const temp: Item = <Item>dict[i.id]
    if (temp.pid === 0) {
      tree.push(temp)
    } else {
      if (!dict[i.pid]) {
        dict[i.pid] = { children: [] }
      }
      dict[i.pid].children?.push(temp)
    }
  }
  return tree
}

export const useRootStore = defineStore('root', () => {
  const [isCollapse, toggleCollapse] = useToggle(false)
  const { isFullscreen, toggle: toggleFullscreen } = useFullscreen()
  const menus: Ref<Item[]> = ref([])

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
