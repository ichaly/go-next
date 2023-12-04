import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', () => {
  const [isCollapse, toggleCollapse] = useToggle(false)
  const { isFullscreen, toggle: toggleFullscreen } = useFullscreen()

  return {
    isCollapse,
    toggleCollapse,
    isFullscreen,
    toggleFullscreen
  }
})
