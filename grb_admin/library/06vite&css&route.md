# 单元重点

```sh
1.React router React最重要的第三方库，路由库
2.建立一单页面应用
3.css模块化
```

## 使用 vite 建立 React 项目

1.使用 vite 创建程序，可以指定版本 例如

```sh
# 本课我们使用第四版
npm create vite@4
# 也可以使用最新版
npm create vite@latest
```

按照提示依次输入：

```sh
# 例如worldwise
Project name 项目名称

#选择框架，按上下键移动，按enter键确定
# 不仅react实际上Vue3也可以这样
Select a framework

#选择开发语言 我们选js
Select a variant
```

然后就可以进入项目目录，安装依赖了。

### 了解 vite 项目初始结构

和 create-react-app 的区别

```sh
1.index.html不在public文件夹中，而是在根目录下
2.入口文件不再时`src/index.js`而是`src/main.jsx`
3.package.json下的默认启动命令，由`npm start`变成了`npm run dev`
```

### 配置 vite

和 create-react-app 内置了 eslint 规则不同，vite 几乎是空的，所以我们要配置
**配置 eslint**
第一步就是配置 React 适用的 eslint，在终端执行,

```sh
# 这里安装了三个包
npm install eslint vite-plugin-eslint eslint-config-react-app --save--dev
```

安装完成后，在根目录下新建`.eslintrc.json`文件来配置 eslint 规则，主要是在原有规则上拓展 react 规则

```json
{
  "extends": "react-app"
}
```

然后在`vite.config.js`中引入刚才安装的规则，先来看最初状态

```js
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
});
```

然后进行添加

```js
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
// 引入eslint包
import eslint from "vite-plugin-eslint";

export default defineConfig({
  // 在插件数组中调用
  plugins: [react(), eslint()],
});
```

## css 模块

### React 样式设置

React 样式选择五花八门,由于 React 只是一个库，他没有自己的样式准则，所以样式的选择很多，下面我们来梳理一下

|         style 选择          |         在哪书写         |     如何使用     |        作用域        | 基于 |
| :-------------------------: | :----------------------: | :--------------: | :------------------: | :--: |
|   inline css(即内联样式)    |       JSX elements       |    style prop    |     Jsx element      | css  |
| 外部 css 文件(又叫全局样式) |         额外文件         |  className prop  | 引入样式的该文件全局 | css  |
|          css 模块           | 一个额外的文件为每个组件 |  className prop  |         组件         | css  |
|          css-in-js          | 一个额外的文件为每个组件 | 创建一个新的组件 |         组件         |  js  |
|      Utility-first CSS      |       JSX elements       |  className prop  |     Jsx element      | css  |

**特点解读**

```sh
【inline css】
内联样式，一般只影响到该元素或其子元素

【外部css文件】
这个影响很大，当我们在app文件或者main引入后，全局都会受到影响，在大型项目中甚至不知道影响组件的具体样式来源，所以在大型项目中一般不会使用全局样式文件，或者尽量削弱它的影响

【css模块】
与常规css文件非常相似，不同的是我们只为每个组件编写一个css文件，然后改文件中的样式只局限于该组件，以便没有其他组件可以使用它们，这使得组件更加模块化和可重用

【css-in-js】
实际上是在js文件中编写css，它允许我们创建React组件，将样式直接应用于它们，然后像普通组件一样使用

【Utility-first CSS】
例如tailwindcss，使用预定义的实用程序类来定义单个样式，

【UI组件库】
其实还有一种方式就是使用成熟的UI组件库，然后不写样式(个人感觉UI组件库也是要跳整样式的)，例如MUI、chakra UI、Mantine
```

### css 模块的使用规则

css 模块是 React 预设的一种处理方式，所以我们不需要额外安装包，他的处理方式就是,创建一个`xx.module.css`文件，然后在组件中引入，我们一般会遵循如下规范

```sh
1.css模块名和组件名一致
2.css模块一般放在组件同级目录下
```

示例:

新建模块`PageNav.module.css`,

```css
.nav {
  display: flex;
  justify-content: space-between;
}

ul {
  list-style: none;
}

ul li {
  background-color: aqua;
}
```

在组件中使用，例如在`PageNav.jsx`

