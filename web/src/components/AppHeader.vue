<script setup lang="ts">
const color = ref('#0E9502')

const isDark = useDark()
const toggleDark = useToggle(isDark)

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
          <Navigation/>
        </el-menu>
      </div>
      <div class="flex-1 flex items-center justify-end">
        <el-button link>
          <i class="i-ri:search-line dark:text-white text-2xl"/>
        </el-button>
        <el-button @click="toggleDark()" link>
          <i class="i-ri:moon-line text-2xl dark:text-white" v-if="isDark"/>
          <i class="i-ri:sun-line text-2xl dark:text-white" v-else/>
        </el-button>
        <el-button link>
          <i class="i-ri:github-line dark:text-white text-2xl"/>
        </el-button>
        <el-button link>
          <client-only>
            <el-color-picker v-model="color"/>
            <template #fallback>
              <p class="w-8"></p>
            </template>
          </client-only>
        </el-button>
      </div>
    </div>
  </header>
</template>

<style scoped lang="scss">
.bg-nav {
  @apply sticky w-full z-50 bg-[length:4px_4px] bg-transparent border-b-1 top-0;
  backdrop-filter: saturate(50%) blur(4px);
  //background-image: radial-gradient(transparent 1px, white 1px);
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