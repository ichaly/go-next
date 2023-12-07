<script setup lang="ts">
const router = useRouter()
const onClick = () => {
  router.push('/RoleManageradmin/Role123123')
}
</script>

<template>
  <div @click="onClick" class="h-full w-full bg-blue">Home Page</div>
</template>

<style scoped></style>