```jsx
import { NavLink } from "react-router-dom";
// 引入css模块
import styles from "./PageNav.module.css";
function PageNav() {
  return (
    // 使用模块样式
    <nav className={styles.nav}>
      <ul>
        <li>
          <NavLink to="/">去首页</NavLink>
        </li>
        <li>
          <NavLink to="/home">去home</NavLink>
        </li>
        <li>
          <NavLink to="/product">去产品页</NavLink>
        </li>
      </ul>
    </nav>
  );
}
```

> 规则 1：引入的时候必须全名引入

即引入的时候必须将`xx.module.css`全名引入,自动引入时经常忽略 css 后缀,这样会找不到文件

> 规则 2：应该使用类选择器来限制样式

css 模块本身就是 css 文件，所以理论上来说所有 css 规则都是可行的，例如下面这种

```css
//类选择器
.nav {
  display: flex;
  justify-content: space-between;
}

//标签选择器
ul {
  list-style: none;
}

ul li {
  background-color: aqua;
}

//id选择器
#app {
  height: 200px;
}
```

但是存在一个问题，在我们导入模块文件时，标签选择器会立即对组件生效，而除类选择器以外的选择器我们无法控制，这要说到 css 的模块的使用步骤。

所以我们应该用类选择器来控制样式

#### css 模块的使用步骤

准备好模块文件，例如

```css
.nav {
  display: flex;
  justify-content: space-between;
}

.nav ul {
  list-style: none;
}

.nav ul li {
  background-color: aqua;
}
```

我们可以使用默认导入也可以使用局部导入

**默认导入**

默认导入后，使用`导入名.类名`来使用样式

```jsx
import { NavLink } from "react-router-dom";
// 引入css模块
import styles from "./PageNav.module.css";
function PageNav() {
  return (
    // 使用模块样式
    <nav className={styles.nav}>
      <ul>
        <li>
          <NavLink to="/">去首页</NavLink>
        </li>
        <li>
          <NavLink to="/home">去home</NavLink>
        </li>
        <li>
          <NavLink to="/product">去产品页</NavLink>
        </li>
      </ul>
    </nav>
  );
}
```

**局部导入**

局部导入就是将 css 模块的类名进行导入，例如

```jsx
import { NavLink } from "react-router-dom";
// 引入css模块
import { nav } from "./PageNav.module.css";
function PageNav() {
  return (
    // 使用模块样式
    <nav className={nav}>
      <ul>
        <li>
          <NavLink to="/">去首页</NavLink>
        </li>
        <li>
          <NavLink to="/home">去home</NavLink>
        </li>
        <li>
          <NavLink to="/product">去产品页</NavLink>
        </li>
      </ul>
    </nav>
  );
}
```

不过不是很推荐，因为它可能和其他变量的命名冲突，而且代码可读性差

**小结**

通过使用步骤我们也知道了为什么要使用类选择器书写 css 模块，因为我们只能使用它

#### 样式提升

css 模块化的样式实际上都有自生成的前缀，例如`PageNav.module.css`

```css
.nav {
  display: flex;
  justify-content: space-between;
}

.test {
  background-color: bisque;
}
```

在组件中引入并检查

```jsx
import { NavLink } from "react-router-dom";
// 引入css模块
import styles from "./PageNav.module.css";
function PageNav() {
  return (
    // 使用模块样式
    <nav className={styles.nav}>
      <ul>
        <li className={styles.test}>
          <NavLink to="/">去首页</NavLink>
        </li>
        <li>
          <NavLink to="/home">去home</NavLink>
        </li>
        <li>
          <NavLink to="/product">去产品页</NavLink>
        </li>
      </ul>
    </nav>
  );
}
```

我们会发现他的类名实际上为`_test_1606i_11`,当然这也是随机的，就是为了防止此样式对其他文件产生影响，所以以下这种是读不到的

```jsx
function Product() {
  return <div className="test">这是一个产品</div>;
}

export default Product;
```

这里的`div`身上使用了`test`类，全局的自然会对他产生影响，而我们 css 模块的就不会。所以我们产生出了一个想法，就是将 css 模块的样式提升到全局，语法为

```css
:global(类选择器) {
}
```

例如

```css
.nav {
  display: flex;
  justify-content: space-between;
}

:global(.test) {
  background-color: bisque;
}
```

这样的话 test 类就被提升到全局了，就可以使用`className="test"`来访问了，而且

```jsx
<li className={styles.test}>
```

这就会访问不到，因为该类被提升到全局了

#### 全局样式与 css 模块

全局样式就是引入的外部 css 文件，它的样式是对所有 React 组件生效的，可以使用

