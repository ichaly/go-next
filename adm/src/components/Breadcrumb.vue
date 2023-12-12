<template>
  <el-breadcrumb class="hidden md:block" :separator-icon="ArrowRight">
    <el-breadcrumb-item :to="{ path: '/' }" v-for="item in matched">
      {{ item.meta.title }}
    </el-breadcrumb-item>
    <!--    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>-->
    <!--    <el-breadcrumb-item>-->
    <!--      <router-link to="/adm/public">用户管理</router-link>-->
    <!--    </el-breadcrumb-item>-->
    <!--    <el-breadcrumb-item>用户详情</el-breadcrumb-item>-->
  </el-breadcrumb>
</template>

<script setup lang="ts">
import ArrowRight from '@/components/ArrowRight.vue'

// const route = useRoute()
// const { menus } = toRefs(useRootStore())
// const parse = (items: Item[]) => {
//   for (let item of items) {
//     if (item.children) {
//       return [...parse(item.children)]
//     } else {
//       if (item.name === route.path) {
//         return [{ path: item.name, title: item.title }]
//       }
//     }
//   }
// }
</script>