<template>
  <el-dropdown class="h-full" @command="onCommand">
    <flat-button>
      <el-icon :size="18">
        <i class="i-icon-park-outline:translate" />
      </el-icon>
    </flat-button>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item v-for="({label, value}) in langs" :command="value" :disabled="current===value" :key="value">
          <span>{{ label }}</span>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts">
import { RELOAD_KEY } from '@/plugins/router'

const { locale } = useI18n()
const reload = inject(RELOAD_KEY)
const storage = useStorage('lang', 'zh')
const langs = ref([
  { label: '简体中文', value: 'zh' },
  { label: 'English', value: 'en' }
])
const current = computed(() => {
  return locale.value || storage.value
})

const onCommand = (val: string) => {
  locale.value = val
  storage.value = val
  reload?.()
}
</script>

<style scoped lang="scss"></style>