```jsx
className = "类名";
```

来使用他的样式，常见的全局样式有 main.jsx 和 App.jsx 引入，当然其他任何组件都可以，这就会造成样式污染，这也是 css 模块出现的原因。

**css 模块**
引入模块 css 文件，且会带有随机生成的前缀或者后缀，防止和其他页面产生污染，这时候我们可能会有一个疑问，例如 NavLink 的内置类名是`active`，如果我们采用如下方式

```css
.nav {
  display: flex;
  justify-content: space-between;
}

.active {
  background-color: bisque;
}
```

这样是无法对 NavLink 生效的，因为引入后加了前缀或者后缀，但如果采用以下方式

```css
.nav {
  display: flex;
  justify-content: space-between;
}

:global(.active) {
  background-color: bisque;
}
```

虽然有效，但是对全局产生了污染，所以我们应该采用以下做法

```css
.nav {
  display: flex;
  justify-content: space-between;
}

.nav :global(.test) {
  background-color: bisque;
}
```

这样的话在使用的时候就形成了组件内约束,因为类`.nav`是针对单一模块的

```jsx
import { NavLink } from "react-router-dom";
// 引入css模块
import styles from "./PageNav.module.css";
function PageNav() {
  return (
    // 使用模块样式
    <nav className={styles.nav}>
      <ul>
        <li>
          <NavLink to="/">去首页</NavLink>
        </li>
        <li>
          <NavLink to="/home">去home</NavLink>
        </li>
        <li>
          <NavLink to="/product">去产品页</NavLink>
        </li>
      </ul>
    </nav>
  );
}
```

**小结**
在`xx.moudule.css`中使用

```css
:global(类选择器) {
}
```

实际上的作用就是使用全局类名，受到全局的影响，而不会加上随机的前后缀

## 概念：路由和单页面应用

### 路由 route

```sh
1.使用路由，基本上可以将不同的url匹配到用户界面的不同视图,在React中也就是将每个URL匹配到特定的React组件，当url被访问时将呈现相应的React组件
2.用户只需要使用链接或者浏览器url就可以在应用程序的不同屏幕之间导航
3.路由保证用户界面和当前浏览器url有着良好的同步
4.路由是构造单页面应用的必要组成部分
```

大多数前端框架都内置了路由库，但是 React 不是框架，它的路由是依托第三方库的，我们使用的最多的是 React Router

### 什么是单页面应用 SPA

**特点 1**
完全在 client 上运行的 Web 应用程序

**特点 2**
单页面应用严重依赖路由，不同的 URL 对应不同视图，一般按以下步骤进行：

```sh
1.当用户点击路由器提供的特殊链接时，
2.浏览器中的URL改变(注：只是改变)
3.React router监听到浏览器url变化，触发DOM更新
```

**特点 3**
由 js(也就是 React)触发页面 DOM 更新

**特点 4**
和传统跳转页面重新加载 html 不同，单页面应用通过 js 更新页面内容，所以页面永远不会硬重载

**特点 5**
看起来像一个移动应用

## 路由的使用

在 React 项目中,路由组件一般会被抽离到一个单独的文件夹，例如 pages 或 routes，目的仅仅是为了与其他组件进行区分

**安装路由库**
我们要注意 React router 的版本，随着 React 版本的变化，这个库的版本也有变化，React18 适用第六版

```sh
npm i react-router-dom@6
```

### 路由规则 Route

我们知道路由就是根据不同的 url 呈现不同的 React 组件，要呈现在页面的什么地方，这就是占位符所做的事情

```jsx
import { BrowserRouter } from "react-router-dom";
import { Routes } from "react-router-dom";
import { Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import Product from "./pages/Product";

function App() {
  return (
    // 路由模式：分为hash路由和history路由,BrowserRouter是history路由模式
    <BrowserRouter>
      {/* Routes用于包裹所有的声明式路由块，目的就是规范化 */}
      <Routes>
        {/* 
			具体的每条路由规则 
			path:访问时的路径，用当前组件的路由路径作为首路径拼接
			element:要呈现的jsx，其实和正常React组件return的返回值一样
		*/}
        <Route path="/" element={<p>欢迎</p>}></Route>
        {/* 直接呈现组件,现在这种方式已经不行了 */}
        <Route path="home" element={HomePage}></Route>
        {/* 这种形式方便传props */}
        <Route path="product" element={<Product />}></Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
```

