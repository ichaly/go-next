<template>
  <div class="h-full">
    <flat-button @click="visible = true">
      <IconSearch />
    </flat-button>
    <el-dialog v-model="visible" :show-close="false" @open="onDialogOpen">
      <template #header>
        <el-input v-model="input" clearable size="large" ref="inputRef" placeholder="搜索">
          <template #prefix>
            <IconSearch />
          </template>
        </el-input>
      </template>
      <div class="search-no-data">暂无搜索结果</div>
      <template #footer>
        <div class="search-footer">
          <span class="search-footer-item">
            <el-icon :size="16">
              <i class="i-mdi:subdirectory-arrow-left" />
            </el-icon>
          </span>
          <span>确认</span>
          <span class="search-footer-item">
            <el-icon :size="16">
              <i class="i-mdi:arrow-up" />
            </el-icon>
          </span>
          <span class="search-footer-item">
            <el-icon :size="16">
              <i class="i-mdi:arrow-down" />
            </el-icon>
          </span>
          <span>切换</span>
          <span class="search-footer-item">
            <el-icon :size="16">
              <i class="i-mdi:keyboard-esc" />
            </el-icon>
          </span>
          <span>关闭</span>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
const input = ref('')
const visible = ref(false)
const inputRef = ref<HTMLElement | null>(null)

const onDialogOpen = () => {
  nextTick(() => {
    unref(inputRef)?.focus()
  })
}
</script>

<style scoped lang="scss">
:deep(.el-dialog__header) {
  margin-right: 0;
  padding-bottom: 0;
}

:deep(.el-dialog__footer) {
  padding: 0;
}

.search {
  &-no-data {
    @apply flex center;
    height: 100px;
    color: rgb(150 159 175);
  }

  &-footer {
    display: flex;
    position: relative;
    flex-shrink: 0;
    align-items: center;
    height: 44px;
    padding: 0 20px;
    border-top: 1px solid var(--el-border-color);
    border-radius: 0 0 16px 16px;
    color: #666;
    font-size: 12px;

    &-item {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 20px;
      height: 18px;
      margin-right: 0.4em;
      padding-bottom: 2px;
      border-radius: 2px;
      background-color: linear-gradient(-225deg, #d5dbe4, #f8f8f8);
      box-shadow:
        inset 0 -2px 0 0 #cdcde6,
        inset 0 0 1px 1px #fff,
        0 1px 2px 1px rgb(30 35 90 / 40%);

      &:nth-child(2),
      &:nth-child(3),
      &:nth-child(6) {
        margin-left: 14px;
      }
    }
  }
}
</style>
