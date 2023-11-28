<template>
  <el-container class="w-full h-full">
    <div class="flex flex-col w-auto">
      <div class="logo">
        <img src="@/assets/logo.svg" class="w-8 h-8" />
        <span class="ml-4" v-if="!isCollapse">Vue Admin</span>
      </div>
      <el-aside class="flex-1 !w-auto bg-[var(--left-menu-bg-color)]">
        <el-scrollbar>
          <el-menu
            default-active="2"
            :collapse="isCollapse"
            textColor="var(--left-menu-text-color)"
            backgroundColor="var(--left-menu-bg-color)"
            activeTextColor="var(--left-menu-text-active-color)"
          >
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
      <el-header class="flex items-center b-0 b-b-1 b-solid b-[var(--el-border-color)]">
        <flat-button @click="toggleCollapse()">
          <el-icon :size="18">
            <i class="i-icon-park-outline:menu-fold-one" v-if="isCollapse" />
            <i class="i-icon-park-outline:menu-unfold-one" v-else />
          </el-icon>
        </flat-button>
        <div class="flex-1">
          <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item><a href="/">用户管理</a></el-breadcrumb-item>
            <el-breadcrumb-item>用户详情</el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <flat-button>
          <el-icon :size="18">
            <i class="i-icon-park-outline:search" />
          </el-icon>
        </flat-button>
        <flat-button>
          <el-badge :value="200" :max="99">
            <el-icon :size="18">
              <i class="i-icon-park-outline:remind" />
            </el-icon>
          </el-badge>
        </flat-button>
        <flat-button @click="toggleFullscreen()">
          <el-icon :size="18">
            <i class="i-icon-park-outline:off-screen-one" v-if="isFullscreen" />
            <i class="i-icon-park-outline:full-screen-one" v-else />
          </el-icon>
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
        <el-dropdown>
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
        <router_view />
      </el-main>
      <el-footer class="flex items-center">Footer</el-footer>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
const [isCollapse, toggleCollapse] = useToggle()
const { isFullscreen, toggle: toggleFullscreen } = useFullscreen()
</script>

<style scoped lang="scss">
.logo {
  @apply w-full flex items-center text-base px-4 text-white;
  height: $item-height;
  background-color: var(--left-menu-bg-color);
}

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

:deep(.el-button + .el-button) {
  margin-left: 0;
}
</style>
