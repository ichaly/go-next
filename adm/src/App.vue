<template>
  <RouterView v-if="isAlive" />
  <div v-if="false">
    <i class="i-ri:apps-line" />
    <i class="i-ri:settings-4-line" />
    <i class="i-ri:file-user-line" />
    <i class="i-ri:shield-keyhole-line" />
    <i class="i-ri:contacts-book-3-line" />
    <i class="i-ri:t-shirt-2-line" />
    <i class="i-ri:team-line" />
    <i class="i-ri:organization-chart" />
    <i class="i-ri:file-edit-line" />
    <i class="i-ri:file-list-line" />
    <i class="i-ri:file-info-line" />
  </div>
</template>
<script setup lang="ts">
import { RELOAD_KEY } from '@/plugins/router'

const isAlive = ref(true)
const reload = () => {
  isAlive.value = false
  nextTick(() => {
    isAlive.value = true
  })
}
provide(RELOAD_KEY, reload)
</script>
