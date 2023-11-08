<script setup lang="ts">
const route = useRoute()
const {data: page} = await useAsyncData(route.path, () => queryContent(route.path).findOne())
if (!page.value) {
  throw createError({statusCode: 404, statusMessage: 'Page not found', fatal: true})
}
</script>

<template>
  <ContentRenderer v-if="page?.body" :value="page"/>
</template>

<style scoped lang="ts">

</style>