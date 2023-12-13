<template>
  <el-breadcrumb class='hidden md:block' :separator-icon='ArrowRight'>
    <el-breadcrumb-item v-for='item in matched'>
      {{ item.title }}
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script setup lang='ts'>
import ArrowRight from '@/components/ArrowRight.vue'

const route = useRoute()
const { menus } = toRefs(useRootStore())
const search = (items: Item[]): Item[] => {
  for (let item of items) {
    if (`/${item.name}` === route.path) {
      return [item]
    } else if (item.children?.length) {
      let nodes = search(item.children)
      if (nodes.length) {
        return [item, ...nodes]
      }
    }
  }
  return []
}
const matched = computed(() => {
  return [{ title: '首页' }, ...search(menus.value)]
})
</script>