<template>
  <component :is="as" :class="containerClass" v-bind="attrs">
    <slot/>
  </component>
</template>

<script lang="ts">
import {computed, toRef, defineComponent} from 'vue'
import type {PropType} from 'vue'
// import {twMerge, twJoin} from 'tailwind-merge'
import type {Strategy} from "~/types";
import {twJoin, twMerge} from "tailwind-merge";
// import {mergeConfig} from '../../utils'
// @ts-expect-error
// import appConfig from '#build/app.config'
// import {container} from '#ui/ui.config'

const config = mergeConfig<typeof container>(appConfig.ui.strategy, appConfig.ui.container, container)

export default defineComponent({
  inheritAttrs: false,
  props: {
    as: {
      type: String,
      default: 'div'
    },
    class: {
      type: [String, Object, Array] as PropType<any>,
      default: undefined
    },
    ui: {
      type: Object as PropType<Partial<typeof config & { strategy?: Strategy }>>,
      default: undefined
    }
  },
  setup(props) {
    const {ui, attrs} = useUI('container', toRef(props, 'ui'), config)

    const containerClass = computed(() => {
      return twMerge(twJoin(
          ui.value.base,
          ui.value.padding,
          ui.value.constrained
      ), props.class)
    })

    return {ui, attrs, containerClass}
  }
})
</script>