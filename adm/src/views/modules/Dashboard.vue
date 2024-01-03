<template>
  <div class="h-full w-full">
    Dashboard
    <div>
      <el-button @click="loadMenus">加载菜单</el-button>
      <el-button @click="$router.push('/adduser')">添加用户</el-button>
      <div>{{ $t('message') }}</div>
      <div>{{ $t('say', { name: 'Lisa' }) }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
const { loadMenus } = useRootStore()
</script>

<style scoped lang="scss"></style>
