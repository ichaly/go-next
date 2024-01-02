<template>
  <ul class="list">
    <li class="item" v-for="item in getData" :key="item.id">
      <el-avatar v-if="item.avatar" :size="30" class="avatar" :src="item.avatar" />
      <div class="content">
        <Ellipsis :content="item.title" class="m-b-2" />
        <Ellipsis :content="item.description" :line="2" class="m-b-1" />
        <span class="datetime">{{ item.datetime }}</span>
      </div>
    </li>
  </ul>
</template>

<script setup lang="ts">
import type { ListItem } from '@/components/Notify/data'

const props = defineProps({
  list: {
    type: Array as PropType<ListItem[]>,
    default: () => []
  },
  pageSize: {
    type: [Boolean, Number] as PropType<Boolean | Number>,
    default: 5
  },
  currentPage: {
    type: Number,
    default: 1
  },
  titleRows: {
    type: Number,
    default: 1
  },
  descRows: {
    type: Number,
    default: 2
  },
  onTitleClick: {}
})
const current = ref(props.currentPage || 1)

const getData = computed(() => {
  const { pageSize, list } = props
  if (pageSize === false) return []
  let size: number = 5
  return list.slice(size * (unref(current) - 1), size * unref(current))
})
</script>

<style scoped lang="scss">
.list {
  @apply b b-rd b-light;
  & .item {
    @apply w-full p-1.5 b-b b-light box-border flex flex-row cursor-pointer;
    &:last-child {
      @apply b-0;
    }

    & .avatar {
      @apply m-r w-7.5 h-7.5 bg-white;
    }

    & .content {
      @apply flex-1 flex flex-col overflow-hidden text-black;

      & .datetime {
        @apply text-xs text-slate-400;
      }
    }
  }
}
</style>
