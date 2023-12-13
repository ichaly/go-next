import 'virtual:uno.css'
import 'virtual:unocss-devtools'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import { setupRouter } from '@/router'
import { useRootStore } from '@/stores/root'

const app = createApp(App)

app.use(createPinia())
//提前加载数据解决刷新404的问题
const { loadMenus } = useRootStore()
await loadMenus()
setupRouter(app)

app.mount('#app')
