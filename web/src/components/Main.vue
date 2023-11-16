<script setup lang="ts">
type Link = {
  title: string
  href: string
  target?: string
}
type Feature = {
  title: string
  description: string
  icon: string
  link?: Link
}
const features = ref<Feature[]>([
  {
    title: "内置权限管理",
    description: "内置完善的rbac权限管理体系，开箱即用。",
    icon: "i-ri:account-pin-circle-line",
    link: {
      title: "了解一下",
      href: "https://casbin.org/zh/docs/overview/"
    }
  },
  {
    title: "多语言支持",
    description: "支持多种不同语言，可以自己添加语言包，加载语言包针对性翻译支持。",
    icon: "i-ri:earth-line"
  },
  {
    title: "支持多个框架",
    description: "支持包括像gin，echo，fasthttp，chi，beego等流行的web框架无缝接入，选用您擅长的框架。",
    icon: "i-ri:checkbox-multiple-blank-line",
    link: {
      title: "快速开始",
      href: "https://casbin.org/zh/docs/overview/"
    }
  },
  {
    title: "pjax无刷新体验",
    description: "系统内采用pjax技术，实现无刷新的良好用户体验，不用再为页面刷新卡住而烦恼。",
    icon: "i-ri:refresh-line"
  },
  {
    title: "插件提升效率",
    description: "不同的插件可以简化你的开发工作量，插件的接入十分简易。",
    icon: "i-ri:plug-line",
    link: {
      title: "马上尝试",
      href: "https://casbin.org/zh/docs/overview/"
    }
  },
  {
    title: "编译成二进制",
    description: "有了golang的支持，整个系统可以编译成一个二进制文件，便于部署与分发。",
    icon: "i-ri:file-list-line"
  },
  {
    title: "不同UI主题",
    description: "将提供不同的ui主题模板以适应您个性化的需求。",
    icon: "i-ri:brush-2-line",
    link: {
      title: "帮我设计",
      href: "https://casbin.org/zh/docs/overview/"
    }
  },
  {
    title: "性能更好",
    description: "得益于golang的良好性能，GoAdmin也同步拥有优于其他语言同类框架的性能特性。",
    icon: "i-ri:dashboard-3-line"
  },
])
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
      <div class="w-full flex flex-row flex-wrap content-start gap-4 box-border">
        <div class="grow-1 shrink-0 basis-1/4" v-for='(f,i) in features' :key='i'>
          <div
              class="flex flex-col p-8 bg-slate-100 dark:bg-[#202127] dark:border-[#202127] rounded-xl h-full cursor-pointer border-1 hover:border-[--el-color-primary]">
            <div class="w-12 h-12 rounded-md flex items-center justify-center bg-slate-300 mb-5 dark:bg-slate-600">
              <i :class="`${f.icon} block text-2xl text-slate-800 dark:text-slate-100`"/>
            </div>
            <h2 class="leading-6 text-base font-semibold text-base">{{ f.title }}</h2>
            <p class="grow pt-2 text-sm text-slate-500 pt-2 flex-1">{{ f.description }}</p>
            <p class="flex items-center text-sm font-medium mt-2 text-[--el-color-primary]" v-if="f.link">
              {{f.link?.title}}
              <i class="i-ep:right ml-1.5"/>
            </p>
          </div>
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