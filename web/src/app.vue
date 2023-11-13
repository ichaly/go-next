<template>
  <el-config-provider :locale="language">
    <NuxtLayout>
      <NuxtLoadingIndicator/>
      <NuxtPage/>
      <ElBacktop/>
    </NuxtLayout>
  </el-config-provider>
</template>
<script setup lang="ts">
import zh from 'element-plus/es/locale/lang/zh-cn'
import en from 'element-plus/es/locale/lang/en'

const {locale} = useI18n()
const language = computed(() => {
  return locale.value === 'zh' ? zh : en
})
</script>