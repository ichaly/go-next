<template>
  <div class='router-history'>
    <el-tabs
      type='card'
      :closable='closable'
      v-model='$route.path'
      @tab-change='changeTab'
      @tab-remove='removeTab'
      @contextmenu.prevent='openContextMenu($event)'
    >
      <el-tab-pane v-for='({meta:{ title }},name) in histories' :key='name' :name='name' :label='title'>
        <template #label>
          <span>
            <i class='dot' />
            {{ title }}
          </span>
        </template>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang='ts'>
import type { TabPaneName } from 'element-plus'
import type { RouteLocationNormalizedLoaded } from 'vue-router'

const route = useRoute()
const router = useRouter()
const histories: Ref<Record<string, RouteLocationNormalizedLoaded>> = ref({})
const closable = computed(() => !((Object.keys(histories.value)).length === 1 && route.meta.default))
// 路由切换
const changeTab = (tab: TabPaneName) => {
  router.push(histories.value[tab])
}
//删除路由
const removeTab = (tab: TabPaneName) => {
  delete (histories.value[tab])
  const keys = Object.keys(histories.value)
  if (keys.length === 0) {
    router.push('/')
  } else {
    router.push(keys[0])
  }
}
const openContextMenu = (e: any) => {
}
watchEffect(() => {
  histories.value[route.path] = { ...route }
})
</script>

<style scoped lang='scss'>
.el-tabs__item .dot {
  content: '';
  width: 9px;
  height: 9px;
  margin-right: 8px;
  display: inline-block;
  border-radius: 50%;
  transition: background-color 0.2s;
}

.router-history {
  @apply bg-white p-0 border-t border-l-0 border-r-0 border-b-0 border-solid border-gray-100;
  :deep(.el-tabs__header) {
    @apply m-0;
    .el-tabs__item {
      @apply border-solid border-r border-t-0 border-gray-100 border-b-0 border-l-0;
      .dot {
        content: '';
        width: 9px;
        height: 9px;
        margin-right: 8px;
        display: inline-block;
        border-radius: 50%;
        transition: background-color 0.2s;
        background-color: #ddd;
      }
    }

    .el-tabs__item {
      &.is-active {
        @apply bg-blue-500 bg-opacity-5;
        .dot {
          background-color: var(--el-color-primary);
        }

        span {
          color: var(--el-color-primary);
        }
      }
    }

    .el-tabs__nav {
      @apply b-0;
    }
  }
}
</style>
