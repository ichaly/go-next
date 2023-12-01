<template>
  <div class="error-content">
    <div class="pic-error">
      <img :alt="type" class="pic-error-parent" :src="src" />
      <img :alt="type" class="pic-error-left" src="@/assets/images/cloud.png" />
      <img :alt="type" class="pic-error-mid" src="@/assets/images/cloud.png" />
      <img :alt="type" class="pic-error-right" src="@/assets/images/cloud.png" />
    </div>
    <div class="bullshit">
      <div class="bullshit-oops">抱歉!</div>
      <div class="bullshit-headline">{{ title }}</div>
      <div class="bullshit-info">{{ description }}</div>
      <el-button round size="large" type="primary" @click="onBack()" class="bullshit-button">
        返回首页（{{ time }}s）
      </el-button>
    </div>
  </div>
</template>
<script setup lang="ts">
type Prop = {
  type: '403' | '404'
  title: string
  duration?: number
  description: string
}
const { duration, type } = withDefaults(defineProps<Prop>(), {
  duration: 500
})

const src = new URL(`../assets/images/${type}.svg`, import.meta.url).href

const router = useRouter()
const onBack = () => {
  router.push('/')
}

const time = ref(duration)
const { pause, resume } = useIntervalFn(
  () => {
    if (time.value > 1) {
      time.value--
    } else {
      pause()
      onBack()
    }
  },
  1000,
  { immediate: false }
)
resume()
</script>

<style scoped lang="scss">
.error-content {
  @apply w-full h-full flex center flex-col lg:flex-row;
  .pic-error {
    @apply w-80 h-80 relative;
    overflow: hidden;

    & > img {
      position: absolute;
    }

    &-parent {
      width: 100%;
      height: 100%;
    }

    &-left {
      top: 17px;
      left: 220px;
      width: 80px;
      opacity: 0;
      animation-name: cloudLeft;
      animation-duration: 2s;
      animation-timing-function: linear;
      animation-delay: 1s;
      animation-fill-mode: forwards;
    }

    &-mid {
      top: 10px;
      left: 420px;
      width: 46px;
      opacity: 0;
      animation-name: cloudMid;
      animation-duration: 2s;
      animation-timing-function: linear;
      animation-delay: 1.2s;
      animation-fill-mode: forwards;
    }

    &-right {
      top: 100px;
      left: 500px;
      width: 62px;
      opacity: 0;
      animation-name: cloudRight;
      animation-duration: 2s;
      animation-timing-function: linear;
      animation-delay: 1s;
      animation-fill-mode: forwards;
    }

    @keyframes cloudLeft {
      0% {
        top: 17px;
        left: 220px;
        opacity: 0;
      }

      20% {
        top: 33px;
        left: 188px;
        opacity: 1;
      }

      80% {
        top: 81px;
        left: 92px;
        opacity: 1;
      }

      100% {
        top: 97px;
        left: 60px;
        opacity: 0;
      }
    }

    @keyframes cloudMid {
      0% {
        top: 10px;
        left: 420px;
        opacity: 0;
      }

      20% {
        top: 40px;
        left: 360px;
        opacity: 1;
      }

      70% {
        top: 130px;
        left: 180px;
        opacity: 1;
      }

      100% {
        top: 160px;
        left: 120px;
        opacity: 0;
      }
    }

    @keyframes cloudRight {
      0% {
        top: 100px;
        left: 500px;
        opacity: 0;
      }

      20% {
        top: 120px;
        left: 460px;
        opacity: 1;
      }

      80% {
        top: 180px;
        left: 340px;
        opacity: 1;
      }

      100% {
        top: 220px;
        left: 240px;
        opacity: 0;
      }
    }
  }

  .bullshit {
    position: relative;
    float: left;
    width: 300px;
    padding: 30px 0;
    overflow: hidden;

    &-oops {
      margin-bottom: 20px;
      font-size: 32px;
      font-weight: bold;
      line-height: 40px;
      color: var(--el-color-primary);
      opacity: 0;
      animation-name: slideUp;
      animation-duration: 0.5s;
      animation-fill-mode: forwards;
    }

    &-headline {
      margin-bottom: 10px;
      font-size: 20px;
      font-weight: bold;
      line-height: 24px;
      color: #222;
      opacity: 0;
      animation-name: slideUp;
      animation-duration: 0.5s;
      animation-delay: 0.1s;
      animation-fill-mode: forwards;
    }

    &-info {
      margin-bottom: 30px;
      font-size: 13px;
      line-height: 21px;
      color: var(--el-color-info);
      opacity: 0;
      animation-name: slideUp;
      animation-duration: 0.5s;
      animation-delay: 0.2s;
      animation-fill-mode: forwards;
    }

    &-button {
      opacity: 0;
      animation-name: slideUp;
      animation-duration: 0.5s;
      animation-delay: 0.3s;
      animation-fill-mode: forwards;
    }

    @keyframes slideUp {
      0% {
        opacity: 0;
        transform: translateY(60px);
      }

      100% {
        opacity: 1;
        transform: translateY(0);
      }
    }
  }
}
</style>