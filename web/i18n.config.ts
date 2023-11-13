import en from "assets/lang/en";
import zh from "assets/lang/zh";

export default defineI18nConfig(() => ({
    legacy: true,
    locale: 'zh',
    messages: {
        en: en,
        zh: zh
    },
}))