import Loading from '@/components/Loading.vue'
import type { AsyncComponentLoader } from 'vue'

export function useLoadingComponent(loader: AsyncComponentLoader) {
  return defineAsyncComponent({
    loader,
    loadingComponent: Loading
  })
}
