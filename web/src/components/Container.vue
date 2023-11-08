<template>
  <component :is="as" :class="clazz" v-bind="$attrs">
    <slot/>
  </component>
</template>

<script lang="ts" setup>
const {padded, fluid} = withDefaults(defineProps<{
  as?: string
  fluid?: boolean
  padded?: boolean
}>(), {
  as: 'div',
  fluid: true,
  padded: false,
})
const clazz = computed(() => {
  return ['base', {padding: padded, constrained: fluid}]
})
</script>
<style scoped lang="scss">
.base {
  @apply m-a w-full;
}

.padding {
  @apply px-4 sm:px-6 lg:px-8;
}

.constrained {
  @apply max-w-7xl;
}
</style>