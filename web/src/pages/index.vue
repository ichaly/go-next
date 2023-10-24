<template>
  <div>
    <div class="box">
      <ElButton @click="increment">Hello {{ username }} {{ count }}!</ElButton>
    </div>
    <NuxtWelcome/>
  </div>
</template>
<script setup>
const store = useMainStore()
const {increment} = store
const {count, username} = storeToRefs(store)
</script>
<style lang="scss">
.box {
  position: fixed;
}
</style>