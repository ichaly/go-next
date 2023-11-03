<template>
  <div>
    <!--    <header>-->
    <!--      <h1>欢迎来到易师傅的列表页面</h1>-->
    <!--    </header>-->
    <!--    <div class="fixed flex w-full bg-emerald-400 p-1 justify-end">-->
    <!--      <ElButton type="primary" @click="changeLang">{{$t("welcome")}} {{ username }} {{ count }}!</ElButton>-->
    <!--    </div>-->
    <header class="bg-nav">
      <el-menu mode="horizontal" :ellipsis="false" class="bg-transparent">
        <el-menu-item index="0">
          <img class="w-24" src="@/assets/images/element-plus-logo.svg" alt="Element logo"/>
        </el-menu-item>
        <div class="flex-grow"/>
        <el-menu-item index="1">Processing Center</el-menu-item>
        <el-sub-menu index="2">
          <template #title>Workspace</template>
          <el-menu-item index="2-1">item one</el-menu-item>
          <el-menu-item index="2-2">item two</el-menu-item>
          <el-menu-item index="2-3">item three</el-menu-item>
          <el-sub-menu index="2-4">
            <template #title>item four</template>
            <el-menu-item index="2-4-1">item one</el-menu-item>
            <el-menu-item index="2-4-2">item two</el-menu-item>
            <el-menu-item index="2-4-3">item three</el-menu-item>
          </el-sub-menu>
        </el-sub-menu>
      </el-menu>
    </header>
    <NuxtWelcome/>
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
<style lang="scss" scoped>
.bg-nav {
  @apply fixed w-full z-50 bg-[size:4px_4px];
  background-color: transparent;
  backdrop-filter: saturate(50%) blur(4px);
  background-image: radial-gradient(transparent 1px, white 1px);
}
</style>