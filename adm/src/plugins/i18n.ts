import { createI18n } from 'vue-i18n'

const storage = useStorage('lang', 'cn')

const messages = Object.entries(
  import.meta.glob<Record<string, unknown>>('/src/langs/*.ts', { eager: true })
).reduce((result, [path, module]) => {
  let key = path.match(/\/src\/langs\/(.*)\.ts$/)?.[1]
  return { ...result, [key!]: module.default }
}, {})

export const i18n = createI18n({
  messages,
  legacy: false, // Componsition API需要设置为false
  locale: storage.value, // 当前使用的语言类型
  globalInjection: true // 可以在template模板中使用$t
})