被`<BrowserRouter/>`包裹的部分就是受路由控制的部分，根据不同的路由规则呈现不同的路由组件，当我们不想让内容受路由干扰时可以

```jsx
import { BrowserRouter } from "react-router-dom";
import { Routes } from "react-router-dom";
import { Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import Product from "./pages/Product";

function App() {
  return (
    <div>
      {/* 不受路由控制，无论路由如何跳转都对他无影响 */}
      <h1>Hello Router</h1>
      {/* 用于呈现路由组件 */}
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<p>欢迎</p>}></Route>
          <Route path="home" element={<HomePage />}></Route>
          <Route path="product" element={<Product />}></Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}
```

让我们来详细理解一下这部分

```sh
【BrowserRouter】
路由器，所有路由相关的内容，或者通俗点说`react-router-dom`引出的内容都应该在路由器中，否则会报超出范围
规划出一片受路由控制的区域，整个项目只有一个，他在哪开辟，那个组件就是根路径
从main.jsx规定了我们的根组件App，然后在该组件通过路由来呈现不同子组件

【Routes】
路由规则区
所有的路由规则必须写在其中

【Route】
路由规则
同时也是路由组件占位符，根据不同的规则，将自身替换成路由组件
```

### 路由跳转 Link

我们可以尝试用不同的方式实现跳转，最自然的就是使用 a 标签，例如

```jsx
import { BrowserRouter } from "react-router-dom";
import { Routes } from "react-router-dom";
import { Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import Product from "./pages/Product";

function App() {
  return (
    <div>
      {/* 不受路由控制，无论路由如何跳转都对他无影响 */}
      <a herf="/home">Hello Router</a>
      {/* 用于呈现路由组件 */}
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<p>欢迎</p>}></Route>
          <Route path="home" element={<HomePage />}></Route>
          <Route path="product" element={<Product />}></Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}
```

经过验证，a 标签可以实现页面的跳转，但是我们会发现它是以传统的方式重载了整个页面,所以不符合单页面的设想，为此 React router 提供了用于跳转的标签

```jsx
import { Link } from "react-router-dom";
function HomePage() {
  return (
    <div>
      首页
      {/* to 就是要前往的路由，这里使用了绝对路径，从根路由出发 */}
      {/* 他是对a标签的封装，底层实际上就是a标签 */}
      <Link to="/product">去产品页</Link>
    </div>
  );
}

export default HomePage;
```

Link 标签就是封装过的 a 标签
**理解**

```sh
1.除了BrowserRouter是要出现在根组件将该组件变成根路由(当然也可以把其他组件设为根路由，不过那样很容易混乱)
2.路由占位符Route出现在根路由所在组件
3.路由跳转链接可出现在任意组件(普通组件，路由组件都可以),因为它本质是个a标签
```

#### NavLink

有时我们想知道当前处于哪个路由标签，应该高亮当前路由，这时候就可以使用 NavLink，例如

```jsx
import { NavLink } from "react-router-dom";
function PageNav() {
  return (
    <nav>
      <ul>
        <li>
          <NavLink to="/">去首页</NavLink>
        </li>
        <li>
          <NavLink to="/home">去home</NavLink>
        </li>
        <li>
          <NavLink to="/product">去产品页</NavLink>
        </li>
      </ul>
    </nav>
  );
}

export default PageNav;
```

基本功能与 Link 一样，就是给当前路由链接加了一个高亮样式(注:只是给底层 a 标签加了一个 active 样式名，具体样式要我们自己实现)

### 注意事项

像 Link、NavLink 这种路由相关标签，必须出现在 BrowserRouter 这种路由器包裹中，否则会提示超出范围使用，例如

```jsx
//这是一个普通组件
import { NavLink } from "react-router-dom";

function PageNav() {
  return (
    <nav>
      <ul>
        <li>
          <NavLink to="/">去首页</NavLink>
        </li>
        <li>
          <NavLink to="/home">去home</NavLink>
        </li>
        <li>
          <NavLink to="/product">去产品页</NavLink>
        </li>
      </ul>
    </nav>
  );
}
```

使用该组件

```jsx
function App() {
  return (
    <div>
      <PageNav />
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<p>欢迎</p>}></Route>
          <Route path="home" element={<HomePage></HomePage>}></Route>
          <Route path="product" element={<Product />}></Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
```

报错：我们不该超范围使用 NavLink 标签，使用 NavLink 标签的组件应该位于路由器（例如 BroserRouter）控制区内，即

