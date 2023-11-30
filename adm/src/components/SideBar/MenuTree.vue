<template>
  <template v-for="item in list" :key="item.path">
    <el-sub-menu :index="item.path" v-if="item.children?.length">
      <template #title>
        <el-icon><i :class="item.icon" /></el-icon>
        <span>{{ item.name }}</span>
      </template>
      <MenuTree :list="item.children" />
    </el-sub-menu>
    <el-menu-item :index="item.path" v-else>
      <el-icon><i :class="item.icon" /></el-icon>
      <span>{{ item.name }}</span>
    </el-menu-item>
  </template>
</template>
<script setup lang="ts">
type Node = {
  name: string
  path: string
  icon?: string
  children?: Node[]
}
withDefaults(defineProps<{ list?: Node[] }>(), {
  list: () => [
    {
      name: 'Navigator One',
      path: '/index1',
      icon: 'i-ep:location',
      children: [
        {
          name: 'item one',
          path: '/index1/sub1'
        },
        {
          name: 'item two',
          path: '/index1/sub2'
        },
        {
          name: 'item three',
          path: '/index1/sub3'
        },
        {
          name: 'item four',
          path: '/index1/sub4',
          children: [
            {
              name: 'item on',
              path: '/index1/sub4/sub1'
            }
          ]
        }
      ]
    },
    {
      name: 'Navigator Two',
      path: '/index2',
      icon: 'i-ep:menu'
    },
    {
      name: 'Navigator Three',
      path: 'http://baidu.com',
      icon: 'i-ep:document'
    },
    {
      name: 'Navigator Four',
      path: '/home',
      icon: 'i-ep:setting'
    }
  ]
})
</script>

<style scoped lang="ts"></style>
