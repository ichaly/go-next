import type { RouteLocationRaw, Router } from 'vue-router'

//利用AutoImports的别名机制完成偷梁换柱,使路由支持打开外部链接
export function useRouter(): Router {
  const { push, ...rest } = useBaseRouter()
  return {
    push: (to: RouteLocationRaw) => {
      if (typeof to === 'string' && /^https?:\/\//.test(to)) {
        window.open(to)
        return Promise.resolve()
      }
      return push(to)
    },
    ...rest
  }
}
