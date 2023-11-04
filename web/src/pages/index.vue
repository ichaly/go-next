<template>
  <div>
    <AppHeader/>
    <NuxtWelcome/>
    <NuxtWelcome/>
  </div>
</template>
<script setup>
const store = useMainStore()
const {increment} = store
const {count, username} = storeToRefs(store)
</script>
<style lang="scss" scoped>
</style>