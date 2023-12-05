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
  shortcuts: [['center', 'items-center justify-center']],
  presets: [
    presetUno(),
    presetIcons({
      // cdn: 'https://esm.sh/',
      collections: {
        ep: () => import('@iconify-json/ep/icons.json').then((i) => i.default),
        ri: () => import('@iconify-json/ri/icons.json').then((i) => i.default),
        mdi: () => import('@iconify-json/mdi/icons.json').then((i) => i.default),
        carbon: () => import('@iconify-json/carbon/icons.json').then((i) => i.default),
        'icon-park-outline': () =>
          import('@iconify-json/icon-park-outline/icons.json').then((i) => i.default)
      }
    }),
    presetAttributify()
  ],
  transformers: [transformerDirectives(), transformerVariantGroup()]
})
