import Loading from '@/components/Loading.vue'
import type { AsyncComponentLoader } from 'vue'

export function useLoadingComponent(loader: AsyncComponentLoader) {
  const component = defineAsyncComponent({
    loader,
    loadingComponent: Loading
  })
  //消除VueRouter的警告,参考https://github.com/vuejs/router/issues/1469#issuecomment-1458400428
  component.__warnedDefineAsync = true
  return component
}
