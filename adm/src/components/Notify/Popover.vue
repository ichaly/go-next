<template>
  <el-popover
    placement="bottom"
    title="Title"
    :width="200"
    trigger="hover"
    content="this is content, this is content, this is content"
  >
    <template #reference>
      <flat-button>
        <el-badge :value="200" :max="99" class="h-4.5">
          <el-icon :size="18">
            <i class="i-icon-park-outline:remind" />
          </el-icon>
        </el-badge>
      </flat-button>
    </template>
  </el-popover>
</template>

<script setup lang="ts">

</script>

<style scoped lang="scss">

</style>