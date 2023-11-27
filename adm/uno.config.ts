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
      cdn: 'https://esm.sh/'
    }),
    presetAttributify()
  ],
  transformers: [transformerDirectives(), transformerVariantGroup()]
})
