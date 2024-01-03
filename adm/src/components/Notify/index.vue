<template>
  <el-popover placement="bottom" :width="280" trigger="hover">
    <template #reference>
      <flat-button>
        <el-badge :value="total" :max="99" class="h-4.5">
          <el-icon :size="18">
            <i class="i-icon-park-outline:remind" />
          </el-icon>
        </el-badge>
      </flat-button>
    </template>
    <el-tabs v-model="activeName">
      <el-tab-pane :name="item.key" v-for="item in listData" :key="item.key">
        <template #label>
          {{ item.name }}
          <span v-if="item.list.length !== 0">{{ getTotal(item.list) }}</span>
        </template>
        <NoticeList :list="item.list" v-if="item.key === '1'" @title-click="onNoticeClick" />
        <NoticeList :list="item.list" v-else />
      </el-tab-pane>
    </el-tabs>
  </el-popover>
</template>

<script setup lang="ts">
import { type ListItem, tabListData } from './data'

const activeName = ref('1')
const listData = ref(tabListData)
const total = computed(() => {
  return listData.value.reduce((pre, cur) => {
    return pre + cur.list.filter((item) => !item.titleDelete).length
  }, 0)
})
const getTotal = computed(() => (list: ListItem[]) => {
  const length = list.filter((item) => !item.titleDelete).length
  return length ? `(${length})` : ''
})
const onNoticeClick = (record: ListItem) => {
  ElMessage.success({
    message: '你点击了通知，ID=' + record.id
  })
  // 可以直接将其标记为已读（为标题添加删除线）,此处演示的代码会切换删除线状态
  record.titleDelete = !record.titleDelete
}
</script>

<style scoped lang="scss">
:deep(.el-tabs__nav-wrap) {
  .el-tabs__nav {
    justify-content: space-between;
    width: 100%;
  }

  &::after {
    height: 1px;
  }
}
</style>
