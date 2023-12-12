<template>
  <div class="h-full w-full bg-red">
    Dashboard
    <div>
      <el-button @click="loadMenus">加载菜单</el-button>
      <el-button @click="onClick">角色管理</el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
const { loadMenus } = useRootStore()
const router = useRouter()
const onClick = () => {
  router.push('/RoleManageradmin/Role123123')
}
</script>

<style scoped lang="scss"></style>
