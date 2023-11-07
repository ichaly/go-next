<template>
  <component :is="as" :class="containerClass" v-bind="$attrs">
    <slot/>
  </component>
</template>

<script lang="ts" setup>
import type {ClassNameValue} from "tailwind-merge";
import {twJoin, twMerge} from "tailwind-merge";

const props = withDefaults(defineProps<{
  as?: string
  class?: ClassNameValue
}>(), {
  as: 'div',
  class: undefined,
})
const appConfig = useAppConfig()
// fluid?: boolean
// padded?: boolean
const containerClass = computed(() => {
  return twMerge(twJoin(
  ), props.class)
})
</script>