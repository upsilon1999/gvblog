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
      <!-- 
        头部区域 
        分为左右区域，左边为面包屑，右边是功能按钮和个人信息
      -->
      <div class="gvb-head">
        <!-- 左侧头部面包屑 -->
        <div class="gvb-bread-crumbs">
          <a-breadcrumb>
            <a-breadcrumb-item>Home</a-breadcrumb-item>
            <a-breadcrumb-item>Channel</a-breadcrumb-item>
            <a-breadcrumb-item>News</a-breadcrumb-item>
          </a-breadcrumb>
        </div>
        <!-- 右侧功能区域  -->
        <div class="gvb-func-area">
          <!-- 前往首页 -->
          <IconHome class="action-icon"></IconHome>
          <!-- 主题切换按钮 -->
          <div class="gvb_theme">
            <IconSun class="action-icon"></IconSun>
          </div>
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
      </div>

      <div class="gvb-tabs">
        <span class="gvb-tab active">首页</span>
        <span class="gvb-tab">用户列表</span>
        <span class="gvb-tab">文章列表</span>
      </div>
      <!-- 内容区域 -->
      <div class="gvb-container">
        <RouterView />
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { RouterView } from "vue-router";
import {
  IconHome,
  IconSun,
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
  }

  main {
    //注意calc的减号两端要留空格，否则会报错
    width: calc(100% - 240px);

    .gvb-head {
      width: 100%;
      height: 60px;
      //底部分隔线
      border-bottom: 1px solid var(--bg);

      display: flex;
      justify-content: space-between; //两边对齐
      padding: 0 20px;
      align-items: center; //垂直居中

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
      min-height: calc(100dvh - 90px);
      padding: 20px;
    }
  }
}
</style>
