<template>
  <el-tooltip :disabled="isShow" :content="text" placement="top" popper-class="tooltip">
    <div class="ellipsis" @mouseover="onMouseOver">
      <span ref="refName">{{ text }}</span>
    </div>
  </el-tooltip>
</template>

<script setup lang="ts">
defineProps<{ text: string }>()

const refName = ref()
const isShow = ref(false)
const onMouseOver = () => {
  const parentWidth = refName.value.parentNode.offsetWidth
  const contentWidth = refName.value.offsetWidth
  // 判断是否开启 tooltip 功能，如果溢出显示省略号，则子元素的宽度势必短于父元素的宽度
  if (contentWidth > parentWidth) {
    isShow.value = false
  } else {
    isShow.value = true
  }
}
</script>

<style scoped lang="scss">
.ellipsis {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  //@apply m-b-1 text-ellipsis overflow-hidden;
  //display: -webkit-box;
  //-webkit-line-clamp: 2;
  //-webkit-box-orient: vertical;
}
</style>
