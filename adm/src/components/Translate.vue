<template>
  <el-dropdown class="h-full">
    <flat-button>
      <el-icon :size="18">
        <i class="i-icon-park-outline:translate" />
      </el-icon>
    </flat-button>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item>
          <span>简体中文</span>
        </el-dropdown-item>
        <el-dropdown-item>
          <span>English</span>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts"></script>

<style scoped lang="scss"></style>
