import { defineStore } from 'pinia'
import { getPermission } from '@/apis/permission'
import { addRouter, resetRouter, views } from '@/router'

function formatMenu(items: Item[]) {
  const dict: Record<number, Partial<Item>> = {}
  const tree: Item[] = []
  //结构化菜单
  for (const item of items) {
    const { id, pid, name, ...rest } = item
    //如果不是菜单，或者不在菜单中显示则跳过
    if (item.type !== 'menu' || item.hidden) {
      continue
    }
    //填充树形结构和索引结构
    dict[id] = { id, pid, name, ...rest, children: [] }
    if (pid === 0) {
      tree.push(<Item>dict[id])
    } else {
      if (!dict[pid]) {
        dict[pid] = { children: [] }
      }
      dict[pid].children?.push(<Item>dict[id])
    }
  }
  //递归查询所有的祖辈节点
  const findParents = (item: Item): Item[] => {
    if (item.pid === 0) {
      return [item]
    } else {
      return findParents(<Item>dict[item.pid]).concat(item)
    }
  }
  //动态添加路由
  for (const item of items) {
    const { name } = item
    //如果不是路由，或者没有对应的页面则跳过
    if (item.type !== 'menu' || !views[name]) {
      continue
    }
    //添加动态路由到名称为root到根路由下
    addRouter({
      path: name,
      component: views[name],
      meta: { items: findParents(item) }
    })
  }
  return tree
}

export const useRootStore = defineStore('root', () => {
  const [isCollapse, toggleCollapse] = useToggle(false)
  const { isFullscreen, toggle: toggleFullscreen } = useFullscreen()
  const menus: Ref<Item[]> = ref([])

  const loadMenus = async () => {
    const res = await getPermission()
    //移除之前的动态路由
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
