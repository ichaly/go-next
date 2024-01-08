<template>
  <ul class="list">
    <li class="item" v-for="item in getData" :key="item.id">
      <el-avatar v-if="item.avatar" :size="30" class="avatar" :src="item.avatar" />
      <div class="content">
        <div class="title" @click="handleTitleClick(item)">
          <Ellipsis :content="item.title" class="!m-b-2">
            <span :class="{ '!line-through': item.titleDelete }">{{ item.title }}</span>
          </Ellipsis>
          <el-tag effect="light" v-if="item.extra" size="small" :type="item.color">
            {{ item.extra }}
          </el-tag>
        </div>
        <Ellipsis :content="item.description" :line="2" class="!m-b-1" v-if="item.description" />
        <span class="datetime">{{ item.datetime }}</span>
      </div>
    </li>
    <el-pagination
      layout="prev, pager, next"
      :total="list.length"
      :page-size="5"
      :pager-count="5"
      :current-page="current"
      @current-change="onCurrentChange"
      :hide-on-single-page="true"
    />
  </ul>
</template>

<script setup lang="ts">
import type { ListItem } from '@/components/Notify/data'
import isNumber from 'lodash/isNumber'

const props = withDefaults(
  defineProps<{
    list: ListItem[]
    pageSize?: boolean | number
    currentPage?: number
    onTitleClick?: (item: ListItem) => void
  }>(),
  {
    pageSize: 5,
    currentPage: 1
  }
)

const current: Ref<number> = ref(props.currentPage || 1)
const getData = computed((): ListItem[] => {
  const { pageSize, list } = props
  if (pageSize === false) return []
  let size: number = isNumber(pageSize) ? pageSize : 5
  return list.slice(size * (unref(current) - 1), size * unref(current))
})
const onCurrentChange = (page: number) => {
  current.value = page
}
const handleTitleClick = (item: ListItem) => {
  props.onTitleClick?.(item)
}
</script>

<style scoped lang="scss">
.list {
  @apply b b-rd b-light flex flex-col items-center;
  & .item {
    @apply w-full p-1.5 b-t b-light box-border flex flex-row cursor-pointer;
    &:first-child {
      @apply b-0;
    }

    & .avatar {
      @apply m-r w-7.5 h-7.5 bg-white;
    }

    & .content {
      @apply flex-1 flex flex-col overflow-hidden text-black;

      & .title {
        @apply flex flex-row;
      }

      & .datetime {
        @apply text-xs text-slate-400;
      }
    }
  }
}
</style>
