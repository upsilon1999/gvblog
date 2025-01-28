## UI框架的选取

Vue3中个人接触较多的是element系列。
React中则是antd。
由于该网站需要提供Vue3和React两个版本，所以选择了字节开源的同时支持这两种前端框架的arco-design。

```sh
npm install --save-dev @arco-design/web-vue
```

采用全局引用的方式

```ts
import { createApp } from 'vue'
import ArcoVue from '@arco-design/web-vue'
import App from './App.vue'
import '@arco-design/web-vue/dist/arco.css'

const app = createApp(App)
app.use(ArcoVue)
app.mount('#app')
```

### arco-icon

Arco图标是一个独立的库，需要额外引入并注册使用。

```ts
import { createApp } from 'vue'
import ArcoVue from '@arco-design/web-vue'
// 额外引入图标库
import ArcoVueIcon from '@arco-design/web-vue/es/icon'
import App from './App.vue'
import '@arco-design/web-vue/dist/arco.css'

const app = createApp(App)
app.use(ArcoVue)
app.use(ArcoVueIcon)
app.mount('#app')
```

## 设计views目录

由于本系统分为两层

```sh
admin 中后台管理系统
web   前台博客系统
```
