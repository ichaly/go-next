import { defineStore } from 'pinia'

export type Permission = {
  title: string
  view: string
  kind: 'menu' | 'action' | 'path'
  icon?: string
  hidden?: boolean
  external?: boolean
  weight?: number
  parent?: Permission
  children?: Permission[]
}

export const useRootStore = defineStore('root', () => {
  const [isCollapse, toggleCollapse] = useToggle(false)
  const { isFullscreen, toggle: toggleFullscreen } = useFullscreen()

  return {
    isCollapse,
    toggleCollapse,
    isFullscreen,
    toggleFullscreen
  }
})