```sh
1.NavLink、Link本身可以使用在任何组件中，它就是a标签
2.但是NavLink、Link必须受路由器（例如BroserRouter）控制，所以合法位置是路由组件及其子组件或者路由器直属
```

正确示范:

```jsx
function App() {
  return (
    <div>
      <BrowserRouter>
        {/*BrowserRouter的控制区中*/}
        <PageNav />
        <Routes>
          <Route path="/" element={<p>欢迎</p>}></Route>
          <Route path="home" element={<HomePage></HomePage>}></Route>
          <Route path="product" element={<Product />}></Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
```

路由组件，原因是，路由组件本身就处于路由器（例如 BrowserRouter）的控制区中

```jsx
function HomePage() {
  return (
    <div>
      <PageNav />
    </div>
  );
}
```

**小结**

实际上路由组件和普通组件没差别，就是因为它充当了路由跳转的目标而已，

所以可以理解为所有于路由相关的标签、组件都必须要受`BrowserRouter`这种路由器的控制

### 嵌套路由

目的：随着路由 url 改变，页面的一部分发生变化，有些地方不变

#### 嵌套路由规则的书写

```jsx
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<p>欢迎</p>}></Route>
        <Route path="product" element={<Product />}></Route>
        {/* 嵌套路由的规则就是在Route的便签体中写入嵌套路由的路由规则Route */}
        <Route path="app" element={<AppLayout />}>
          <Route path="countries" element={<CountryList />} />
          <Route path="form" element={<Form />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}
```

#### 嵌套路由的呈现

按照路由规则，现在路由 app 下有子路由 countries 和 form，当子路由匹配成功后应该要在父路由组件中呈现子路由组件，此时就要用到 Outlet 标签，

```jsx
import { Outlet } from "react-router-dom";
function AppLayout() {
  return (
    <div>
      <Logo />
      {/* 子路由组件占位符，当子路由被匹配到时在这里呈现 */}
      <Outlet />
    </div>
  );
}
```

#### 索引路由

对于嵌套路由，当回到父路由时，由于没有任何子路又被匹配，Outlet 占位处不会有任何呈现，我们想要设置一个默认的子路由，此时就可以使用索引路由

```jsx
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<p>欢迎</p>}></Route>
        <Route path="product" element={<Product />}></Route>
        {/* 嵌套路由的规则就是在Route的便签体中写入嵌套路由的路由规则Route */}
        <Route path="app" element={<AppLayout />}>
          {/* 索引路由，即默认子路由 */}
          <Route index element={<Form />} />
          <Route path="countries" element={<CountryList />} />
          <Route path="form" element={<Form />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}
```
#### 导航组件Navigate
我们在嵌套路由中，使用了根组件默认跳转，实际上应该配合重定向来使用，以达到刷新的作用
```jsx
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<p>欢迎</p>}></Route>
        <Route path="product" element={<Product />}></Route>
        {/* 嵌套路由的规则就是在Route的便签体中写入嵌套路由的路由规则Route */}
        <Route path="app" element={<AppLayout />}>
          {/* 索引路由，即默认子路由 */}
          {/* 使用重定向跳转来刷新 */}
          <Route index element={<Navigate replace to="form"/>} />
          <Route path="countries" element={<CountryList />} />
          <Route path="form" element={<Form />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

```
这个组件一般用来做强制导航和重定向
### Url存储状态

URL 也是存储状态的极佳位置，原因:

```sh
1.在URL中放置状态时存储状态的一种简单方法，在应用程序中的所有组件都可以轻松访问的全局位置
2.是一种将数据从一页传递到下一页的好方法，而不必将数据存储在应用程序的某个临时位置
3.使得书签或者使用页面具有确切UI状态共享页面，例如我们分享一个网址给别人时，例如购物，如果URL中已经带有颜色尺码信息，她获取到链接访问的时候就会自动选中
```

所谓 URL 存储状态，实际上就是如下形式

```sh
www.zgry.fun?name=zhansan&age=20
www.zgry.fun/185289
```

其实就是路由传参

#### path 传参

**书写路由规则**

路由规则的书写,用`:变量`来书写占位符

```jsx
<Routes>
    <Route path="app/:id/:name" element={<AppItem/>}></Route>
<Routes>
```

**参数传递**

