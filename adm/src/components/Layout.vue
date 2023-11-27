<template>
  <el-container class="w-full h-full">
    <div class="w-50 flex flex-col bg-[var(--left-menu-bg-color)]">
      <div class="w-full h-12 flex items-center text-base px-4 text-white">
        <img src="@/assets/logo.svg" class="w-9 h-9" />Admin
      </div>
      <el-aside class="flex-1 !w-full">
        <el-scrollbar>
          <el-menu :default-openeds="['1', '3']">
            <el-sub-menu index="1">
              <template #title>
                <el-icon>
                  <message />
                </el-icon>
                Navigator One
              </template>
              <el-menu-item-group>
                <template #title>Group 1</template>
                <el-menu-item index="1-1">Option 1</el-menu-item>
                <el-menu-item index="1-2">Option 2</el-menu-item>
              </el-menu-item-group>
              <el-menu-item-group title="Group 2">
                <el-menu-item index="1-3">Option 3</el-menu-item>
              </el-menu-item-group>
              <el-sub-menu index="1-4">
                <template #title>Option4</template>
                <el-menu-item index="1-4-1">Option 4-1</el-menu-item>
              </el-sub-menu>
            </el-sub-menu>
            <el-sub-menu index="2">
              <template #title>
                <el-icon>
                  <icon-menu />
                </el-icon>
                Navigator Two
              </template>
              <el-menu-item-group>
                <template #title>Group 1</template>
                <el-menu-item index="2-1">Option 1</el-menu-item>
                <el-menu-item index="2-2">Option 2</el-menu-item>
              </el-menu-item-group>
              <el-menu-item-group title="Group 2">
                <el-menu-item index="2-3">Option 3</el-menu-item>
              </el-menu-item-group>
              <el-sub-menu index="2-4">
                <template #title>Option 4</template>
                <el-menu-item index="2-4-1">Option 4-1</el-menu-item>
              </el-sub-menu>
            </el-sub-menu>
            <el-sub-menu index="3">
              <template #title>
                <el-icon>
                  <setting />
                </el-icon>
                Navigator Three
              </template>
              <el-menu-item-group>
                <template #title>Group 1</template>
                <el-menu-item index="3-1">Option 1</el-menu-item>
                <el-menu-item index="3-2">Option 2</el-menu-item>
              </el-menu-item-group>
              <el-menu-item-group title="Group 2">
                <el-menu-item index="3-3">Option 3</el-menu-item>
              </el-menu-item-group>
              <el-sub-menu index="3-4">
                <template #title>Option 4</template>
                <el-menu-item index="3-4-1">Option 4-1</el-menu-item>
              </el-sub-menu>
            </el-sub-menu>
          </el-menu>
        </el-scrollbar>
      </el-aside>
    </div>
    <el-container>
      <el-header class="flex items-center">Header</el-header>
      <el-main class="h-full">
        <el-button type="primary">Button</el-button>
        <router_view />
      </el-main>
      <el-footer class="flex items-center">Footer</el-footer>
    </el-container>
  </el-container>
</template>

<script setup lang="ts"></script>

<style scoped lang="scss">
:deep(.el-menu) {
  border-right: none;
  background-color: var(--el-menu-bg-color);
  //  // 设置选中时子标题的颜色
  //  .is-active {
  //    & > .el-sub-menu__title {
  //      color: var(--left-menu-text-active-color) !important;
  //      background-color: var(--left-menu-bg-active-color) !important;
  //    }
  //  }
}
</style>

<!--<style scoped lang="scss">-->
<!--$namespace: 'el';-->

<!--$prefix-cls: '#{$namespace}-menu';-->

<!--// .is-active&#45;&#45;after {-->
<!--//   position: absolute;-->
<!--//   top: 0;-->
<!--//   right: 0;-->
<!--//   width: 4px;-->
<!--//   height: 100%;-->
<!--//   background-color: var(&#45;&#45;el-color-primary);-->
<!--//   content: '';-->
<!--// }-->

