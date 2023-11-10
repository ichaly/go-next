<template>
  <el-config-provider :locale="zh">
    <NuxtLayout>
      <NuxtLoadingIndicator/>
      <NuxtPage/>
      <ElBacktop/>
    </NuxtLayout>
  </el-config-provider>
</template>
<script setup lang="ts">
// import {useI18n} from 'vue-i18n'
import zh from 'element-plus/es/locale/lang/zh-cn'
// import en from 'element-plus/es/locale/lang/en'

// let locale11 = ref(en)
// const {t, locale} = useI18n({useScope: 'global'})
// const switchLang = () => {
//   locale11 = zh
//   locale.value = 'zh'
// }
</script>