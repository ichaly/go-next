<template>
  <div class="root">
    <div class="name">
      <img src="@/assets/images/logo.svg" class="logo" />
      <span class="text">Go Next Admin</span>
    </div>
    <div class="form">
      <div class="corner text-4xl" @click="toggleScan()">
        <i class="i-ri:qr-code-line" v-if="isScan" />
        <i class="i-ri:computer-line" v-else />
      </div>
      <div class="left"></div>
      <div class="right">
        <div class="login">
          <div class="account" v-if="!isScan">
            <el-tabs v-model="activeTab" stretch>
              <el-tab-pane name="first">
                <template #label><span class="label">密码登录</span></template>
                <el-form :model="accountForm" ref="accountRef" :rules="accountRules">
                  <el-form-item prop="username">
                    <el-input v-model="accountForm.username" clearable placeholder="手机/邮箱/账号">
                      <template #prefix>
                        <i class="i-ri:user-5-line" />
                      </template>
                    </el-input>
                  </el-form-item>
                  <el-form-item prop="password">
                    <el-input
                      v-model="accountForm.password"
                      type="password"
                      clearable
                      show-password
                      placeholder="请输入密码">
                      <template #prefix>
                        <i class="i-ri:lock-password-line" />
                      </template>
                    </el-input>
                  </el-form-item>
                  <el-row class="w-full items-center m-b-4.5" justify="space-between">
                    <el-checkbox v-model="isRemember" label="记住密码" size="large" />
                    <el-link type="primary">找回密码</el-link>
                  </el-row>
                  <Captcha @result="onSubmitResult">
                    <el-button type="primary" class="w-full">立即登录</el-button>
                  </Captcha>
                </el-form>
              </el-tab-pane>
              <el-tab-pane name="second">
                <template #label><span class="label">免密登录</span></template>
                <el-form :model="mobileForm" ref="mobileRef" :rules="mobileRules">
                  <el-form-item prop="mobile">
                    <el-input v-model="mobileForm.mobile" clearable placeholder="手机/邮箱">
                      <template #prefix>
                        <i class="i-ri:tablet-line" />
                      </template>
                    </el-input>
                  </el-form-item>
                  <el-form-item prop="captcha">
                    <div class="flex w-full flex-row">
                      <el-input v-model="mobileForm.captcha" clearable placeholder="请输入验证码">
                        <template #prefix>
                          <i class="i-ri:shield-user-line" />
                        </template>
                      </el-input>
                      <div class="w-3"></div>
                      <Captcha @result="onCaptchaResult">
                        <el-button type="primary">获取验证码</el-button>
                      </Captcha>
                    </div>
                  </el-form-item>
                  <el-row class="desc w-full m-b-4.5">
                    <el-checkbox v-model="isRemember" size="large" />
                    <span class="ml-2"></span>
                    我已阅读并同意&nbsp;
                    <el-link type="primary" class="link">用户协议</el-link>
                    &nbsp;和&nbsp;
                    <el-link type="primary" class="link">隐私政策</el-link>
                  </el-row>
                  <el-button type="primary" class="w-full" @click="onSubmit(mobileRef)">登录/注册</el-button>
                </el-form>
              </el-tab-pane>
            </el-tabs>
          </div>
          <div class="scan" v-else>
            <div class="title">手机扫码，安全登陆</div>
            <qrcode-vue :value="content" size="132" class="code" />
            <div class="desc">
              使用&nbsp;
              <el-link type="primary" class="link">客户端</el-link>
              &nbsp;扫描二维码安全快速登录
            </div>
          </div>
        </div>
        <div class="third">
          <el-divider><span class="divider">其他方式登录</span></el-divider>
          <div class="icons">
            <div class="icon" v-for="(v, i) in platforms" :key="i" :style="`background-color:${v.color} ;`">
              <i :class="v.icon" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import QrcodeVue from 'qrcode.vue'
import type { FormInstance, FormRules } from 'element-plus'

