// uno.config.ts
import {
  defineConfig,
  presetUno,
  presetIcons,
  presetAttributify,
  transformerDirectives,
  transformerVariantGroup
} from 'unocss'

export default defineConfig({
  shortcuts: [
    ['line', 'border-solid border-[var(--el-menu-border-color)] border-0 border-b-1'],
    ['line-t', 'border-solid border-[var(--el-menu-border-color)] border-0 border-t-1']
  ],
  presets: [
    presetUno(),
    presetIcons({
      cdn: 'https://esm.sh/'
    }),
    presetAttributify()
  ],
  transformers: [transformerDirectives(), transformerVariantGroup()]
})
