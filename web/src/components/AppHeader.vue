<script setup lang="ts">
const {locale, locales, setLocale} = useI18n()
const changeLang = () => {
  setLocale('en')
};
</script>

<template>
  <div class="bg-nav">
    <el-menu mode="horizontal" :ellipsis="false" class="!bg-transparent">
      <el-menu-item index="0">
        <img class="h-8" src="@/assets/images/go-next-logo.svg" alt="Go Next Logo"/>
        <span class="ml-1 text-2xl">GONEXT</span>
      </el-menu-item>
      <div class="flex-grow"/>
      <el-sub-menu index="1">
        <template #title>
          <i class="i-ri:earth-line text-2xl" />
        </template>
        <el-menu-item index="2-1">English</el-menu-item>
        <el-menu-item index="2-2">简体中文</el-menu-item>
      </el-sub-menu>
      <el-menu-item index="2">
        <template #title>
          <i class="i-ri:github-fill text-2xl" />
        </template>
      </el-menu-item>
      <el-menu-item index="3">登录/注册</el-menu-item>
    </el-menu>
  </div>
</template>

<style scoped lang="scss">
.bg-nav {
  @apply fixed w-full z-50 bg-[length:4px_4px] bg-transparent;
  backdrop-filter: saturate(50%) blur(4px);
  background-image: radial-gradient(transparent 1px, white 1px);
}
</style>