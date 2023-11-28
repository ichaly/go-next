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
  presets: [
    presetUno(),
    presetIcons({
      // cdn: 'https://esm.sh/',
      collections: {
        ep: () => import('@iconify-json/ep/icons.json').then(i => i.default),
        ri: () => import('@iconify-json/ri/icons.json').then(i => i.default),
        'icon-park-outline': () => import('@iconify-json/icon-park-outline/icons.json').then(i => i.default),
      }
    }),
    presetAttributify()
  ],
  transformers: [transformerDirectives(), transformerVariantGroup()]
})
