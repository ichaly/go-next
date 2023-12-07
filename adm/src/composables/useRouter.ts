// const router = useBaseRouter()

export function useRouter() {
  // console.log(router)
  // router.push = (to: RouteLocationRaw): Promise<NavigationFailure | void | undefined> => {
  //   if (typeof to === 'string') {
  //     //正则判断是不是网页链接
  //     if (/^https?:\/\//.test(to)) {
  //       window.open(to)
  //       return Promise.resolve()
  //     }
  //     //判断是否以斜杠开头,如果不是则在前边拼接斜杠
  //     if (!/^\//.test(to)) {
  //       to = `/${to}`
  //     }
  //     // } else if (to instanceof MatcherLocationAsPath) {
  //     //   //正则判断是不是网页链接
  //     //   if (/^https?:\/\//.test(to.path)) {
  //     //     window.open(to.path)
  //     //     return Promise.resolve()
  //     //   }
  //     //   //判断是否以斜杠开头,如果不是则在前边拼接斜杠
  //     //   if (!/^\//.test(to.path)) {
  //     //     to.path = `/${to.path}`
  //     //   }
  //   }
  //
  //   return router.push(to)
  // }

  return useBaseRouter()
}
