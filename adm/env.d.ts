/// <reference types="vite/client" />
import 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    items?: Item[]
    icon?: string
    title?: string
    default?: boolean
  }
}

declare global {
  interface ImportMetaEnv {
    readonly VITE_BASE_API: string
  }
}
