/// <reference types="vite/client" />
import 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    icon?: string
    title?: string
    weight?: number
    routes?: Item[]
  }
}

declare global {
  interface ImportMetaEnv {
    readonly VITE_BASE_API: string
  }
}
