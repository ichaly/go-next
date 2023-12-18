/// <reference types="vite/client" />
import 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    items?: Item[]
  }
}

declare global {
  interface ImportMetaEnv {
    readonly VITE_BASE_API: string
  }
}
