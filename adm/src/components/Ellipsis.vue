<template>
  <el-tooltip
    v-bind="$attrs"
    :disabled="enable"
    :content="content"
    :placement="placement"
    :popper-class="popperClassName"
  >
    <div :class="['ellipsis', $attrs.class]" @mouseover="onMouseOver">
      <span ref="refName" :class="className">
        <span>
          <slot>{{ content }}</slot>
        </span>
      </span>
    </div>
  </el-tooltip>
</template>

<script setup lang="ts">
import type { ElTooltipProps } from 'element-plus'

type TooltipProps = Pick<ElTooltipProps, 'content' | 'popperClass' | 'disabled' | 'placement'>

interface Props extends Partial<TooltipProps> {
  line?: number
}

const props = withDefaults(defineProps<Props>(), {
  line: 1,
  placement: 'top'
})

const { line } = toRefs(props)

const enable = computed(() => {
  return props.disabled || isEllipsis.value == false
})

const className = computed(() => {
  return ['ellipsis', props.line > 1 ? 'multi-line-ellipsis' : 'single-line-ellipsis']
})

const popperClassName = computed(() => {
  return [props.popperClass || '', 'ellipsis-tooltip']
})

const refName = ref()
const isEllipsis = ref(false)
const onMouseOver = () => {
  let parentSize, contentSize
  if (line.value > 1) {
    parentSize = refName.value.offsetHeight
    contentSize = refName.value.children[0].offsetHeight
  } else {
    parentSize = refName.value.offsetWidth
    contentSize = refName.value.children[0].offsetWidth
  }
  isEllipsis.value = contentSize > parentSize
}
</script>

<style scoped lang="scss">
.ellipsis {
  display: inline-block;
  overflow: hidden;
  width: 100%;
}

.single-line-ellipsis {
  @apply truncate;
}

.multi-line-ellipsis {
  @apply text-ellipsis overflow-hidden;
  display: -webkit-box;
  -webkit-line-clamp: v-bind(line);
  -webkit-box-orient: vertical;
}
</style>

<style lang="scss">
.ellipsis-tooltip {
  max-width: 200px !important;
}
</style>
