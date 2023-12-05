import { defineStore } from 'pinia'
import type { Menu } from '@/components/Layout/index.vue'
import type { Permission } from '@/apis/permission'
import { getPermission } from '@/apis/permission'

function formatMenu(items: Permission[]) {
  let temp: Record<number, Menu> = {}
  let tree: Menu[] = []
  for (const i of items) {
    if (i.hidden || i.type !== 'menu') {
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

  getPermission().then((res) => {
    const items = res.data ?? []
    menus.value = formatMenu(items)
    console.log(menus.value)
  })

  return {
    menus,
    isCollapse,
    toggleCollapse,
    isFullscreen,
    toggleFullscreen
  }
})
