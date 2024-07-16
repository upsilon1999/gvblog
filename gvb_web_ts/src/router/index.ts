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
              component: () => import("../views/admin/UserCenter/UserInfo.vue"),
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
              component: () => import("../views/admin/Article/ArticleList.vue"),
            },
          ],
        },
        {
          path: "chatGroup",
          //群聊管理
          name: "chatGroup",
          children: [
            {
              path: "chatList",
              name: "chatList",
              component: () => import("../views/admin/ChatGroup/ChatList.vue"),
            },
          ],
        },
        {
          path: "system",
          //系统管理
          name: "system",
          children: [
            {
              path: "menuList",
              name: "menuList",
              component: () => import("../views/admin/System/MenuList.vue"),
            },
          ],
        },
        {
          path: "users",
          //用户管理
          name: "users",
          children: [
            {
              path: "userList",
              name: "userList",
              component: () => import("../views/admin/Users/UserList.vue"),
            },
          ],
        },
      ],
    },
  ],
});

export default router;
