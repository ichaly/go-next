<template>
  <el-dropdown class="h-full">
    <flat-button>
      <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png">
        user
      </el-avatar>
    </flat-button>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item>
          <i class="i-icon-park-outline:user" />
          <span>个人信息</span>
        </el-dropdown-item>
        <el-dropdown-item>
          <i class="i-icon-park-outline:edit" />
          <span>修改密码</span>
        </el-dropdown-item>
        <el-dropdown-item divided>
          <i class="i-icon-park-outline:help" />
          <span>帮助文档</span>
        </el-dropdown-item>
        <el-dropdown-item>
          <i class="i-icon-park-outline:tips" />
          <span>功能更新</span>
        </el-dropdown-item>
        <el-dropdown-item divided>
          <i class="i-icon-park-outline:lock" />
          <span>锁定屏幕</span>
        </el-dropdown-item>
        <el-dropdown-item @click="$router.push('/login')">
          <i class="i-icon-park-outline:power" />
          <span>退出登录</span>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts"></script>

<style scoped lang="scss"></style>