<!--.#{$prefix-cls} {-->
<!--  position: relative;-->
<!--  transition: width var(&#45;&#45;transition-time-02);-->

<!--  // &:after {-->
<!--  //   position: absolute;-->
<!--  //   top: 0;-->
<!--  //   right: 0;-->
<!--  //   height: 100%;-->
<!--  //   width: 1px;-->
<!--  //   background-color: var(&#45;&#45;el-border-color);-->
<!--  //   content: '';-->
<!--  // }-->

<!--  :deep(.#{$prefix-cls}) {-->
<!--    width: 100% !important;-->
<!--    border-right: none;-->

<!--    // 设置选中时子标题的颜色-->
<!--    .is-active {-->
<!--      & > .#{$namespace}-sub-menu__title {-->
<!--        color: var(&#45;&#45;left-menu-text-active-color) !important;-->
<!--        // background-color: var(&#45;&#45;left-menu-bg-color) !important;-->
<!--      }-->
<!--    }-->

<!--    // 设置子菜单悬停的高亮和背景色-->
<!--    .#{$namespace}-sub-menu__title,-->
<!--    .#{$namespace}-menu-item {-->
<!--      &:hover {-->
<!--        color: var(&#45;&#45;left-menu-text-active-color) !important;-->
<!--        background-color: var(&#45;&#45;left-menu-bg-color) !important;-->
<!--      }-->
<!--    }-->

<!--    // 设置选中时的高亮背景和高亮颜色-->
<!--    .#{$namespace}-menu-item.is-active {-->
<!--      color: var(&#45;&#45;left-menu-text-active-color) !important;-->
<!--      background-color: var(&#45;&#45;left-menu-bg-active-color) !important;-->

<!--      &:hover {-->
<!--        background-color: var(&#45;&#45;left-menu-bg-active-color) !important;-->
<!--      }-->
<!--    }-->

<!--    .#{$namespace}-menu-item.is-active {-->
<!--      position: relative;-->

<!--      // &:after {-->
<!--      //   .is-active&#45;&#45;after;-->
<!--      // }-->
<!--    }-->

<!--    // 设置子菜单的背景颜色-->
<!--    .#{$namespace}-menu {-->
<!--      .#{$namespace}-sub-menu__title,-->
<!--      .#{$namespace}-menu-item:not(.is-active) {-->
<!--        background-color: var(&#45;&#45;left-menu-bg-light-color) !important;-->
<!--      }-->
<!--    }-->
<!--  }-->

<!--  // 折叠时的最小宽度-->
<!--  :deep(.#{$namespace}-menu&#45;&#45;collapse) {-->
<!--    width: var(&#45;&#45;left-menu-min-width);-->

<!--    & > .is-active,-->
<!--    & > .is-active > .#{$namespace}-sub-menu__title {-->
<!--      position: relative;-->
<!--      background-color: var(&#45;&#45;left-menu-collapse-bg-active-color) !important;-->

<!--      // &:after {-->
<!--      //   .is-active&#45;&#45;after;-->
<!--      // }-->
<!--    }-->
<!--  }-->

<!--  // 折叠动画的时候，就需要把文字给隐藏掉-->
<!--  :deep(.horizontal-collapse-transition) {-->
<!--    // transition: 0s width ease-in-out, 0s padding-left ease-in-out, 0s padding-right ease-in-out !important;-->
<!--    .#{$prefix-cls}__title {-->
<!--      display: none;-->
<!--    }-->
<!--  }-->

<!--  // 水平菜单-->
<!--  &__horizontal {-->
<!--    height: calc(var(&#45;&#45;top-tool-height)) !important;-->

<!--    :deep(.#{$namespace}-menu&#45;&#45;horizontal) {-->
<!--      height: calc(var(&#45;&#45;top-tool-height));-->
<!--      border-bottom: none;-->
<!--      // 重新设置底部高亮颜色-->
<!--      & > .#{$namespace}-sub-menu.is-active {-->
<!--        .#{$namespace}-sub-menu__title {-->
<!--          border-bottom-color: var(&#45;&#45;el-color-primary) !important;-->
<!--        }-->
<!--      }-->

<!--      .#{$namespace}-menu-item.is-active {-->
<!--        position: relative;-->

<!--        &::after {-->
<!--          display: none !important;-->
<!--        }-->
<!--      }-->

<!--      .#{$prefix-cls}__title {-->
<!--        /* stylelint-disable-next-line */-->
<!--        max-height: calc(var(&#45;&#45;top-tool-height) - 2px) !important;-->
<!--        /* stylelint-disable-next-line */-->
<!--        line-height: calc(var(&#45;&#45;top-tool-height) - 2px);-->
<!--      }-->
<!--    }-->
<!--  }-->
<!--}-->
<!--</style>-->

<!--<style lang="scss">-->
<!--$prefix-cls: 'el-menu-popper';-->
<!--.#{$prefix-cls}&#45;&#45;vertical,-->
<!--.#{$prefix-cls}&#45;&#45;horizontal {-->
<!--  // 设置选中时子标题的颜色-->
<!--  .is-active {-->
<!--    & > .el-sub-menu__title {-->
<!--      color: var(&#45;&#45;left-menu-text-active-color) !important;-->
<!--      // background-color: var(&#45;&#45;left-menu-bg-color) !important;-->
<!--    }-->
<!--  }-->

<!--  // 设置子菜单悬停的高亮和背景色-->
<!--  .el-sub-menu__title,-->
<!--  .el-menu-item {-->
<!--    &:hover {-->
<!--      color: var(&#45;&#45;left-menu-text-active-color) !important;-->
<!--      background-color: var(&#45;&#45;left-menu-bg-color) !important;-->
<!--    }-->
<!--  }-->

<!--  // 设置选中时的高亮背景-->
<!--  .el-menu-item.is-active {-->
<!--    position: relative;-->
<!--    background-color: var(&#45;&#45;left-menu-bg-active-color) !important;-->

<!--    &:hover {-->
<!--      background-color: var(&#45;&#45;left-menu-bg-active-color) !important;-->
<!--    }-->

<!--    // &:after {-->
<!--    //   .is-active&#45;&#45;after;-->
<!--    // }-->
<!--  }-->
<!--}-->
<!--</style>-->
