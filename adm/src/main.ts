import 'virtual:uno.css'
import 'virtual:unocss-devtools'
import '@unocss/reset/tailwind.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import { setupRouter } from '@/plugins/router'
import { useRootStore } from '@/stores/root'
import { i18n } from '@/plugins/i18n'

const app = createApp(App)

// 状态管理
app.use(createPinia())
// 提前加载数据解决刷新404的问题
const { loadMenus } = useRootStore()
await loadMenus()
setupRouter(app)
// 国际化
app.use(i18n)

app.mount('#app')
