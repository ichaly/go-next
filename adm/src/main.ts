import 'virtual:uno.css'
import 'virtual:unocss-devtools'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import { setupRouter } from '@/router'

const app = createApp(App)

app.use(createPinia())
setupRouter(app)

app.mount('#app')
