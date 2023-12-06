// const views = import.meta.glob('/src/views/modules/**/*.vue')

export const pages: Record<string, { path: string, component: RawRouteComponent }> = {
  'dashboard': {
    path: '/dashboard',
    component: () => import('@/views/modules/Dashboard.vue')
  }
}