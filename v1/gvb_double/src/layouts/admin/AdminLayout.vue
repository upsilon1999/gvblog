<template>
  <div class="gvb-admin">
    <!-- 
            aside 左侧侧边栏
            固定240px，当屏幕宽度过窄时需要自动收起
            分为头部logo区域和下部菜单区域
        -->
    <aside :class="{ collapsed: isCollapsed }">
      <!-- logo区域 分为左右结构，左侧是图片，右侧是文字  -->
      <Logo />
      <!-- 菜单区域 -->
      <Menu />
    </aside>
    <main>
      <!-- 
            头部区域 
            分为左右区域，左边为面包屑，右边是功能按钮和个人信息
          -->
      <div class="gvb-head">
        <!-- 左侧头部面包屑 -->
        <BreadCrumbs />
        <!-- 右侧功能区域  -->
        <FuncArea />
      </div>

      <div class="gvb-tabs">
        <!-- <span class="gvb-tab active">首页</span>
          <span class="gvb-tab">用户列表</span>
          <span class="gvb-tab">文章列表</span> -->
        <Tabs />
      </div>
      <!-- 内容区域 -->
      <div class="gvb-container">
        <RouterView />
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import Logo from './components/AdminLogo.vue'
import Menu from './components/Menu/MenuIndex.vue'
// import Tabs from './components/Tabs.vue'
import BreadCrumbs from './components/BreadCrumbs.vue'
import FuncArea from './components/FuncArea/FuncArea.vue'
import useSettingStore from '@/stores/modules/settings'
import { computed } from 'vue'
const settingStore = useSettingStore()
const isCollapsed = computed(() => settingStore.getCollapsed())
</script>

<style lang="scss">
.gvb-admin {
  display: flex;

  aside {
    width: 240px;

    //左侧灰色分隔线
    border-right: 1px solid var(--bg);
    height: 100vh;

    background-color: var(--color-bg-1);
    transition: all 0.3s;
    position: relative;
    .gvb-menu {
      height: calc(100vh - 90px);
      overflow-y: auto;
      overflow-x: hidden;
    }
  }

  aside.collapsed {
    width: 48px;

    & ~ main {
      width: calc(100vw - 48px);
    }
  }

  aside:hover {
    .gvb-menu {
      .arco-menu-collapse-button {
        opacity: 1;
      }
    }
  }

  main {
    //注意calc的减号两端要留空格，否则会报错
    width: calc(100vw - 240px);

    .gvb-head {
      width: 100%;
      height: 60px;
      //底部分隔线
      border-bottom: 1px solid var(--bg);

      display: flex;
      justify-content: space-between; //两边对齐
      padding: 0 20px;
      align-items: center; //垂直居中
    }

    .gvb-tabs {
      height: 30px;
      width: 100%;
      //底部分隔线
      border-bottom: 1px solid var(--bg);
      padding: 0 20px;

      display: flex;
      align-items: center;

      .gvb-tab {
        border-radius: 5px;
        border: 1px solid var(--bg);
        padding: 2px 8px;
        margin-right: 10px;
        cursor: pointer;

        &.active {
          background-color: var(--active);
          color: white;
          border: none;
        }
      }
    }

    .gvb-container {
      background-color: var(--bg);
      min-height: calc(100vh - 90px);
      padding: 20px;
    }
  }
}
</style>
