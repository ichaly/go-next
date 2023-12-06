<template>
  <div class='h-full w-full bg-red' @click='loadMenus'>
    Dashboard
  </div>
</template>

<script setup lang='ts'>
const { loadMenus } = useRootStore()
</script>


<style scoped lang='scss'></style>
