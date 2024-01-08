<template>
  <div class="h-full w-full">
    Dashboard
    <div>
      <el-button @click="loadMenus">加载菜单</el-button>
      <el-button @click="$router.push('/adduser')">添加用户</el-button>
      <div v-t="'dashboard.message'"></div>
      <div v-t="{ path: 'dashboard.say', args: { name: 'Lisa' } }"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
const { loadMenus } = useRootStore()
</script>

<style scoped lang="scss"></style>
