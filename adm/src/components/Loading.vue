<template>
  <div class='w-full h-full flex flex-col center'>
    <div class='loader'></div>
    <div class='load_title'>正在努力加载中</div>
  </div>
</template>
<style scoped>
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
  border-radius: 50%;
  position: relative;
  border: 3px solid transparent;
  border-top-color: #409eff;
  animation: spin 2s linear infinite;

  &:before {
    content: '';
    position: absolute;
    top: 5px;
    left: 5px;
    right: 5px;
    bottom: 5px;
    border-radius: 50%;
    border: 3px solid transparent;
    border-top-color: #409eff;
    animation: spin 2s linear infinite;
  }

  &:after {
    content: '';
    position: absolute;
    top: 15px;
    left: 15px;
    right: 15px;
    bottom: 15px;
    border-radius: 50%;
    border: 3px solid transparent;
    border-top-color: #409eff;
    animation: spin 2s linear infinite;
  }
}

.load_title {
  margin-top: 30px;
  color: #409eff;
  font-size: 16px;
  line-height: 30px;
}
</style>
