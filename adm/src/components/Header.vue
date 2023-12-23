<template>
  <div class="flex items-center h-full">
    <flat-button @click="toggleCollapse()">
      <el-icon :size="18" :class="['transition-transform', { '-scale-x-100': isCollapse }]">
        <i class="i-icon-park-outline:indent-left" />
      </el-icon>
    </flat-button>
    <div class="flex-1 pl-2">
      <Breadcrumb />
    </div>
    <Search />
    <flat-button @click="toggleFullscreen()">
      <el-icon :size="18">
        <i class="i-icon-park-outline:off-screen-one" v-if="isFullscreen" />
        <i class="i-icon-park-outline:full-screen-one" v-else />
      </el-icon>
    </flat-button>
    <flat-button>
      <el-badge :value="200" :max="99" class="h-4.5">
        <el-icon :size="18">
          <i class="i-icon-park-outline:remind" />
        </el-icon>
      </el-badge>
    </flat-button>
    <el-dropdown class="h-full">
      <flat-button>
        <el-icon :size="18">
          <i class="i-icon-park-outline:translate" />
        </el-icon>
      </flat-button>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>
            <span>简体中文</span>
          </el-dropdown-item>
          <el-dropdown-item>
            <span>English</span>
          </el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
    <el-dropdown class="h-full">
      <flat-button>
        <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png">
          user
        </el-avatar>
      </flat-button>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>
            <i class="i-icon-park-outline:user" />
            <span>个人信息</span>
          </el-dropdown-item>
          <el-dropdown-item>
            <i class="i-icon-park-outline:edit" />
            <span>修改密码</span>
          </el-dropdown-item>
          <el-dropdown-item divided>
            <i class="i-icon-park-outline:help" />
            <span>帮助文档</span>
          </el-dropdown-item>
          <el-dropdown-item>
            <i class="i-icon-park-outline:tips" />
            <span>功能更新</span>
          </el-dropdown-item>
          <el-dropdown-item divided>
            <i class="i-icon-park-outline:lock" />
            <span>锁定屏幕</span>
          </el-dropdown-item>
          <el-dropdown-item @click="$router.push('/login')">
            <i class="i-icon-park-outline:power" />
            <span>退出登录</span>
          </el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </div>
</template>

<script setup lang="ts">
const rootStore = useRootStore()
const { toggleCollapse, toggleFullscreen } = rootStore
const { isCollapse, isFullscreen } = toRefs(rootStore)
</script>

<style scoped lang="scss"></style>
