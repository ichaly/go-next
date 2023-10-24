// https://nuxt.com/docs/api/configuration/nuxt-config
// https://juejin.cn/post/7170746000112353293
// https://juejin.cn/post/7291133720902811707
export default defineNuxtConfig({
    srcDir: 'src/',
    devtools: {enabled: true},
    modules: [
        '@element-plus/nuxt',
        '@pinia/nuxt',
        '@pinia-plugin-persistedstate/nuxt',
    ],
    elementPlus: {
        importStyle: 'scss',
    },
    vite: {
        css: {
            preprocessorOptions: {
                scss: {
                    additionalData: '@use "@/assets/styles/default.scss" as *;'
                }
            }
        }
    },
    piniaPersistedstate: {
        cookieOptions: {
            sameSite: 'strict',
        },
        storage: 'localStorage',
    },
})
