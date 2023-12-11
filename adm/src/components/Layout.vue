<template>
  <el-container class="w-full h-full">
    <el-aside class="!w-auto bg-[var(--left-menu-bg-color)] flex flex-col">
      <LogoView />
      <el-scrollbar class="flex-1">
        <el-menu
          :collapse="isCollapse"
          @select="selectMenuItem"
          default-active="/index1/sub1"
          text-color="var(--left-menu-text-color)"
          background-color="var(--left-menu-bg-color)"
          active-text-color="var(--left-menu-text-active-color)"
        >
          <MenuTree :menus="menus" />
        </el-menu>
      </el-scrollbar>
    </el-aside>
    <el-container>
      <el-header>
        <Header />
      </el-header>
      <el-main class="h-full">
        <router-view />
      </el-main>
      <el-footer>
        <Footer />
      </el-footer>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
const router = useRouter()
const rootStore = useRootStore()
const { menus, isCollapse } = toRefs(rootStore)
const selectMenuItem = (index: any) => {
  router.push(index)
}
</script>

<style scoped lang="scss">
:deep(.el-menu) {
  width: var(--left-menu-max-width);
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

// 折叠时的最小宽度
:deep(.el-menu--collapse) {
  width: var(--left-menu-min-width);

  & > .is-active,
  & > .is-active > .el-sub-menu__title {
    color: var(--left-menu-text-active-color) !important;
    background-color: var(--left-menu-collapse-bg-active-color) !important;
  }
}

//解决下拉菜单的focus样式有个黑框的问题
:deep(.el-tooltip__trigger:focus-visible) {
  outline: unset;
}
</style>
