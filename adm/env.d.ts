/// <reference types="vite/client" />
import 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    name: string
    icon?: string
    title?: string
    items?: Item[]
    default?: boolean
  }
}

declare global {
  interface ImportMetaEnv {
    readonly VITE_BASE_API: string
  }
}
