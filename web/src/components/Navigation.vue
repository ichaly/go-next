<script setup lang="ts">
type Node = {
  name: string,
  path: string,
  icon?: string,
  children?: Node[]
}
withDefaults(defineProps<{ list?: Node[] }>(), {
  list: () => [
    {
      name: '首页',
      path: '/index',
    },
    {
      name: '文档',
      path: '/docs',
      children: [
        {
          name: '指南',
          path: '/index/sub1',
          icon: 'i-ri:compass-3-line',
        },
        {
          name: 'API',
          path: '/index/sub2',
          icon: 'i-ant-design:api-outlined',
        },
      ],
    },
  ],
})
</script>

<template>
  <template v-for="item in list" :key="item.path">
    <el-sub-menu :index="item.path" v-if="item.children??[].length>0">
      <template #title>
        <el-icon><i :class="item.icon"/></el-icon>
        <span>{{ item.name }}</span>
      </template>
      <Navigation :list="item.children"/>
    </el-sub-menu>
    <el-menu-item :index="item.path" v-else>
      <el-icon><i :class="item.icon"/></el-icon>
      <span>{{ item.name }}</span>
    </el-menu-item>
  </template>
</template>

<style scoped lang="ts">

</style>