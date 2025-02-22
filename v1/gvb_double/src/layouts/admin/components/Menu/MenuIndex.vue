<template>
  <div class="gvb-menu">
    <a-menu
      @menu-item-click="clickMenu"
      v-model:selected-keys="selectedKeys"
      v-model:open-keys="openKeys"
      show-collapse-button
      @collapse="collapse"
    >
      <template v-for="item in menuList" :key="item.path">
        <!-- 要点:根据有无child来渲染子菜单 -->
        <a-menu-item :key="item.path" v-if="item.child?.length === 0">
          {{ item.title }}
          <template #icon>
            <component :is="item.icon"></component>
          </template>
        </a-menu-item>
        <a-sub-menu v-if="item.child?.length !== 0" :key="item.path">
          <template #icon>
            <component :is="item.icon"></component>
          </template>
          <template #title>{{ item.title }}</template>
          <a-menu-item :key="sub.path" v-for="sub in item.child">
            {{ sub.title }}
            <template #icon>
              <component :is="sub.icon"></component>
            </template>
          </a-menu-item>
        </a-sub-menu>
      </template>
    </a-menu>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import type { Component } from 'vue'
import useSettingStore from '@/stores/modules/settings'
import {
  IconMenu,
  IconUser,
  IconSettings,
  IconMessage,
  IconUserGroup,
  IconBook,
  IconHome,
  IconStorage,
  IconFile,
  IconShareAlt,
  IconImage,
  IconTags,
  IconRobot,
  IconSave,
  IconCustomerService,
  IconEar,
} from '@arco-design/web-vue/es/icon'
import { useRoute, useRouter } from 'vue-router'

/*
	侧边栏的收缩状态，如何存储
	除了menu组件用，index组件也要用，logo最近也要用，所以可以将这个状态写入到pinia
	也可以写入在index组件用v-model，在menu组件变化，父组件知道，兄弟组件也知道
*/
const settingStore = useSettingStore()
const router = useRouter()
const route = useRoute()

//菜单类型
interface MenuType {
  title: string
  icon?: Component
  name?: string // 路由名字
  path?: string //路由路径
  child?: MenuType[]
}

//菜单列表
//name是跳转页面
const menuList: MenuType[] = [
  {
    title: '首页',
    icon: IconHome,
    name: 'AdminIndex',
    path: '/admin',
    child: [],
  },
  {
    title: '个人中心',
    icon: IconUser,
    name: 'UserCenter',
    path: '/admin/userCenter',
    child: [
      {
        title: '我的信息',
        icon: IconTags,
        name: 'UserInfo',
        path: '/admin/userCenter/userInfo',
      },
      {
        title: '我的发布',
        icon: IconRobot,
        name: 'UserArticle',
        path: '/admin/userCenter/userArticle',
      },
      {
        title: '我的收藏',
        icon: IconSave,
        name: 'UserCollects',
        path: '/admin/userCenter/collects',
      },
      {
        title: '我的消息',
        icon: IconMessage,
        name: 'UserMessage',
        path: '/admin/userCenter/messages',
      },
    ],
  },
  {
    title: '文章管理',
    icon: IconBook,
    name: 'ArticleMgr',
    path: '/admin/articleMgr',
    child: [
      {
        title: '文章列表',
        icon: IconBook,
        name: 'ArticleList',
        path: '/admin/articleMgr/articleList',
      },
      {
        title: '图片列表',
        icon: IconImage,
        name: 'ImageList',
        path: '/admin/articleMgr/imageList',
      },
      {
        title: '评论列表',
        icon: IconCustomerService,
        name: 'CommentList',
        path: '/admin/articleMgr/commentList',
      },
    ],
  },
  {
    title: '用户管理',
    icon: IconUserGroup,
    name: 'UsersMgr',
    path: '/admin/usersMgr',
    child: [
      {
        title: '用户列表',
        icon: IconUserGroup,
        name: 'UserList',
        path: '/admin/usersMgr/userList',
      },
      {
        title: '消息列表',
        icon: IconMessage,
        name: 'MessageList',
        path: '/admin/usersMgr/messageList',
      },
    ],
  },
  {
    title: '群聊管理',
    icon: IconMessage,
    name: 'chatGroup',
    path: 'ChatGroup',
    child: [
      {
        title: '聊天记录',
        icon: IconMessage,
        name: 'ChatList',
        path: '/admin/chatGroup/chatList',
      },
    ],
  },
  {
    title: '系统管理',
    icon: IconSettings,
    name: 'SystemMgr',
    path: '/admin/systemMgr',
    child: [
      {
        title: '菜单列表',
        icon: IconMenu,
        name: 'MenuList',
        path: '/admin/systemMgr/menuList',
      },
      {
        title: '用户反馈',
        icon: IconEar,
        name: 'FeedbackList',
        path: '/admin/systemMgr/feedbackList',
      },
      {
        title: '广告列表',
        icon: IconShareAlt,
        name: 'PromotionList',
        path: '/admin/systemMgr/promotionList',
      },
      {
        title: '系统日志',
        icon: IconFile,
        name: 'LogList',
        path: '/admin/systemMgr/logList',
      },
      {
        title: '系统配置',
        icon: IconStorage,
        name: 'ConfigList',
        path: '/admin/systemMgr/configList',
      },
    ],
  },
]

//点击菜单跳转路由
//参数是菜单key值
const clickMenu = (path: string) => {
  /*
    为了配合以后的权限路由，最好改成path跳转
  */
  router.push({
    path: path,
  })
}

const selectedKeys = ref([route.path])
const openKeys = ref([route.matched[1].path])
watch(
  () => route.path,
  () => {
    selectedKeys.value = [route.path]
    openKeys.value = [route.matched[1].path]
  },
)

function collapse(collapsed: boolean) {
  settingStore.setCollapsed(collapsed)
}
</script>

<style lang="scss">
.gvb-menu {
  .arco-menu {
    position: inherit;
  }

  .arco-menu-collapse-button {
    position: absolute;
    right: 0;
    top: 50%;
    transform: translate(50%, -50%);
    opacity: 0;
    transition: all 0.3s;
  }
}
</style>
