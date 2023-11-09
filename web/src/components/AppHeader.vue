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
    <div class="container flex flex-row items-center m-a">
      <logo class="w-auto h-12 m-1"/>
      <div class="flex-grow justify-center items-center"/>
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
  </header>
</template>

<style scoped lang="scss">
.bg-nav {
  @apply sticky w-full z-50 bg-[length:4px_4px] bg-transparent border-b-1 top-0;
  backdrop-filter: saturate(50%) blur(4px);
  background-image: radial-gradient(transparent 1px, white 1px);
}
</style>