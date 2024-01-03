<template>
  <RouterView v-if="isAlive" />
  <div v-if="false">
    <i class="i-icon-park-outline:system" />
    <i class="i-icon-park-outline:setting-two" />
    <i class="i-icon-park-outline:permissions" />
    <i class="i-icon-park-outline:audit" />
    <i class="i-icon-park-outline:add-user" />
    <i class="i-icon-park-outline:personal-privacy" />
    <i class="i-icon-park-outline:data-user" />
    <i class="i-icon-park-outline:table-file" />
    <i class="i-icon-park-outline:edit-one" />
    <i class="i-icon-park-outline:view-list" />
    <i class="i-icon-park-outline:doc-detail" />
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
