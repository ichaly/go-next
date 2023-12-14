<template>
  <template v-for='item in menus' :key='item.path'>
    <el-sub-menu :index='item.name' v-if='item.children?.length'>
      <template #title>
        <el-icon><i :class='item.icon' /></el-icon>
        <span>{{ item.title }}</span>
      </template>
      <MenuTree :menus='item.children' />
    </el-sub-menu>
    <el-menu-item :index='item.name' v-else>
      <el-icon><i :class='item.icon' /></el-icon>
      <span>{{ item.title }}</span>
    </el-menu-item>
  </template>
</template>
<script setup lang='ts'>
defineProps<{ menus?: Item[] }>()
</script>
