import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import("../views/web/web.vue"),
      children: [
        {
          //path留空代表默认填充的子路由
          path: "",
          name: "index",
          component: () => import("../views/web/index.vue"),
        },
      ],
    },
    {
      path: "/admin",
      name: "admin",
      meta: {
        title: "首页",
        isLogin: true,
      },
      component: () => import("../views/admin/index.vue"),
      children: [
        {
          //path留空代表默认填充的子路由
          path: "",
          name: "home",
          component: () => import("../views/admin/Home/index.vue"),
        },
        {
          path: "userCenter",
          name: "userCenter",
          //由于没有配置component，所以直接访问 /admin/userCenter会为404
          //但是访问/admin/userCenter/userInfo是正常的
          children: [
            {
              path: "userInfo",
              //个人信息
              name: "userInfo",
              meta: {
                title: "我的信息",
              },
              component: () =>
                import("../views/admin/UserCenter/UserInfo/index.vue"),
            },
            {
              path: "article",
              name: "userArticle",
              meta: {
                title: "我的发布",
              },
              component: () =>
                import("../views/admin/UserCenter/UserArticle/index.vue"),
            },
            {
              path: "collects",
              name: "collects",
              meta: {
                title: "我的收藏",
              },
              component: () =>
                import("../views/admin/UserCenter/UserArticle/index.vue"),
            },
            {
              path: "messages",
              name: "messages",
              meta: {
                title: "我的消息",
              },
              component: () =>
                import("../views/admin/UserCenter/UserMessage/index.vue"),
            },
          ],
        },
        {
          path: "article",
          //文章管理
          name: "article",
          children: [
            {
              path: "articleList",
              name: "articleList",
              meta: {
                title: "文章列表",
              },
              component: () =>
                import("../views/admin/Article/ArticleList/index.vue"),
            },
            {
              path: "imageList",
              name: "imageList",
              meta: {
                title: "图片列表",
              },
              component: () =>
                import("../views/admin/Article/ImageList/index.vue"),
            },
            {
              path: "CommentList",
              name: "CommentList",
              meta: {
                title: "评论列表",
              },
              component: () =>
                import("../views/admin/Article/CommentList/index.vue"),
            },
          ],
        },
        {
          path: "chatGroup",
          //群聊管理
          name: "chatGroup",
          meta: {
            title: "群聊管理",
            isAdmin: true,
            isTourist: true,
          },
          children: [
            {
              path: "chatList",
              name: "chatList",
              meta: {
                title: "聊天记录",
              },
              component: () =>
                import("../views/admin/ChatGroup/ChatList/index.vue"),
            },
          ],
        },
        {
          path: "system",
          //系统管理
          name: "system",
          meta: {
            title: "系统管理",
            isAdmin: true,
            isTourist: false,
          },
          children: [
            {
              path: "menuList",
              name: "menuList",
              meta: {
                title: "菜单列表",
              },
              component: () =>
                import("../views/admin/System/MenuList/index.vue"),
            },

            {
              path: "feedbackList",
              name: "feedbackList",
              meta: {
                title: "用户反馈",
              },
              component: () =>
                import("../views/admin/System/FeedbackList/index.vue"),
            },
            {
              path: "promotionList",
              name: "promotionList",
              meta: {
                title: "广告列表",
              },
              component: () =>
                import("../views/admin/System/PromotionList/index.vue"),
            },
            {
              path: "logList",
              name: "logList",
              meta: {
                title: "系统日志",
              },
              component: () =>
                import("../views/admin/System/LogList/index.vue"),
            },
            {
              path: "config",
              name: "systemConfig",
              meta: {
                title: "系统配置",
              },
              redirect: "/admin/system/system/site",
              component: () => import("../views/admin/System/Config/index.vue"),
              children: [
                {
                  path: "site",
                  name: "siteConfig",
                  meta: {
                    title: "网站设置",
                  },
                  component: () =>
                    import("../views/admin/System/Config/SiteConfig/index.vue"),
                },
                {
                  path: "email",
                  name: "emailConfig",
                  meta: {
                    title: "邮箱设置",
                  },
                  component: () =>
                    import(
                      "../views/admin/System/Config/EmailConfig/index.vue"
                    ),
                },
                {
                  path: "qiniu",
                  name: "qiniuConfig",
                  meta: {
                    title: "七牛云设置",
                  },
                  component: () =>
                    import(
                      "../views/admin/System/Config/QiniuConfig/index.vue"
                    ),
                },
                {
                  path: "qq",
                  name: "qqConfig",
                  meta: {
                    title: "QQ设置",
                  },
                  component: () =>
                    import("../views/admin/System/Config/QqConfig/index.vue"),
                },
                {
                  path: "jwt",
                  name: "jwtConfig",
                  meta: {
                    title: "jwt设置",
                  },
                  component: () =>
                    import("../views/admin/System/Config/JwtConfig/index.vue"),
                },
                {
                  path: "gaode",
                  name: "gaode_config",
                  meta: {
                    title: "高德设置",
                  },
                  component: () =>
                    import(
                      "../views/admin/System/Config/GaoDeConfig/index.vue"
                    ),
                },
              ],
            },
          ],
        },
        {
          path: "users",
          //用户管理
          name: "users",
          meta: {
            title: "用户管理",
            isAdmin: true,
            isTourist: true,
          },
          children: [
            {
              path: "userList",
              name: "userList",
              meta: {
                title: "用户列表",
              },
              component: () =>
                import("../views/admin/Users/UserList/index.vue"),
            },
            {
              path: "messageList",
              name: "messageList",
              meta: {
                title: "消息列表",
              },
              component: () =>
                import("../views/admin/Users/MessageList/index.vue"),
            },
          ],
        },
      ],
    },
    {
      path: "/:pathMatch(.*)*", // 页面不存在的情况下会跳到404页面
      name: "notFound",
      component: () => import("../views/web/NotFound.vue"),
    },
  ],
});

export default router;
