<template>
  <div class="gvb-logo">
    <img src="/images/logo.jpg" alt="" />
    <div class="logo-head">
      <span>小邓知识库</span>
      <span>upsilon's store</span>
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped lang="scss">
.gvb-logo {
  height: 90px;
  display: flex;
  padding: 20px;
  align-items: center;

  //底部分隔线
  border-bottom: 1px solid var(--bg);

  img {
    width: 60px;
    height: 60px;
    border-radius: 50%;
  }
  .logo-head {
    margin-left: 20px;
    > span:nth-child(1) {
      font-size: 22px;
      font-weight: 600;
      margin-bottom: 5px;
    }
    > span:nth-child(2) {
      font-size: 12px;
    }
  }
}
</style>
