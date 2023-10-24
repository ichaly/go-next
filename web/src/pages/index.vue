<template>
  <div>
    <div class="fixed flex w-full bg-emerald-400 p-1 justify-end">
      <ElButton type="primary" @click="changeLang">{{$t("welcome")}} {{ username }} {{ count }}!</ElButton>
    </div>
    <NuxtWelcome/>
  </div>
</template>
<script setup>
const store = useMainStore()
const {increment} = store
const {count, username} = storeToRefs(store)

const {locale, locales, setLocale} = useI18n()
const changeLang = () => {
  setLocale('en')
};
</script>
<style lang="scss">
.box {
}
</style>