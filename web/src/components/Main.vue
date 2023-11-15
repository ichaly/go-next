<script setup lang="ts">

</script>

<template>
  <Container>
    <main class="w-full min-h-screen">
      <div class="flex flex-row px-16 py-32">
        <div class="main text flex-1">
          <h1>
            <span class="clip bg-rainbow">GoNext</span>
          </h1>
          <p class="text-[56px]">基于Golang + Nuxt的中后台系统脚手架</p>
          <p class="text-6 leading-9 text-[#3c3c43]/78 pt-3">简单·强大·高效</p>
        </div>
        <div class="image flex-1 flex justify-center items-center">
          <div class="w-80 h-80 bg-red:10 absolute rounded-full bg-rainbow z-1 blur-3xl"></div>
          <img src="~/assets/images/go-next-logo.svg" alt="GoNext" class="w-55 h-55 z-2 pointer-events-none">
        </div>
      </div>
    </main>
  </Container>
</template>

<style scoped lang="scss">
@import url(~/assets/styles/rainbow.scss);

.text {
  font-size: 56px;
  line-height: 64px;
  font-weight: 700;
  white-space: pre-wrap;
}

.clip {
  @apply line-height-16 bg-clip-text  text-transparent antialiased;
}

.bg-rainbow {
  @apply bg-gradient-to-br from-indigo-500 from-30% to-pink-500;
}

.image {
  .bg-rainbow {
    &:hover {
      @apply opacity-75;
    }
  }
}

:root {
  animation: rainbow 12s linear infinite;
}
</style>