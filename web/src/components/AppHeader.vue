<script setup lang="ts">
import Logo from "~/components/Logo.vue";

const color = ref('#0E9502')

const colorMode = useColorMode()
const isDark = computed({
  get() {
    return colorMode.value === 'dark'
  },
  set() {
    colorMode.preference = colorMode.value === 'dark' ? 'light' : 'dark'
  }
})

const {locale, locales, setLocale} = useI18n()
const changeLang = () => {
  setLocale('en')
};
</script>

<template>
  <header class="bg-nav">
    <div class="flex flex-row items-center m-a px-4 sm:px-6 lg:px-8 max-w-7xl">
      <div class="flex-1">
        <logo class="w-auto h-12 m-1"/>
      </div>
      <div class="sm:block">
        <el-menu class="!border-b-0 !bg-transparent" mode="horizontal">
          <el-menu-item index="1">Processing Center</el-menu-item>
          <el-sub-menu index="2">
            <template #title>Workspace</template>
            <el-menu-item index="2-1">item one</el-menu-item>
            <el-menu-item index="2-2">item two</el-menu-item>
            <el-menu-item index="2-3">item three</el-menu-item>
            <el-sub-menu index="2-4">
              <template #title>item four</template>
              <el-menu-item index="2-4-1">item one</el-menu-item>
              <el-menu-item index="2-4-2">item two</el-menu-item>
              <el-menu-item index="2-4-3">item three</el-menu-item>
            </el-sub-menu>
          </el-sub-menu>
          <el-menu-item index="3" disabled>Info</el-menu-item>
          <el-menu-item index="4">Orders</el-menu-item>
        </el-menu>
      </div>
      <div class="flex-1 flex items-center justify-end">
        <el-button link>
          <client-only>
            <el-color-picker v-model="color"/>
          </client-only>
        </el-button>
        <el-button link>
          <i class="i-ri:search-line text-2xl"/>
        </el-button>
        <el-button @click="isDark = !isDark" link>
          <i class="i-ri:moon-line text-2xl" v-if="isDark"/>
          <i class="i-ri:sun-line text-2xl" v-else/>
        </el-button>
        <el-button link>
          <i class="i-ri:github-line text-2xl"/>
        </el-button>
        <el-button type="primary" plain class="w-0">登录 / 注册</el-button>
      </div>
    </div>
  </header>
</template>

<style scoped lang="scss">
.bg-nav {
  @apply sticky w-full z-50 bg-[length:4px_4px] bg-transparent border-b-1 top-0;
  backdrop-filter: saturate(50%) blur(4px);
  background-image: radial-gradient(transparent 1px, white 1px);
}

.base {
  @apply m-a w-full;
}

.padding {
  @apply px-4 sm:px-6 lg:px-8;
}

.constrained {
  @apply max-w-7xl;
}
</style>