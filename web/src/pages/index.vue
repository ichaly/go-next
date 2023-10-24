<template>
  <div>
    <div class="box">
      <ElButton @click="ElMessage('hello')">Hello {{ store.username }}!</ElButton>
    </div>
    <NuxtWelcome/>
  </div>
</template>
<script setup>
const store = mainStore()
</script>
<style lang="scss">
.box {
  position: fixed;
}
</style>