```jsx
function City(props){
  return (
    {/* 模板字符串需要js解析 */}
    <Link to={`app/${props.id}/${props.name}`} />
  )
}
```

**参数接收**
接收 path 参数，需要使用到 useParams 钩子，他会返回一个对象，对象的键就是 path 传参的占位符，值也就是对应的接收值，例如我们实际调用传入

```
www.zgry.fun/app/1854/zhansan
```

则接收为

```jsx
import { useParams } from "react-router-dom";
function AppItem() {
  // 输出的结果为
  /*
    {id:1854,name:"张三"}
  */
  const x = useParams();
  return <p>...</p>;
}
```

#### queryString 传参

**路由传参**

```jsx
function City(props){
  return (
    {/* 模板字符串需要js解析 */}
    <Link to={`app?id=${props.id}&name=${props.name}`} />
  )
}
```

**参数接收**
使用 useSearchParams 钩子进行接收，该钩子返回的是一个数组，数组第一个元素是参数类 map，第二个参数是更新方法

```jsx
import { useSearchParams } from "react-router-dom";
function AppItem(){
  const [searchParams,setsearchParams] = useSearchParams()
  // 获取对应的参数值
  const id = searchParams.get("id")
  return <p>...</p>;
}
```
更新方法的使用，和useState提供的setter方法很像，用一个新的对象取覆盖点原来的，例如
```jsx
import { useSearchParams } from "react-router-dom";
function AppItem(){
  const [searchParams,setsearchParams] = useSearchParams()
  // 获取对应的参数值
  const id = searchParams.get("id")

  function change(){
    // 注意这里的更新不是我们预想的使用set，而是使用预设的setter方法，传入新值
    setsearchParams({id:4852,name:"lisi"})
  }
  return <p>...</p>;
}
```

### useNavigate

我们以往跳转都需要Link或者NavLink，也就是传统的声明式路由导航，这样局限性很大，所以有了编程式路由导航。使用步骤:
**1.引入钩子创建导航对象**
```jsx
// 引入钩子
import { useNavigate } from "react-router-dom";
function City(){
  //创建导航对象，对象名大家一般都用navigate
  const navigate = useNavigate()
}
```
**2.使用导航对象**
```jsx
import { useNavigate } from "react-router-dom";
function City(){
  //创建导航对象，对象名大家一般都用navigate
  const navigate = useNavigate()

  function toApp(){
    // 这个功能就和Link中to一样，相当于 to="app"
    navigate("app")
  }

  return <p>...</p>
}
```

## css模块补充

当我们的类名有破折号时，例如
```css
.cityItem--active{
    color:red;
}
```

此时使用时应该用

```jsx
import styles from "xx.module.css"
function App(){
    return <p className={styles["cityItem--active"]}>...</p>
}
```

## 使用地图

**强烈推荐，真的很强:**对应P230-237

1.安装基本的库,一个是react-leaflet，一个是leaflet

```sh
npm i react-leaflet leaflet
```

这两个是实现地图的最大的开源库，在使用的时候要学会查看文档，搜索`react-leaflet`,从这个文档再进入`leaflet`的文档从那里获取初始css，这样就不会产生更多问题

2.从react-leaflet引入地图代码，可以深入研究，但记住一定要给地图设置高度

3.react-leaflet不支持onClick封装，他的所有事件都需要用自定义组件来实现，就是将事件逻辑写在组件里，然后组件返回null即可，感兴趣可以研究

### 使用日期选择器包

1.安装常用的日期选择器包

```sh
npm i react-datepicker
```

看文档学习，搜索`react-datepicker`

## 假的身份验证

用户的身份验证通常分三步进行:

```sh
1.从登录表单中获取用户的电子邮件和密码，并使用API进行检查，
2.如果是正确的，校验通过，就将用户重定向到主应用程序，并将用户信息保存到状态中
3.应该保护应用免受未经授权的访问，未登录的用户将功能受限
```

之所以说是假的身份验证，是因为我们第一步没有用API检查，而是硬编码，和已知信息做简单校验

### 1.创建验证Context

创建一个`FakeAuthContext.jsx`

```jsx
import { createContext, useContext, useReducer } from "react";

const AuthContext = createContext();


function AuthProvider({ children }) {
  
  return (
    <AuthContext.Provider>
      {children}
    </AuthContext.Provider>
  );
}

function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined)
    throw new Error("AuthContext was used outside AuthProvider");
  return context;
}

export { AuthProvider, useAuth };
```

