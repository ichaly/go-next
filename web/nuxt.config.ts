// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    srcDir: 'src/',
    devtools: {enabled: true},
    modules: [
        'nuxt-icon',
        '@unocss/nuxt',
        '@element-plus/nuxt',
        '@nuxtjs/i18n',
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
