<template>
  <el-breadcrumb class="hidden md:block" :separator-icon="ArrowRight">
    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>
      <router-link to="/adm/public">用户管理</router-link>
    </el-breadcrumb-item>
    <el-breadcrumb-item>用户详情</el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script setup lang="ts">
import ArrowRight from '@/components/ArrowRight.vue'
</script>