简易的架子搭建都是相同的，然后考虑使用的地方，我们是在登录的时候校验，登出的时候去除用户状态,因为用户信息和是否授权几乎是同时更新的，所以推荐用reducer来维护(不嫌麻烦也可以使用两个state)

```jsx
import { createContext, useContext, useReducer } from "react";

const AuthContext = createContext();

//初始值
const initialState = {
  user: null,
  isAuthenticated: false,
};

//reducer逻辑
function reducer(state, action) {
  switch (action.type) {
    case "login":
      return { ...state, user: action.payload, isAuthenticated: true };
    case "logout":
      return { ...state, user: null, isAuthenticated: false };
    default:
      throw new Error("Unknown action");
  }
}

//硬编码数据
const FAKE_USER = {
  name: "Jack",
  email: "jack@example.com",
  password: "qwerty",
  avatar: "https://i.pravatar.cc/100?u=zz",
};


function AuthProvider({ children }) {
  
  //使用useReducer来维护状态
  const [{ user, isAuthenticated }, dispatch] = useReducer(
    reducer,
    initialState
  );  
    
  //登入事件
  function login(email, password) {
    //由于我们是直接比较，没有请求API，所以是假验证
    if (email === FAKE_USER.email && password === FAKE_USER.password){
        dispatch({ type: "login", payload: FAKE_USER });
	}
  }

  //登出事件
  function logout() {
    dispatch({ type: "logout" });
  }
  
  return (
    <AuthContext.Provider value={{ user, isAuthenticated, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined)
    throw new Error("AuthContext was used outside AuthProvider");
  return context;
}

export { AuthProvider, useAuth };
```

**注意事项**

在实际前端开发中，千万不要把重要数据写在代码中，假设:

```jsx
//假设这是真的用户名账号密码
const FAKE_USER = {
  name: "Jack",
  email: "jack@example.com",
  password: "qwerty",
  avatar: "https://i.pravatar.cc/100?u=zz",
};
```

因为前端的代码会被浏览器下载，即使经过简易编码，也还是能识别的(这就是为什么很多前端开发人员区别人的网页爬取校验规则)

### 2.使用Context

**1.Provider位置**

由于是关于登陆组件，甚至有时候某些组件功能会用到用户信息，所以推荐包裹整个路由器

```jsx
function App() {
  return (
    //包裹整个路由器，方便用户信息的使用
    <AuthProvider>
      <CitiesProvider>
        <BrowserRouter>
            <Routes>
              <Route index element={<Homepage />} />
              <Route path="product" element={<Product />} />
              <Route path="pricing" element={<Pricing />} />
              <Route path="login" element={<Login />} />
              <Route path="app" element={<AppLayout />}>
                <Route index element={<Navigate replace to="cities" />} />
                <Route path="cities" element={<CityList />} />
                <Route path="cities/:id" element={<City />} />
                <Route path="countries" element={<CountryList />} />
                <Route path="form" element={<Form />} />
              </Route>
              <Route path="*" element={<PageNotFound />} />
            </Routes>
        </BrowserRouter>
      </CitiesProvider>
    </AuthProvider>
  );
}
```

**2.登录组件执行登陆事件**

```jsx
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Button from "../components/Button";
import PageNav from "../components/PageNav";
import { useAuth } from "../contexts/FakeAuthContext";
import styles from "./Login.module.css";

export default function Login() {
  // PRE-FILL FOR DEV PURPOSES
  const [email, setEmail] = useState("jack@example.com");
  const [password, setPassword] = useState("qwerty");

  //从上下文中获取登录方法和权限判断
  const { login, isAuthenticated } = useAuth();
  const navigate = useNavigate();

  //表单提交事件
  function handleSubmit(e) {
    //去除掉默认事件
    e.preventDefault();

    //如果邮箱密码都存在就执行登录事件
    if (email && password) login(email, password);
  }

  //监听isAuthenticated，如果isAuthenticated为真就触发跳转
  //使用replace的原因，如果不用的话，点击回退会再次来到login组件，触发navigate再次跳到app
  //所以使用replace
  useEffect(
    function () {
      if (isAuthenticated) navigate("/app", { replace: true });
    },
    [isAuthenticated, navigate]
  );

  return (
    <main className={styles.login}>
      <PageNav />

      <form className={styles.form} onSubmit={handleSubmit}>
        <div className={styles.row}>
          <label htmlFor="email">Email address</label>
          <input
            type="email"
            id="email"
            onChange={(e) => setEmail(e.target.value)}
            value={email}
          />
        </div>

        <div className={styles.row}>
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            onChange={(e) => setPassword(e.target.value)}
            value={password}
          />
        </div>

        <div>
          <Button type="primary">Login</Button>
        </div>
      </form>
    </main>
  );
}
```

