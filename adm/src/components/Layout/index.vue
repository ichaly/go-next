<template>
  <el-container class="w-full h-full">
    <el-aside class="!w-auto bg-[var(--left-menu-bg-color)] flex flex-col">
      <logo />
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
      <el-header class="flex items-center b-0 b-b-1 b-solid b-[var(--el-border-color)]">
        <flat-button @click="toggleCollapse()">
          <el-icon :size="18" :class="['transition-transform', { '-scale-x-100': isCollapse }]">
            <i class="i-icon-park-outline:menu-unfold-one" />
          </el-icon>
        </flat-button>
        <div class="flex-1 pl-2">
          <el-breadcrumb class="hidden md:block">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>
              <router-link to="/adm/public">用户管理</router-link>
            </el-breadcrumb-item>
            <el-breadcrumb-item>用户详情</el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <flat-button>
          <el-icon :size="18">
            <i class="i-icon-park-outline:search" />
          </el-icon>
        </flat-button>
        <flat-button @click="toggleFullscreen()">
          <el-icon :size="18">
            <i class="i-icon-park-outline:off-screen-one" v-if="isFullscreen" />
            <i class="i-icon-park-outline:full-screen-one" v-else />
          </el-icon>
        </flat-button>
        <flat-button>
          <el-badge :value="200" :max="99" class="h-4.5">
            <el-icon :size="18">
              <i class="i-icon-park-outline:remind" />
            </el-icon>
          </el-badge>
        </flat-button>
        <el-dropdown class="h-full">
          <flat-button>
            <el-icon :size="18">
              <i class="i-icon-park-outline:translate" />
            </el-icon>
          </flat-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>
                <span>简体中文</span>
              </el-dropdown-item>
              <el-dropdown-item>
                <span>English</span>
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
        <el-dropdown class="h-full">
          <flat-button>
            <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png">
              user
            </el-avatar>
          </flat-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>
                <i class="i-icon-park-outline:user" />
                <span>个人信息</span>
              </el-dropdown-item>
              <el-dropdown-item>
                <i class="i-icon-park-outline:edit" />
                <span>修改密码</span>
              </el-dropdown-item>
              <el-dropdown-item divided>
                <i class="i-icon-park-outline:help" />
                <span>帮助文档</span>
              </el-dropdown-item>
              <el-dropdown-item>
                <i class="i-icon-park-outline:tips" />
                <span>功能更新</span>
              </el-dropdown-item>
              <el-dropdown-item divided>
                <i class="i-icon-park-outline:lock" />
                <span>锁定屏幕</span>
              </el-dropdown-item>
              <el-dropdown-item>
                <i class="i-icon-park-outline:power" />
                <span>退出登录</span>
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-header>
      <el-main class="h-full">
        <router-view />
      </el-main>
      <el-footer
        class="text-gray-500 text-sm flex flex-col items-center justify-center md:flex-row md:justify-between b-0 b-solid b-t-1 b-[var(--el-border-color)]"
      >
        <div>Released under the MIT License.</div>
        <div>Copyright © 2023 iChaly.</div>
      </el-footer>
    </el-container>
  </el-container>
  <div v-if="false">
    <i class="i-icon-park-outline:system" />
    <i class="i-icon-park-outline:setting-two" />
    <i class="i-icon-park-outline:permissions" />
    <i class="i-icon-park-outline:audit" />
    <i class="i-icon-park-outline:add-user" />
    <i class="i-icon-park-outline:personal-privacy" />
    <i class="i-icon-park-outline:data-user" />
    <i class="i-icon-park-outline:table-file" />
  </div>
</template>

<script setup lang="ts">
const router = useRouter()
const rootStore = useRootStore()
const { toggleCollapse, toggleFullscreen } = rootStore
const { menus, isCollapse, isFullscreen } = toRefs(rootStore)
const selectMenuItem = (index: any, path: any, item: any, result: any) => {
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
