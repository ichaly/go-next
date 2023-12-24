<template>
  <div class="h-full">
    <flat-button @click="visible = true">
      <IconSearch />
    </flat-button>
    <el-dialog v-model="visible" :show-close="false" @open="onDialogOpen">
      <template #header>
        <el-input
          clearable
          size="large"
          ref="input"
          v-model="keyword"
          @input="onSearch"
          placeholder="搜索"
        >
          <template #prefix>
            <IconSearch />
          </template>
        </el-input>
      </template>
      <div class="empty" v-if="result.length === 0">暂无搜索结果</div>
      <ul class="list" v-else>
        <li
          v-for="(item, index) in result"
          :key="item.name"
          :data-index="index"
          :class="[
            'list__item',
            {
              'list__item--active': activeIndex === index
            }
          ]"
          @click="onEnter"
          @mouseenter="onMouseEnter"
        >
          <i :class="['text-4', item.icon || 'i-mdi:form-select']" />
          <span class="text"> {{ formatText(item) }} </span>
          <i class="i-mdi:subdirectory-arrow-left text-4 suffix" />
        </li>
      </ul>
      <template #footer>
        <div class="footer">
          <span class="footer__item">
            <el-icon :size="16">
              <i class="i-mdi:subdirectory-arrow-left" />
            </el-icon>
          </span>
          <span>确认</span>
          <span class="footer__item">
            <el-icon :size="16">
              <i class="i-mdi:arrow-up" />
            </el-icon>
          </span>
          <span class="footer__item">
            <el-icon :size="16">
              <i class="i-mdi:arrow-down" />
            </el-icon>
          </span>
          <span>切换</span>
          <span class="footer__item">
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
import type { RouteMeta } from 'vue-router'
import Fuse from 'fuse.js'

const onDialogOpen = () => {
  nextTick(() => {
    unref(input)?.focus()
  })
}

const keyword = ref('')
const visible = ref(false)
const activeIndex = ref(-1)
const result = ref<RouteMeta[]>([])
const input = ref<HTMLElement | null>(null)
const router = useRouter()
const routers = router.getRoutes().map((item) => item.meta)
const fuse = new Fuse(routers, {
  keys: ['title']
})
const formatText = (item: RouteMeta) => {
  if (!item.items) return ''
  let res = ''
  const length = item.items.length
  for (let i = 0; i < length; i++) {
    const c = item.items[i]
    if (i !== 0) {
      res += ' > '
    }
    res += c.title
  }
  return res
}
const handleSearch = (val: string | number) => {
  val = val.toString().trim()
  result.value = fuse
    .search(val)
    .map((item) => item.item)
    .splice(0, 5)
  activeIndex.value = 0
}
const onSearch = useDebounceFn(handleSearch, 200)
const onUp = () => {
  if (!result.value.length) return
  activeIndex.value--
  if (activeIndex.value < 0) {
    activeIndex.value = result.value.length - 1
  }
}
const onDown = () => {
  if (!result.value.length) return
  activeIndex.value++
  if (activeIndex.value > result.value.length - 1) {
    activeIndex.value = 0
  }
}
const onEnter = () => {
  if (activeIndex.value !== -1) {
    router.push(result.value[activeIndex.value].name)
    visible.value = false
  }
}
const onClose = () => {
  visible.value = false
}
const onMouseEnter = (e: any) => {
  const index = e.target.dataset.index
  activeIndex.value = Number(index)
}
onKeyStroke('Enter', onEnter)
onKeyStroke('ArrowUp', onUp)
onKeyStroke('ArrowDown', onDown)
onKeyStroke('Escape', onClose)
</script>

<style scoped lang="scss">
:deep(.el-dialog__header) {
  margin-right: 0;
  padding-bottom: 0;
}

:deep(.el-dialog__footer) {
  padding: 0;
}

.empty {
  @apply flex center;
  height: 100px;
  color: rgb(150 159 175);
}

.list {
  max-height: 472px;
  padding: 1px;
  overflow: hidden;

  &__item {
    display: flex;
    align-items: center;
    height: 48px;
    border-radius: 4px;
    padding: 0 14px;
    margin: 1px 1px 14px 1px;
    box-shadow: 0 1px 3px 0 #d4d9e1;
    color: var(--el-text-color);
    font-size: 14px;
    cursor: pointer;

    &:last-child {
      margin: 0 0 3px 0;
    }

    > .text {
      margin-left: 10px;
      flex: 1;
    }

    > .suffix {
      opacity: 0;
    }

    &--active {
      background-color: var(--el-color-primary);
      color: #ffffff;

      > .suffix {
        opacity: 1;
      }
    }
  }
}

.footer {
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

  &__item {
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
</style>