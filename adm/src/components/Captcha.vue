<template>
  <MySlot />
</template>

<script setup lang="ts">
import { randomString } from '@/utils/strings'

interface Props {
  prefix?: string// 身份标。开通阿里云验证码2.0后，您可以在控制台概览页面的实例基本信息卡片区域，获取身份标
  sceneId?: string// 场景ID。根据步骤二新建验证场景后，您可以在验证码场景列表，获取该场景的场景ID
}

const props = withDefaults(defineProps<Props>(), {
  prefix: '1cs768',
  sceneId: 'xesx3o47'
})

//使用默认插槽生成自己的组件并追加id属性
const slots = useSlots()
const id = randomString(8)
const defaultSlot = slots.default?.()?.[0] || 'span'
const MySlot = h(defaultSlot, { id })

// 获取验证码验证结果和业务结果
const captchaVerifyCallback = async (captchaVerifyParam: {
  sceneId: string,
  certifyId: string,
  deviceToken: string,
  data: string,
}): Promise<{
  captchaResult: boolean,
  bizResult: boolean
}> => {
  // // 1.向后端发起业务请求，获取验证码验证结果和业务结果
  // const result = await xxxx('http://您的业务请求地址', {
  //   captchaVerifyParam: captchaVerifyParam, // 验证码参数
  //   yourBizParam... // 业务参数
  // });
  //
  console.log(captchaVerifyParam.sceneId)
  return { captchaResult: true, bizResult: false }
}

// 业务请求验证结果回调函数
const emits = defineEmits<{
  (e: 'result', result: boolean): void
}>()
const onBizResultCallback = (bizResult: boolean) => {
  emits('result', bizResult)
}

const { locale } = useI18n()
const initCaptcha = () => {
  initAliyunCaptcha({
    mode: 'popup',
    button: `#${id}`,
    prefix: props.prefix,
    SceneId: props.sceneId,
    captchaVerifyCallback: captchaVerifyCallback, // 业务请求(带验证码校验)回调函数，无需修改
    onBizResultCallback: onBizResultCallback, // 业务请求结果回调函数，无需修改
    slideStyle: { width: 360, height: 30 }, // 滑块验证码样式，支持自定义宽度和高度，单位为px。其中，width最小值为320 px
    language: locale.value || 'cn' // 验证码语言类型，支持简体中文（cn）、繁体中文（tw）、英文（en）
  })
}

onMounted(() => {
  const script = document.createElement('script')
  script.type = 'text/javascript'
  script.setAttribute('data-sdk', 'captcha-v2')
  script.src = '//o.alicdn.com/captcha-frontend/aliyunCaptcha/AliyunCaptcha.js'
  script.onload = initCaptcha
  // 引入失败
  script.onerror = function() {
    console.log('AliyunCaptcha jssdk 资源加载失败了')
  }
  document.head.appendChild(script)
})

onUnmounted(() => {
  let script = document.querySelector('script[data-sdk=\'captcha-v2\']')
  script && document.head.removeChild(script)
})
</script>

<style lang="scss">
#aliyunCaptcha-sliding-body {
  box-sizing: content-box;
}

.aliyunCaptcha-sliding-slider {
  background-color: var(--el-color-primary) !important;
}
</style>