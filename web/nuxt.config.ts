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
        '@nuxtjs/tailwindcss',
        '@nuxtjs/i18n',
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
        },
    },
    piniaPersistedstate: {
        cookieOptions: {
            sameSite: 'strict',
        },
        storage: 'localStorage',
    },
    // 国际化支持
    // https://blog.csdn.net/weixin_45978842/article/details/133840855
    i18n: {
        // 不启用国际化语言路由前缀模式
        strategy: 'no_prefix',
        defaultLocale: 'zh',
    }
})
