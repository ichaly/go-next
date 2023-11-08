<script setup lang="ts">
const {locale, locales, setLocale} = useI18n()
const changeLang = () => {
  setLocale('en')
};
</script>

<template>
  <header class="bg-nav">
    <Container :fluid="false">
      <el-menu mode="horizontal" :ellipsis="false" class="!bg-transparent !border-b-0">
        <el-menu-item index="0" class="!inline-flex dark:!hidden !text-red">
          <i class="i-fa:dashcube text-4xl"/>
<!--          <span class="ml-0.5 text-2xl"><span class="text-black dark:text-white">Go</span>Next</span>-->
        </el-menu-item>
        <div class="flex-grow"/>
        <el-sub-menu index="1">
          <template #title>
            <i class="i-ri:earth-line text-2xl"/>
          </template>
          <el-menu-item index="2-1">English</el-menu-item>
          <el-menu-item index="2-2">简体中文</el-menu-item>
        </el-sub-menu>
        <el-menu-item index="2">
          <template #title>
            <i class="i-ri:github-fill text-2xl"/>
          </template>
        </el-menu-item>
        <el-menu-item index="3">登录/注册</el-menu-item>
      </el-menu>
    </Container>
  </header>
</template>

<style scoped lang="scss">
.bg-nav {
  @apply sticky w-full z-50 bg-[length:4px_4px] bg-transparent border-b-1 top-0;
  backdrop-filter: saturate(50%) blur(4px);
  background-image: radial-gradient(transparent 1px, white 1px);
}
</style>