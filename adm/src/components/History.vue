<template>
  <div class="router-history">
    <el-tabs
      type="card"
      :closable="closable"
      v-model="$route.path"
      @tab-change="changeTab"
      @tab-remove="removeTab"
    >
      <el-tab-pane v-for="{ name, title } in histories" :key="name" :name="name" :label="title">
        <template #label>
          <el-dropdown
            :id="name"
            ref="dropdownRef"
            trigger="contextmenu"
            placement="bottom-start"
            @visible-change="onVisibleChange($event, name)"
          >
            <span class="label">
              <i class="dot" />
              {{ title }}
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="removeTab(name)">
                  <el-icon :size="18">
                    <i class="i-icon-park-outline:close" />
                  </el-icon>
                  关闭当前标签页
                </el-dropdown-item>
                <el-dropdown-item
                  @click="removeTab(name, 'left')"
                  v-if="isFirstOrLast(name, 'left')"
                >
                  <el-icon :size="18">
                    <i class="i-icon-park-outline:to-left" />
                  </el-icon>
                  关闭左侧标签页
                </el-dropdown-item>
                <el-dropdown-item
                  @click="removeTab(name, 'right')"
                  v-if="isFirstOrLast(name, 'right')"
                >
                  <el-icon :size="18">
                    <i class="i-icon-park-outline:to-right" />
                  </el-icon>
                  关闭右侧标签页
                </el-dropdown-item>
                <el-dropdown-item @click="removeTab(name, 'other')" v-if="size > 1">
                  <el-icon :size="18">
                    <i class="i-icon-park-outline:off-screen-two" />
                  </el-icon>
                  关闭其他标签页
                </el-dropdown-item>
                <el-dropdown-item @click="removeTab(name, 'all')">
                  <el-icon :size="18">
                    <i class="i-icon-park-outline:full-screen-two" />
                  </el-icon>
                  关闭全部标签页
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import type { TabPaneName } from 'element-plus'
import type { RouteMeta } from 'vue-router'
import Sortable from 'sortablejs'

const route = useRoute()
const router = useRouter()

const histories = useStorage<RouteMeta[]>('router-history', [], sessionStorage)

const size = computed(() => histories.value.length)
const closable = computed(() => !(histories.value.length === 1 && route.meta.default))

// 路由切换
const changeTab = (tab: TabPaneName) => {
  router.push(tab as string)
}
//删除路由
const removeTab = (tab: TabPaneName, type: string = 'self') => {
  const index = histories.value.findIndex((item) => tab === item.name)

  //使用策略模式执行删除逻辑
  const removeStrategies: Record<string, () => void> = {
    all: () => (histories.value = []),
    self: () => histories.value.splice(index, 1),
    left: () => histories.value.splice(0, index),
    right: () => histories.value.splice(index + 1),
    other: () => (histories.value = [histories.value[index]])
  }
  removeStrategies[type]?.()

  //计算当前的路由是否也被删除了是否需要路由跳转
  const current = histories.value.findIndex((item) => route.path === item.name)
  if (current < 0) {
    router.push({ force: true, path: histories.value.length === 0 ? '/' : histories.value[0].name })
  }
}
//控制右键菜单
const dropdownRef = ref()
const onVisibleChange = (visible: boolean, name: string) => {
  if (!visible) return
  dropdownRef.value.forEach((item: { id: string; handleClose: () => void }) => {
    if (item.id === name) return
    item.handleClose()
  })
}
const isFirstOrLast = (name: string, type: string) => {
  const index = histories.value.findIndex((item) => name === item.name)
  return type === 'left' ? index !== 0 : index !== size.value - 1
}
//监听路由变化并记录历史记录
watch(
  () => route.path,
  () => {
    const index = histories.value.findIndex((item) => route.path === item.name)
    if (index < 0) histories.value.push(route.meta)
  },
  { immediate: true }
)
onMounted(() => {
  //找到想要拖拽的那一列
  const el: HTMLElement | null = document.querySelector('.router-history .el-tabs__nav')
  Sortable.create(el!, {
    onEnd: ({ newIndex, oldIndex }) => {
      if (oldIndex == undefined || newIndex == undefined) {
        return
      }
      const item = histories.value.splice(oldIndex, 1)[0]
      histories.value.splice(newIndex, 0, item)
    }
  })
})
</script>

<style scoped lang="scss">
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

      &.is-active {
        @apply bg-blue-500 bg-opacity-5;
        .dot {
          background-color: var(--el-color-primary);
        }

        span {
          color: var(--el-color-primary);
        }
      }

      &:hover {
        span {
          color: var(--el-color-primary); //鼠标移到标签页高亮
        }
      }

      .el-dropdown {
        line-height: inherit; // 统一标签页显示名称行高
      }
    }

    .el-tabs__nav {
      @apply b-0;
    }
  }
}
</style>
