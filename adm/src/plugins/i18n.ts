import { createI18n } from 'vue-i18n'

const storage = useStorage('lang', 'cn')

const messages = Object.entries(
  import.meta.glob<Record<string, unknown>>('/src/locales/*.ts', { eager: true })
).reduce((result, [path, module]) => {
  let key = path.match(/\/src\/locales\/(.*)\.ts$/)?.[1]
  return { ...result, [key!]: module.default }
}, {})

export const i18n = createI18n({
  messages,
  legacy: false, // Componsition API需要设置为false
  fallbackLocale: 'cn', // 默认语言
  locale: storage.value, // 当前使用的语言类型
  globalInjection: true // 可以在template模板中使用$t
})