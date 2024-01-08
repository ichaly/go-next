<template>
  <el-dropdown class="h-full" @command="onCommand">
    <flat-button>
      <el-icon :size="18">
        <i class="i-ri:translate-2" />
      </el-icon>
    </flat-button>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item
          v-for="lang in $i18n.availableLocales"
          :key="lang"
          :command="lang"
          :disabled="locale === lang">
          <span v-t="{path:'name',locale:lang}"></span>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts">
import { RELOAD_KEY } from '@/plugins/router'

const { locale } = useI18n()
const reload = inject(RELOAD_KEY)
const storage = useStorage('lang', 'cn')

const onCommand = (val: string) => {
  locale.value = val
  storage.value = val
  nextTick(() => {
    reload?.()
  })
}
</script>