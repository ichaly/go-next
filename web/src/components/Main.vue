<script setup lang="ts">

</script>

<template>
  <Container>
    <main class="w-full">
      <div class="py-20 flex flex-row w-full">
        <div class="flex flex-col justify-center items-start flex-1">
          <h1 class="text-5xl font-bold tracking-tight text-gray-900 dark:text-white sm:text-7xl mt-10">
            <span>基于<span class="clip rainbow">GO</span>的现代化WEB开发库</span>
<!--            A <span class="clip rainbow">Go Library</span> for Modern Web Apps-->
          </h1>
          <p class="mt-6 text-lg tracking-tight text-gray-600 dark:text-gray-300">
          <span class="mt-6 text-lg tracking-tight text-gray-600 dark:text-gray-300">
            融合了 Golang、Nuxt 和 GraphQL 的强大中后台框架，拥有灵活的架构，<br class="hidden lg:block"/>能够轻松解决各种复杂任务，让你快速开发Web App的梦想变成现实。
            <br/><br/>快来体验一下吧！
          </span>
          </p>
          <div class="mt-10 flex gap-x-6 w-140">
            <el-button type="primary" size="large">
              <template #icon>
                <i class="i-heroicons-rocket-launch h-5 w-5 flex-shrink-0"/>
              </template>
              立即体验
            </el-button>
            <el-input size="large" value="go get github.com/ichaly/go-next@latest" readonly class="w-32">
              <template #prefix>
                <i class="i-heroicons-command-line h-5 w-5 flex-shrink-0"/>
              </template>
              <template #append>
                <el-button class="!p-0">
                  <template #icon>
                    <i class="i-heroicons-clipboard-document h-5 w-5 flex-shrink-0"/>
                  </template>
                </el-button>
              </template>
            </el-input>
          </div>
        </div>
        <div class="image flex-1 flex justify-center items-center">
          <div class="rainbow w-80 h-80 bg-red:10 absolute rounded-full z-1 blur-[240px]"></div>
          <img src="~/assets/images/go-next-logo.svg" alt="GoNext" class="w-55 h-55 z-2 pointer-events-none">
        </div>
      </div>
      <div class="flex flex-wrap -m-2">
        <div class="p-8 bg-slate-50 rounded-xl">
          <div class="w-12 h-12 rounded-md flex items-center justify-center bg-slate-100 mb-5">
            <i class="i-carbon:ibm-toolchain block text-2xl text-slate"/>
          </div>
          <h2 class="leading-6 text-base font-semibold text-sm">Fully Customizable</h2>
          <p class="grow pt-2 text-sm text-slate-500 pt-2">No core utilities, all functionalities are provided via presets.</p>
        </div>
        <div class="p-8 bg-slate-50 rounded-xl">
          <div class="w-12 h-12 rounded-md flex items-center justify-center bg-slate-100 mb-5">
            <i class="i-carbon:ibm-toolchain block text-2xl text-slate"/>
          </div>
          <h2 class="leading-6 text-base font-semibold text-sm">Fully Customizable</h2>
          <p class="grow pt-2 text-sm text-slate-500 pt-2">No core utilities, all functionalities are provided via presets.</p>
        </div>
        <div class="p-8 bg-slate-50 rounded-xl">
          <div class="w-12 h-12 rounded-md flex items-center justify-center bg-slate-100 mb-5">
            <i class="i-carbon:ibm-toolchain block text-2xl text-slate"/>
          </div>
          <h2 class="leading-6 text-base font-semibold text-sm">Fully Customizable</h2>
          <p class="grow pt-2 text-sm text-slate-500 pt-2">No core utilities, all functionalities are provided via presets.</p>
        </div>
      </div>
    </main>
  </Container>
</template>

<style scoped lang="scss">
@import url(~/assets/styles/rainbow.scss);

.clip {
  @apply line-height-16 bg-clip-text text-transparent antialiased;
}

.rainbow {
  @apply bg-gradient-to-br from-indigo-500 from-30% to-pink-500;
}

.image {
  .rainbow {
    &:hover {
      @apply opacity-75;
    }
  }
}
</style>