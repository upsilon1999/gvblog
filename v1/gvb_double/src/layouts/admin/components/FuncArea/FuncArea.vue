<template>
  <div class="gvb-func-area">
    <!-- 前往首页 -->
    <IconHome class="action-icon"></IconHome>
    <!-- 主题切换按钮 -->
    <GvbTheme />
    <!-- 
              个人信息区域
              主题是图片名字
              点开是下拉菜单
             -->
    <div class="gvb-person-info">
      <a-dropdown>
        <!-- 
                显示区域 头像加名字,左右结构
                可以直接采用arco的头像组件
                 -->
        <div class="gvb-avatar">
          <img src="/images/logo.jpg" alt="" />
          <span>upsilon</span>
        </div>
        <!-- 下拉菜单 -->
        <template #content>
          <a-doption>Option 1</a-doption>
          <a-doption>Option 3</a-doption>
          <a-doption>Option 4</a-doption>
          <a-doption>Option 5</a-doption>
          <a-doption>Option 6</a-doption>
        </template>
      </a-dropdown>
    </div>
  </div>
</template>

<script setup lang="ts">
import { IconHome } from '@arco-design/web-vue/es/icon'
import GvbTheme from './GvbTheme/GvbTheme.vue'
</script>

<style lang="scss">
.gvb-func-area {
  display: flex;
  align-items: center;
  .action-icon {
    margin-right: 10px;
    font-size: 15px;
    cursor: pointer;
  }
  .gvb-person-info {
    cursor: pointer;
    .gvb-avatar {
      display: flex;
      align-items: center;
      img {
        width: 40px;
        height: 40px;
        border-radius: 50%;
      }
      span {
        display: inline-block;
        margin: 0 5px;
      }
    }
  }
}
</style>
