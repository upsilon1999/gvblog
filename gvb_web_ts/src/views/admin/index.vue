<template>
  <div class="gvb-admin">
    <!-- 
        aside 左侧侧边栏
        固定240px，当屏幕宽度过窄时需要自动收起
        分为头部logo区域和下部菜单区域
    -->
    <aside>
      <!-- 
        logo区域 分为左右结构，左侧是图片，右侧是文字  
      -->
      <div class="gvb-logo">
        <img src="/images/logo.jpg" alt="" />
        <div class="logo-head">
          <span>小邓知识库</span>
          <span>upsilon's store</span>
        </div>
      </div>
      <!-- 菜单区域 -->
      <div class="gvb-menu">
        <a-menu
          :style="{ width: '200px', height: '100%' }"
          :default-open-keys="['0']"
          :default-selected-keys="['0_2']"
        >
          <a-sub-menu key="0">
            <template #icon><icon-apps></icon-apps></template>
            <template #title>Navigation 1</template>
            <a-menu-item key="0_0">Menu 1</a-menu-item>
            <a-menu-item key="0_1">Menu 2</a-menu-item>
            <a-menu-item key="0_2">Menu 3</a-menu-item>
            <a-menu-item key="0_3">Menu 4</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="1">
            <template #icon><icon-bug></icon-bug></template>
            <template #title>Navigation 2</template>
            <a-menu-item key="1_0">Menu 1</a-menu-item>
            <a-menu-item key="1_1">Menu 2</a-menu-item>
            <a-menu-item key="1_2">Menu 3</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="2">
            <template #icon><icon-bulb></icon-bulb></template>
            <template #title>Navigation 3</template>
            <a-menu-item key="2_0">Menu 1</a-menu-item>
            <a-menu-item key="2_1">Menu 2</a-menu-item>
            <a-sub-menu key="2_2" title="Navigation 4">
              <a-menu-item key="2_2_0">Menu 1</a-menu-item>
              <a-menu-item key="2_2_1">Menu 2</a-menu-item>
            </a-sub-menu>
          </a-sub-menu>
        </a-menu>
      </div>
    </aside>
    <main>
      <RouterView />
    </main>
  </div>
</template>

<script setup lang="ts">
import { RouterView } from "vue-router";
import {
  IconMenuFold,
  IconMenuUnfold,
  IconApps,
  IconBug,
  IconBulb,
} from "@arco-design/web-vue/es/icon";
</script>

<style scoped lang="scss">
.gvb-admin {
  display: flex;

  aside {
    width: 240px;

    //左侧灰色分隔线
    border-right: 1px solid var(--bg);
    height: 100dvh;

    .gvb-logo {
      height: 90px;
      display: flex;
      padding: 20px;
      align-items: center;
      img {
        width: 60px;
        height: 60px;
        border-radius: 50%;
      }
      .logo-head {
        margin-left: 20px;
        > span:nth-child(1) {
          font-size: 22px;
        }
        > span:nth-child(2) {
          font-size: 12px;
        }
      }
    }
  }

  main {
    //注意calc的减号两端要留空格，否则会报错
    width: calc(100% - 240px);
  }
}
</style>
