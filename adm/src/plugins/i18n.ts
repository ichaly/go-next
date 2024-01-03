import en from '@/langs/en' // 英文语言配置
import zh from '@/langs/zh' // 中文语言配置
import { createI18n } from 'vue-i18n'

const storage = useStorage('lang', 'zh')
export const i18n = createI18n({
  legacy: false, // Componsition API需要设置为false
  locale: storage.value,// 当前使用的语言类型
  globalInjection: true, // 可以在template模板中使用$t
  messages: { en, zh }
})