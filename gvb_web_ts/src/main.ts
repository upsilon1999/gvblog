import { createApp } from "vue";
import { createPinia } from "pinia";
//重置默认样式
import "@/assets/base.css";
//引入自定义主题样式，主要是一些全局变量
import "@/assets/theme.css";

import App from "./App.vue";
import router from "./router";

//完整引入arco-design
import ArcoVue from "@arco-design/web-vue";
import "@arco-design/web-vue/dist/arco.css";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(ArcoVue); //使用arco-design
app.mount("#app");
