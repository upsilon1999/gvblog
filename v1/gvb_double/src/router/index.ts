import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'WebEnter',
      component: () => import('../views/web/WebEnter.vue'),
      children: [
        {
          //path留空代表默认填充的子路由
          path: '',
          name: 'WebIndex',
          component: () => import('../views/web/WebIndex.vue'),
        },
      ],
    },
    {
      path: '/admin',
      name: 'AdminIndex',
      meta: {
        title: '首页',
        isLogin: true,
      },
      component: () => import('../views/admin/AdminIndex.vue'),
      children: [
        {
          //path留空代表默认填充的子路由
          path: '',
          name: 'AdminHome',
          component: () => import('../views/admin/Home/HomeIndex.vue'),
        },
        {
          path: 'userCenter',
          name: 'UserCenter',
          //由于没有配置component，所以直接访问 /admin/userCenter会为404
          //但是访问/admin/userCenter/userInfo是正常的
          children: [
            {
              path: 'userInfo',
              //个人信息
              name: 'UserInfo',
              meta: {
                title: '我的信息',
              },
              component: () =>
                import('../views/admin/UserCenter/UserInfo/UserInfo.vue'),
            },
            {
              path: 'userArticle',
              name: 'UserArticle',
              meta: {
                title: '我的发布',
              },
              component: () =>
                import('../views/admin/UserCenter/UserArticle/UserArticle.vue'),
            },
            {
              path: 'collects',
              name: 'UserCollects',
              meta: {
                title: '我的收藏',
              },
              component: () =>
                import(
                  '../views/admin/UserCenter/UserCollects/UserCollects.vue'
                ),
            },
            {
              path: 'messages',
              name: 'UserMessage',
              meta: {
                title: '我的消息',
              },
              component: () =>
                import('../views/admin/UserCenter/UserMessage/UserMessage.vue'),
            },
          ],
        },
        {
          path: 'articleMgr',
          //文章管理
          name: 'ArticleMgr',
          children: [
            {
              path: 'articleList',
              name: 'ArticleList',
              meta: {
                title: '文章列表',
              },
              component: () =>
                import('../views/admin/ArticleMgr/ArticleList/ArticleList.vue'),
            },
            {
              path: 'imageList',
              name: 'ImageList',
              meta: {
                title: '图片列表',
              },
              component: () =>
                import('../views/admin/ArticleMgr/ImageList/ImageList.vue'),
            },
            {
              path: 'commentList',
              name: 'CommentList',
              meta: {
                title: '评论列表',
              },
              component: () =>
                import('../views/admin/ArticleMgr/CommentList/CommentList.vue'),
            },
          ],
        },
        {
          path: 'chatGroup',
          //群聊管理
          name: 'ChatGroup',
          meta: {
            title: '群聊管理',
            isAdmin: true,
            isTourist: true,
          },
          children: [
            {
              path: 'chatList',
              name: 'ChatList',
              meta: {
                title: '聊天记录',
              },
              component: () =>
                import('../views/admin/ChatGroup/ChatList/ChatList.vue'),
            },
          ],
        },
        {
          path: 'systemMgr',
          //系统管理
          name: 'SystemMgr',
          meta: {
            title: '系统管理',
            isAdmin: true,
            isTourist: false,
          },
          children: [
            {
              path: 'menuList',
              name: 'MenuList',
              meta: {
                title: '菜单列表',
              },
              component: () =>
                import('../views/admin/SystemMgr/MenuList/MenuList.vue'),
            },

            {
              path: 'feedbackList',
              name: 'FeedbackList',
              meta: {
                title: '用户反馈',
              },
              component: () =>
                import(
                  '../views/admin/SystemMgr/FeedbackList/FeedbackList.vue'
                ),
            },
            {
              path: 'promotionList',
              name: 'PromotionList',
              meta: {
                title: '广告列表',
              },
              component: () =>
                import(
                  '../views/admin/SystemMgr/PromotionList/PromotionList.vue'
                ),
            },
            {
              path: 'logList',
              name: 'LogList',
              meta: {
                title: '系统日志',
              },
              component: () =>
                import('../views/admin/SystemMgr/LogList/LogList.vue'),
            },
            {
              path: 'configList',
              name: 'ConfigList',
              meta: {
                title: '系统配置',
              },
              redirect: '/admin/systemMgr/ConfigList/site',
              component: () =>
                import('../views/admin/SystemMgr/ConfigList/SystemMgr.vue'),
              children: [
                {
                  path: 'site',
                  name: 'SiteConfig',
                  meta: {
                    title: '网站设置',
                  },
                  component: () =>
                    import(
                      '../views/admin/SystemMgr/ConfigList/SiteConfig/SiteConfig.vue'
                    ),
                },
                {
                  path: 'email',
                  name: 'EmailConfig',
                  meta: {
                    title: '邮箱设置',
                  },
                  component: () =>
                    import(
                      '../views/admin/SystemMgr/ConfigList/EmailConfig/EmailConfig.vue'
                    ),
                },
                {
                  path: 'qiniu',
                  name: 'QiniuConfig',
                  meta: {
                    title: '七牛云设置',
                  },
                  component: () =>
                    import(
                      '../views/admin/SystemMgr/ConfigList/QiniuConfig/QiniuConfig.vue'
                    ),
                },
                {
                  path: 'qq',
                  name: 'QqConfig',
                  meta: {
                    title: 'QQ设置',
                  },
                  component: () =>
                    import(
                      '../views/admin/SystemMgr/ConfigList/QqConfig/QqConfig.vue'
                    ),
                },
                {
                  path: 'jwt',
                  name: 'JwtConfig',
                  meta: {
                    title: 'jwt设置',
                  },
                  component: () =>
                    import(
                      '../views/admin/SystemMgr/ConfigList/JwtConfig/JwtConfig.vue'
                    ),
                },
                {
                  path: 'gaode',
                  name: 'GaodeConfig',
                  meta: {
                    title: '高德设置',
                  },
                  component: () =>
                    import(
                      '../views/admin/SystemMgr/ConfigList/GaoDeConfig/GaoDeConfig.vue'
                    ),
                },
              ],
            },
          ],
        },
        {
          path: 'usersMgr',
          //用户管理
          name: 'UsersMgr',
          meta: {
            title: '用户管理',
            isAdmin: true,
            isTourist: true,
          },
          children: [
            {
              path: 'userList',
              name: 'UserList',
              meta: {
                title: '用户列表',
              },
              component: () =>
                import('../views/admin/UsersMgr/UserList/UserList.vue'),
            },
            {
              path: 'messageList',
              name: 'MessageList',
              meta: {
                title: '消息列表',
              },
              component: () =>
                import('../views/admin/UsersMgr/MessageList/MessageList.vue'),
            },
          ],
        },
      ],
    },
    {
      path: '/:pathMatch(.*)*', // 页面不存在的情况下会跳到404页面
      name: 'NotFound',
      component: () => import('../views/web/NotFound.vue'),
    },
  ],
})

export default router
