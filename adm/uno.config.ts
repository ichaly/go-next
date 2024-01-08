// uno.config.ts
import {
  defineConfig,
  presetAttributify,
  presetIcons,
  presetUno,
  transformerDirectives,
  transformerVariantGroup
} from 'unocss'

export default defineConfig({
  shortcuts: [
    { 'center': 'items-center justify-center' },
    { 'b-light': 'b-[var(--el-border-color-light)]' }
  ],
  presets: [
    presetUno(),
    presetIcons({
      // autoInstall: true
      // cdn: 'https://esm.sh/',
      // cdn: 'https://cdn.skypack.dev/'
      collections: {
        ep: () => import('@iconify-json/ep/icons.json').then((i) => i.default),
        ri: () => import('@iconify-json/ri/icons.json').then((i) => i.default),
        mdi: () => import('@iconify-json/mdi/icons.json').then((i) => i.default)
      }
    }),
    presetAttributify()
  ],
  transformers: [transformerDirectives(), transformerVariantGroup()]
})