**3.执行登出事件**

```jsx
import { useNavigate } from "react-router-dom";
import { useAuth } from "../contexts/FakeAuthContext";
import styles from "./User.module.css";

//登出事件在User组件里
function User() {
  //从上下文获取登出事件
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  function handleClick() {
    //执行登出事件，用户信息被重置为null
    logout();
    //跳转到主页，因为我们从主页有按钮跳转到登陆页
    //且主页不需要用户信息
    navigate("/");
  }

  return (
    <div className={styles.user}>
      <img src={user.avatar} alt={user.name} />
      <span>Welcome, {user.name}</span>
      {/* 登出事件 */}
      <button onClick={handleClick}>Logout</button>
    </div>
  );
}

export default User;
```

### 3.保护应用程序免受未经授权的访问

基本操作就是当用户访问未经授权的页面时将他们重定向到特定页面，

```sh
1.例如没登录的用户不该访问某些需要登录信息的页面
2.例如没有VIP权限的用户不该访问需要VIP的页面
```

在项目中我们可以用一个组件来统一处理重定向问题，然后将整个APP包裹在该组件中，例如在路由组件中加入`ProtectedRoute.jsx`组件

```jsx
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../contexts/FakeAuthContext";

function ProtectedRoute({ children }) {
  //从上下文中获取授权信息
  const { isAuthenticated } = useAuth();
  const navigate = useNavigate();

  useEffect(
    function () {
      //如果没有授权，回到导航页
      if (!isAuthenticated) navigate("/");
    },
    [isAuthenticated, navigate]
  );

  //有无授权，有授权展示子组件,否则为null
  return isAuthenticated ? children : null;
}

export default ProtectedRoute;
```

使用该组件

```jsx
function App() {
	return (
		<AuthProvider>
			<CitiesProvider>
				<BrowserRouter>
                    <Routes>
                        <Route index element={<Homepage />} />
                        <Route path="product" element={<Product />} />
                        <Route path="pricing" element={<Pricing />} />
                        <Route path="login" element={<Login />} />
                        {/*将所有需要授权的组件都包裹进去，这里是AppLayout组件*/}
                        <Route
                            path="app"
                            element={
                                <ProtectedRoute>
                                    <AppLayout />
                                </ProtectedRoute>
                            }
                            >
                            <Route
                                index
                                element={<Navigate replace to="cities" />}
                                />
                            <Route path="cities" element={<CityList />} />
                            <Route path="cities/:id" element={<City />} />
                            <Route
                                path="countries"
                                element={<CountryList />}
                                />
                            <Route path="form" element={<Form />} />
                        </Route>
                        <Route path="*" element={<PageNotFound />} />
                    </Routes>
				</BrowserRouter>
			</CitiesProvider>
		</AuthProvider>
	);
}
```

有个跳转的小问题，关于图片的获取，例如

```css
.homepage {
  height: calc(100vh - 5rem);
  margin: 2.5rem;
  background-image: linear-gradient(
      rgba(36, 42, 46, 0.8),
      rgba(36, 42, 46, 0.8)
    ),
    url("bg.jpg");
  background-size: cover;
  background-position: center;
  padding: 2.5rem 5rem;
}
```

我们如果采用

```css
url("bg.jpg");
```

这种形式获取，它是基于当前路径的，在重定向的时候可能会获取到历史路径，例如我们从`localhost:9090/app`重定向到`localhost:9090/`,正确获取应该是

```sh
localhost:9090/bg.jpg
```

但是按照上述写法就有可能读到

```sh
localhost:9090/app/bg.jpg
```

解决方案是改为绝对路径

```css
.homepage {
  height: calc(100vh - 5rem);
  margin: 2.5rem;
  background-image: linear-gradient(
      rgba(36, 42, 46, 0.8),
      rgba(36, 42, 46, 0.8)
    ),
    url("/bg.jpg");
  background-size: cover;
  background-position: center;
  padding: 2.5rem 5rem;
}
```

这样每次都会触发绝对路径拼接，就不会出现那个读取历史的问题