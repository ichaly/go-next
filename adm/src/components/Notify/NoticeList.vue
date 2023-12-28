<template>
  <ul class="b b-rd b-solid b-[var(--el-border-color-light)]">
    <li class="h-20" v-for="item in getData" :key="item.id">

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

</style>