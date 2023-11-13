<script setup lang="ts">
const isDark = useDark()
const {locale} = useI18n()
const toggleDark = useToggle(isDark)
</script>

<template>
  <header class="bg-nav">
    <div class="flex flex-row items-center m-a px-4 sm:px-6 lg:px-8 max-w-7xl">
      <div class="flex-1">
        <logo class="w-auto h-12"/>
      </div>
      <div class="sm:block">
        <el-menu class="!border-b-0 !bg-transparent" mode="horizontal">
          <Navigation/>
        </el-menu>
      </div>
      <div class="flex-1 flex items-center justify-end">
        <el-button @click="toggleDark()" link>
          <i class="i-ri:sun-line text-2xl hidden dark:block"/>
          <i class="i-ri:moon-line text-2xl block dark:hidden"/>
        </el-button>
        <el-button link>
          <el-dropdown trigger="click" @command="(lang: string) =>locale = lang">
            <i class="i-ri:translate text-2xl"/>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item :disabled="locale === 'zh'" command="zh">简体中文</el-dropdown-item>
                <el-dropdown-item :disabled="locale === 'en'" command="en">English</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </el-button>
        <el-button link>
          <i class="i-ri:github-fill text-2xl"/>
        </el-button>
      </div>
    </div>
  </header>
</template>

<style scoped lang="scss">
.bg-nav {
  @apply sticky w-full z-50 bg-[length:4px_4px] border-b-1 top-0 backdrop-blur;
  //backdrop-filter: saturate(50%) blur(4px);
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