interface Platform {
  icon: string
  color: string
}

const platforms = ref<Array<Platform>>([
  { icon: 'i-ri:wechat-line', color: '#4caf50' },
  { icon: 'i-ri:weibo-line', color: '#ff9800' },
  { icon: 'i-ri:qq-line', color: '#2196f3' },
  { icon: 'i-ri:github-line', color: '#000000' }
])
const [isScan, toggleScan] = useToggle(false)
const [isRemember] = useToggle(true)
const content = ref('code123')
const activeTab = ref('first')

const accountRef = ref<FormInstance>()
const accountForm = reactive({
  username: '',
  password: ''
})
const accountRules = reactive<FormRules>({
  username: [{ required: true, message: '请输入账号', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
})

const mobileRef = ref<FormInstance>()
const mobileForm = reactive({
  mobile: '',
  captcha: ''
})
const mobileRules = reactive<FormRules>({
  mobile: [{ required: true, message: '请输入接收账号', trigger: 'blur' }],
  captcha: [{ required: true, message: '请输入验证码', trigger: 'blur' }]
})

const onSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      console.log('submit!')
    } else {
      console.log('error submit!', fields)
    }
  })
}

const onCaptchaResult = (result: boolean) => {
  console.log('onCaptchaResult', result)
}

const onSubmitResult = (result: boolean) => {
  console.log('onSubmitResult', result)
  if (result) {
    onSubmit(accountRef.value)
  }
}
</script>

<style scoped lang="scss">
:deep(.el-tabs__nav-wrap) {
  &::after {
    height: 1px;
  }
}

.root {
  @apply flex flex-col h-screen w-full items-center justify-center bg-cover bg-center bg-no-repeat pb-50;
  background-color: #f0f2f5;
  background-image: url('@/assets/images/login/bg.jpg');

  .name {
    @apply flex flex-col center md:flex-row;
    background: linear-gradient(to right, var(--el-color-primary), #F74952) no-repeat right bottom;
    background-size: 0 2px;
    transition: background-size 0.5s ease-in-out;

    &:hover {
      background-position-x: left;
      background-size: 100% 2px;
    }

    .logo {
      @apply w-24 md:w-14 md:mr-4;
    }

    .text {
      @apply text-3xl font-bold cursor-pointer;
      font-family: Courier, monospace;
    }
  }

  .form {
    @apply relative flex overflow-hidden rounded-md bg-white shadow-md w-[350px] md:w-[600px] mt-10 md:mt-20;
    height: 420px;

    .corner {
      @apply absolute right-0 z-50 flex cursor-pointer items-center justify-center rounded-bl-full text-white;
      width: 72px;
      height: 72px;
      background-color: #4d4d4d;
    }

    .left {
      @apply h-full bg-center bg-no-repeat bg-[length:80%] w-0 md:w-[250px] op-80;
      background-color: var(--el-color-primary);
      background-image: url('@/assets/images/login/online_posts.svg');
    }

    .right {
      @apply flex flex-1 flex-col p-6 pt-10;
      .login {
        @apply flex-1;
        .account {
          padding-top: 10px;

          .label {
            font-size: 18px;
            font-weight: normal;
          }

          .eye {
            @apply flex items-center justify-center;
          }
        }

        .scan {
          @apply flex h-full flex-1 flex-col items-center justify-center;
          .title {
            @apply text-lg;
          }

          .code {
            @apply my-8 rounded border border-gray-200 p-2;
          }
        }

        .desc {
          @apply flex items-center text-sm;

          .link {
            @apply text-sm;
          }
        }
      }

      .third {
        @apply flex flex-col;
        .divider {
          @apply text-xs font-light;
          color: #8c92a4;
        }

        .icons {
          @apply flex items-center justify-between text-xl;

          .icon {
            @apply flex h-8 w-8 cursor-pointer items-center justify-center rounded-full text-white;
          }
        }
      }
    }
  }
}
</style>
