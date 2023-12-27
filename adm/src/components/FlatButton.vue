<template>
  <div class="flat-btn">
    <slot />
  </div>
</template>
<style scoped lang="scss">
.flat-btn {
  @apply h-full px-5 cursor-pointer flex center;
  @apply text-[var(--el-button-text-color)] text-[length:var(--el-font-size-base)];
  @apply hover:bg-[var(--el-fill-color-light)] hover:text-[var(--el-color-primary)];
}
</style>
