import { fileURLToPath, URL } from 'node:url'
import fg from 'fast-glob'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import UnoCSS from 'unocss/vite'
import { viteMockServe } from 'vite-plugin-mock'
import topLevelAwait from 'vite-plugin-top-level-await'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import type { ComponentResolver } from 'unplugin-vue-components/types'

export interface IndexResolverOptions {
  exclude?: RegExp,
}

function IndexResolver(options: IndexResolverOptions): ComponentResolver[] {
  return [
    {
      type: 'component',
      resolve: async (name: string) => {
        if (options.exclude && name.match(options.exclude))
          return
        let files = fg.globSync(`./src/components/${name}/index.{ts,js,vue}`)
        for (let file of files) {
          let from = fileURLToPath(new URL(file, import.meta.url))
          return { name: 'default', from }
        }
      }
    }
  ]
}

// https://vitejs.dev/config/
export default defineConfig(({ command }) => ({
  envDir: 'env',
  plugins: [
    vue(),
    UnoCSS(),
    viteMockServe({
      mockPath: './mock',
      enable: command === 'serve'
    }),
    topLevelAwait({
      promiseExportName: '__tla',
      promiseImportName: (i) => `__tla_${i}`
    }),
    AutoImport({
      dirs: ['src/stores', 'src/composables'],
      imports: [
        'vue',
        '@vueuse/core',
        {
          'vue-router': [
            'useLink',
            'useRoute',
            'onBeforeRouteLeave',
            'onBeforeRouteUpdate',
            ['useRouter', 'useBaseRouter']
          ]
        }
      ],
      resolvers: [
        ElementPlusResolver({
          importStyle: 'sass'
        })
      ]
    }),
    Components({
      dirs: ['src/components'],
      resolvers: [
        ElementPlusResolver({
          importStyle: 'sass'
        }),
        IndexResolver({
          exclude: /(RouterLink|RouterView)$/
        })
      ]
    })
  ],
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `@use "@/assets/styles/index.scss" as *;`
      }
    }
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
}))
