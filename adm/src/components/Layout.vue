<template>
  <el-container class="w-full h-full">
    <div :class="['flex', 'flex-col', 'w-auto']">
      <div class="logo">
        <img src="@/assets/logo.svg" class="w-8 h-8" />
        <span class="ml-4" v-if="!isCollapse">Vue Admin</span>
      </div>
      <el-aside class="flex-1 w-auto bg-[var(--left-menu-bg-color)]">
        <el-scrollbar>
          <el-menu default-active="2" :collapse="isCollapse">
            <el-sub-menu index="1">
              <template #title>
                <el-icon><i class="i-ep:location" /></el-icon>
                <span>Navigator One</span>
              </template>
              <el-menu-item index="1-1">item one</el-menu-item>
              <el-menu-item index="1-2">item two</el-menu-item>
              <el-menu-item index="1-3">item three</el-menu-item>
              <el-sub-menu index="1-4">
                <template #title><span>item four</span></template>
                <el-menu-item index="1-4-1">item one</el-menu-item>
              </el-sub-menu>
            </el-sub-menu>
            <el-menu-item index="2">
              <el-icon><i class="i-ep:menu" /></el-icon>
              <template #title>Navigator Two</template>
            </el-menu-item>
            <el-menu-item index="3">
              <el-icon><i class="i-ep:document" /></el-icon>
              <template #title>Navigator Three</template>
            </el-menu-item>
            <el-menu-item index="4">
              <el-icon><i class="i-ep:setting" /></el-icon>
              <template #title>Navigator Four</template>
            </el-menu-item>
          </el-menu>
        </el-scrollbar>
      </el-aside>
    </div>
    <el-container>
      <el-header class="flex items-center">
        <el-button text @click="toggleCollapse()" class="p-2">
          <el-icon :size="18" class="cursor-pointer">
            <i class="i-ep:expand" v-if="isCollapse" />
            <i class="i-ep:fold" v-else />
          </el-icon>
        </el-button>
        Header
      </el-header>
      <el-main class="h-full">
        <router_view />
      </el-main>
      <el-footer class="flex items-center">Footer</el-footer>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
const [isCollapse, toggleCollapse] = useToggle(true)
</script>

<style scoped lang="scss">
.logo {
  @apply w-full flex items-center text-base px-4 text-white;
  height: $item-height;
  background-color: var(--left-menu-bg-color);
}

:deep(.el-menu) {
  max-width: 200px;
  // 去掉右侧的边框
  border-right: none;
  background-color: var(--left-menu-bg-color) !important;

  // 设置子菜单悬停的高亮和背景色
  .el-sub-menu__title,
  .el-menu-item {
    color: var(--left-menu-text-color);

    &:hover {
      color: var(--left-menu-text-active-color) !important;
      background-color: var(--left-menu-bg-color) !important;
    }
  }

  // 设置选中时的背景和颜色
  .el-menu-item.is-active {
    color: var(--left-menu-text-active-color) !important;
    background-color: var(--left-menu-bg-active-color) !important;
  }

  // 设置子菜单的背景颜色
  .el-menu {
    .el-sub-menu__title,
    .el-menu-item:not(.is-active) {
      background-color: var(--left-menu-bg-light-color) !important;
    }
  }
}
</style>
