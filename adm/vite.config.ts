import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import UnoCSS from 'unocss/vite'
import { viteMockServe } from 'vite-plugin-mock'
import topLevelAwait from 'vite-plugin-top-level-await'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import type { ComponentResolver } from 'unplugin-vue-components/types'
import * as fs from 'fs'

export interface IndexResolverOptions {}

function IndexResolver(options: IndexResolverOptions): ComponentResolver[] {
  return [
    {
      type: 'component',
      resolve: async (name: string) => {
        let url = fileURLToPath(new URL(`./src/components/${name}/index.ts`, import.meta.url))
        if (fs.existsSync(url)) {
          return { name: 'default', from: url }
        }
        url = fileURLToPath(new URL(`./src/components/${name}/index.vue`, import.meta.url))
        if (fs.existsSync(url)) {
          return { name: 'default', from: url }
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
        IndexResolver({})
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
