<template>
  <div class="w-full h-full flex flex-col center">
    <div class="loader"></div>
    <div class="load_title">正在努力加载中</div>
  </div>
</template>
<style scoped lang="scss">
@mixin border() {
  border-radius: 50%;
  border: 3px solid transparent;
  border-top-color: #409eff;
  animation: spin 2s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.loader {
  width: 50px;
  height: 50px;
  position: relative;
  @include border;

  &:before {
    content: '';
    position: absolute;
    inset: 5px;
    @include border;
  }

  &:after {
    content: '';
    position: absolute;
    inset: 15px;
    @include border;
    animation-duration: 3s;
  }
}

.load_title {
  margin-top: 30px;
  color: #409eff;
  font-size: 16px;
  line-height: 30px;
}
</style>
