<script setup lang="ts">
const isDark = useDark()
const {locale} = useI18n()
const toggleDark = useToggle(isDark)
const {y} = useWindowScroll()
const clazz = computed(() => {
  return ['bg-nav', {'!border-b-1': y.value > 0}]
})
</script>

<template>
  <header :class="clazz">
    <Container>
      <div class="flex flex-row items-center">
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
            <client-only>
              <el-dropdown trigger="click" @command="(lang: string) =>locale = lang">
                <i class="i-ri:translate text-2xl"/>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item :disabled="locale === 'zh'" command="zh">简体中文</el-dropdown-item>
                    <el-dropdown-item :disabled="locale === 'en'" command="en">English</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </client-only>
          </el-button>
          <el-button link>
            <i class="i-ri:github-fill text-2xl"/>
          </el-button>
        </div>
      </div>
    </Container>
  </header>
</template>

<style scoped lang="scss">
.bg-nav {
  @apply fixed w-full z-50 bg-[length:4px_4px] top-0 backdrop-blur;
  //backdrop-filter: saturate(50%) blur(4px);
  //background-image: radial-gradient(transparent 1px, white 1px);
}
</style>