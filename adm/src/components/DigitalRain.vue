<template>
  <div class="w-screen h-screen absolute inset-0">
    <canvas ref="canvas" class="absolute w-full h-full" />
  </div>
</template>

<script setup lang="ts">
const fontSize = 16
const canvas = ref()
const columnCount = ref(0)
const nextChar = ref<number[]>([])

const ctx = computed(() => canvas.value.getContext('2d'))
const { width: windowWidth, height: windowHeight } = useWindowSize()

const dpr = ref(window.devicePixelRatio || 1)
const size = ref(fontSize * unref(dpr))
const width = ref(unref(windowWidth) * unref(dpr))
const height = ref(unref(windowHeight) * unref(dpr))

function randomColor() {
  const colors = [
    '#33b5e5',
    '#0099cc',
    '#aa66cc',
    '#9933cc',
    '#99cc00',
    '#669900',
    '#ffbb33',
    '#ff8800',
    '#ff4444',
    '#cc0000'
  ]
  return colors[Math.floor(Math.random() * colors.length)]
}

function randomText() {
  const text = 'console.log("hello world")'
  return text[Math.floor(Math.random() * text.length)]
}

function draw() {
  const c = unref(ctx)
  c.fillStyle = 'rgba(255,255,255,0.1)'
  c.fillRect(0, 0, unref(width), unref(height))
  for (let i = 0; i < unref(columnCount); i++) {
    const x = i * unref(size)
    const y = (nextChar.value[i] + 1) * unref(size)
    c.font = `${unref(size)}px "Roboto Mono"`
    c.fillStyle = randomColor()
    c.fillText(randomText(), x, y)
    if (y > unref(height) && Math.random() > 0.9) {
      nextChar.value[i] = 0
    } else {
      nextChar.value[i]++
    }
  }
}

function init() {
  if (!canvas.value) return
  dpr.value = window.devicePixelRatio || 1

  size.value = fontSize * unref(dpr)
  width.value = unref(windowWidth) * unref(dpr)
  height.value = unref(windowHeight) * unref(dpr)

  canvas.value.width = unref(width)
  canvas.value.height = unref(height)

  columnCount.value = Math.ceil(unref(width) / unref(size))
  nextChar.value = new Array(unref(columnCount)).fill(0)
}

watch([windowWidth, windowHeight, canvas], init, { immediate: true })
onMounted(() => {
  setInterval(draw, 33)
})
</script>
