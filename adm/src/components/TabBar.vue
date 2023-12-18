<template>
  <div class="router-history">
    <el-tabs
      type="card"
      v-model="activeTab"
      @tab-click="changeTab"
      @tab-remove="removeTab"
      @contextmenu.prevent="openContextMenu($event)"
      :closable="!(histories.length === 1 && $route.name === '')"
    >
      <el-tab-pane
        v-for="item in histories"
        :key="item.name"
        :label="item.title"
        :name="item.name"
        :tab="item"
        class="gva-tab"
      >
        <template #label>
          <span :tab="item" :style="{ color: '#333' }">
            <i class="dot" :style="{ backgroundColor: '#ddd' }" />
            {{ item.title }}
          </span>
        </template>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import type { TabPaneName } from 'element-plus'

const router = useRouter()
const histories: Ref<Item[]> = ref([
  {
    id: 1,
    pid: 0,
    type: 'menu',
    name: 'home',
    title: '首页'
  },
  {
    id: 2,
    pid: 0,
    type: 'menu',
    name: 'about',
    title: '关于'
  }
])
const activeTab = ref('')

const changeTab = (tab: TabPaneName) => {}
const removeTab = (tab: TabPaneName) => {}
const openContextMenu = (e: any) => {}
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
      }
    }

    .el-tabs__item {
      &.is-active {
        @apply bg-blue-500 bg-opacity-5;
      }
    }

    .el-tabs__nav {
      @apply b-0;
    }
  }
}
</style>