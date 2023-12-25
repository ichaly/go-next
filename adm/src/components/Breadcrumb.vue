<template>
  <el-breadcrumb class="hidden md:block" :separator-icon="IconArrowRight">
    <el-breadcrumb-item>首页</el-breadcrumb-item>
    <el-breadcrumb-item v-for="item in $route.meta.items" :key="item.name">
      {{ item.title }}
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script setup lang="ts">
import IconArrowRight from '@/components/icons/IconArrowRight.vue'
</script>
