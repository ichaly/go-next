<template>
  <div class="logo">
    <img src="@/assets/logo.svg" class="w-8 h-8" />
    <span class="ml-4" v-if="!isCollapse">Vue Admin</span>
  </div>
</template>

<script setup lang="ts">
const { isCollapse } = toRefs(useSettingStore())
</script>

<style scoped lang="scss">
.logo {
  @apply w-full flex items-center text-base px-4 text-white;
  background-color: var(--left-menu-bg-color);
  height: $item-height;
}
</style>
