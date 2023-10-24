import en_us from "assets/lang/en_us";
import zh_cn from "assets/lang/zh_cn";

export default defineI18nConfig(() => ({
    legacy: true,
    locale: 'zh',
    messages: {
        en: en_us,
        zh: zh_cn
    },
}))