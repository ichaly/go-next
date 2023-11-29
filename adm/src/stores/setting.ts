import { defineStore } from 'pinia'

export const useSettingStore = defineStore('setting', () => {
  const [isCollapse, toggleCollapse] = useToggle(true)
  const { isFullscreen, toggle: toggleFullscreen } = useFullscreen()

  return {
    isCollapse,
    toggleCollapse,
    isFullscreen,
    toggleFullscreen
  }
})
