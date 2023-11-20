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
    title: "原生支持多租户",
    description: "多租户设计简化了应用的开发和维护，使开发者能更专注于核心业务功能的实现。",
    icon: "i-ri:checkbox-multiple-blank-line",
    link: {
      title: "快速开始",
      href: "https://casbin.org/zh/docs/overview/"
    }
  },
  {
    title: "前后分离",
    description: "前后端分离的开发模式提高了代码的可维护性与扩展性，加速项目迭代与上线。",
    icon: "i-ri:pin-distance-line"
  },
  {
    title: "插件化开发",
    description: "插件化开发降低了耦合度，提高代码复用率，缩短开发周期，提升整体开发效率。",
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
    description: "得益于golang的良好性能，GoNext也同步拥有优于其他语言同类框架的性能特性。",
    icon: "i-ri:dashboard-3-line"
  },
])

const content = ref<string>("go get github.com/ichaly/go-next@latest")
const { copy, copied } = useClipboard()
</script>

<template>
  <Container>
    <main class="w-full">
      <div class="flex flex-row w-full pt-10">
        <div class="flex flex-col justify-center flex-1">
          <h1 class="text-5xl font-bold tracking-tight text-gray-900 text-center dark:text-white md:text-7xl md:text-start">
            <span>基于<span class="clip rainbow">GO</span>的<br class="md:hidden"/>现代化WEB开发库</span>
          </h1>
          <p class="mt-6 text-lg tracking-tight text-gray-600 text-center dark:text-gray-300 md:text-start">
            <span>融合 Golang、Nuxt 和 GraphQL 的强大中后台框架，拥有灵活的架构，能够轻松解决各种复杂任务，让你快速开发Web App的梦想变成现实。<br/><br/>快来体验一下吧！</span>
          </p>
          <div class="mt-10 flex gap-x-6 justify-start items-start">
            <el-button type="primary" size="large">
              <template #icon>
                <i class="i-heroicons-rocket-launch h-5 w-5 flex-shrink-0"/>
              </template>
              立即体验
            </el-button>
            <el-input size="large" :value="content" readonly>
              <template #prefix>
                <i class="i-heroicons-command-line h-5 w-5 flex-shrink-0"/>
              </template>
              <template #append>
                <el-button class="!p-0" @click="copy(content)">
                  <template #icon>
                    <i class="i-heroicons-clipboard-document-check h-5 w-5 flex-shrink-0 text-[--el-color-primary]" v-if="copied"/>
                    <i class="i-heroicons-clipboard-document h-5 w-5 flex-shrink-0" v-else/>
                  </template>
                </el-button>
              </template>
            </el-input>
          </div>
        </div>
        <div class="image flex-1 justify-center items-center hidden md:flex">
          <div class="rainbow w-80 h-80 absolute rounded-full z-1 blur-[140px]"></div>
          <img src="~/assets/images/go-next-logo.svg" alt="GoNext" class="w-55 h-55 z-2 pointer-events-none">
        </div>
      </div>
      <div class="w-full gap-4 py-20 grid grid-cols-1 md:grid-cols-3 xl:grid-cols-4">
        <div v-for='(f,i) in features' :key='i'>
          <div class="flex flex-col p-8 bg-slate-100 dark:bg-[#202127] dark:border-[#202127] rounded-xl h-full cursor-pointer border-1 hover:border-[--el-color-primary]">
            <div class="w-12 h-12 rounded-md flex items-center justify-center bg-slate-300 mb-5 dark:bg-slate-600">
              <i :class="`${f.icon} block text-2xl text-slate-800 dark:text-slate-100`"/>
            </div>
            <h2 class="leading-6 text-base font-semibold text-base">{{ f.title }}</h2>
            <p class="grow pt-2 text-sm text-slate-500 pt-2 flex-1">{{ f.description }}</p>
            <p class="flex items-center text-sm font-medium mt-2 text-[--el-color-primary]" v-if="f.link">
              {{ f.link?.title }}
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