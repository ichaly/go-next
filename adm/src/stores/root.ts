import { defineStore } from 'pinia'
import { getPermission } from '@/apis/permission'
import { addRouter, resetRouter, views } from '@/router'

function formatMenu(items: Item[]) {
  const dict: Record<number, Partial<Item>> = {}
  const tree: Item[] = []
  for (const item of items) {
    item.name = item.name.toLowerCase()
    //如果不是路由，或者没有对应的页面，则跳过
    if (item.type !== 'menu' || !views[item.name]) {
      continue
    }
    //添加动态路由到名称为root到根路由下保证每个页面都会使用Layout组件装饰
    addRouter({
      path: `${item.name}`,
      component: views[item.name],
      meta: {
        icon: item.icon,
        title: item.title,
        items: []
      }
    })
    //如果隐藏则不添加到菜单
    if (item.hidden) {
      continue
    }
    const { name, ...rest } = item
    dict[item.id] = { name: `/${name}`, ...rest, children: dict[item.id]?.children ?? [] }
    const temp: Item = <Item>dict[item.id]
    if (temp.pid === 0) {
      tree.push(temp)
    } else {
      if (!dict[item.pid]) {
        dict[item.pid] = { children: [] }
      }
      dict[item.pid].children?.push(temp)
    }
  }
  return tree
}

export const useRootStore = defineStore('root', () => {
  const [isCollapse, toggleCollapse] = useToggle(false)
  const { isFullscreen, toggle: toggleFullscreen } = useFullscreen()
  const menus: Ref<Item[]> = ref([])

  const loadMenus = async () => {
    let res = await getPermission()
    //移除之前的路由
    resetRouter()
    //更新新菜单
    menus.value = formatMenu(res.data ?? [])
  }

  return {
    menus,
    loadMenus,
    isCollapse,
    toggleCollapse,
    isFullscreen,
    toggleFullscreen
  }
})
