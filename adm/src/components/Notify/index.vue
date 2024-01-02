<template>
  <el-popover :visible="true" placement="bottom" :width="260" trigger="hover">
    <template #reference>
      <flat-button>
        <el-badge :value="200" :max="99" class="h-4.5">
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
          <span v-if="item.list.length !== 0">({{ item.list.length }})</span>
        </template>
        <!-- 绑定title-click事件的通知列表中标题是“可点击”的-->
        <NoticeList :list="item.list" v-if="item.key === '1'" @title-click="onNoticeClick" />
        <NoticeList :list="item.list" v-else />
      </el-tab-pane>
    </el-tabs>
  </el-popover>
</template>

<script setup lang="ts">
import { type ListItem, tabListData } from './data'
import { ElMessage } from 'element-plus'

const activeName = ref('1')
const listData = ref(tabListData)

const onNoticeClick = (record: ListItem) => {
  ElMessage({
    message: '你点击了通知，ID=' + record.id,
    type: 'success'
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
