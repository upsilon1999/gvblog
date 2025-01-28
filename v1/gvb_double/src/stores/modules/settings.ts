import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useSettingStore = defineStore('settings', () => {
  //收缩状态
  const iscollapsed = ref(false)
  //获取收缩状态
  const getCollapsed = () => {
    return iscollapsed.value
  }
  //设置收缩状态
  const setCollapsed = (state: boolean) => {
    iscollapsed.value = state
  }

  //主题：true白天、false黑夜
  const theme = ref(true)
  const getTheme = () => {
    return theme.value
  }
  //设置收缩状态
  const setTheme = () => {
    theme.value = !theme.value
  }

  return {
    getCollapsed,
    setCollapsed,

    //主题
    getTheme,
    setTheme,
  }
})

export default useSettingStore
