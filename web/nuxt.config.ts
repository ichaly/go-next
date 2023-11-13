// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    srcDir: 'src/',
    devtools: {enabled: true},
    modules: [
        '@nuxtjs/i18n',
        '@vueuse/nuxt',
        '@unocss/nuxt',
        '@nuxt/content',
        '@element-plus/nuxt',
        '@pinia/nuxt',
        '@pinia-plugin-persistedstate/nuxt',
    ],
    css: [
        '@unocss/reset/tailwind.css',
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
    i18n: {
        strategy: 'no_prefix',
        defaultLocale: 'zh',
        locales: ['en', 'zh']
    }
})
