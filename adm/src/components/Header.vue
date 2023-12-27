<template>
  <div class="flex items-center h-full">
    <flat-button @click="toggleCollapse()">
      <el-icon :size="18" :class="['transition-transform', { '-scale-x-100': isCollapse }]">
        <i class="i-icon-park-outline:indent-left" />
      </el-icon>
    </flat-button>
    <div class="flex-1 pl-2">
      <Breadcrumb />
    </div>
    <Search />
    <flat-button @click="toggleFullscreen()">
      <el-icon :size="18">
        <i class="i-icon-park-outline:off-screen-one" v-if="isFullscreen" />
        <i class="i-icon-park-outline:full-screen-one" v-else />
      </el-icon>
    </flat-button>
    <Popover />
    <Translate />
    <Avatar />
  </div>
</template>

<script setup lang="ts">

const rootStore = useRootStore()
const { toggleCollapse, toggleFullscreen } = rootStore
const { isCollapse, isFullscreen } = toRefs(rootStore)
</script>

<style scoped lang="scss"></style>
