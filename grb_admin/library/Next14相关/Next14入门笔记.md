## 单元重点

```sh
1.使用Next构造全栈程序
```

## 服务端渲染

**服务端渲染和客户端渲染比较**

```sh
【客户端渲染】
1.HTML等使用js在客户端上呈现，网页在用户机子上生成

【存在问题】
1.初始页面加载缓慢
a.因为初始要加载的js包需要下载且可能很大
b.大型应用程序有很多数据需要在组件挂载时获取

2.SEO优化差
因为页面需要等js渲染,可能搜索到空白页面

3.适用于不怎么需要SEO，考虑高度互动的应用，例如后台应用程序

【优势】
1.高度互动的用户体验
a.因为初始页面加载完成后，后续请求特别少

```

```sh
【服务端渲染】
1.HTML模板在web服务器上预先生成，服务器发送已经生成的网站

【优势】
1.初始页面加载快
a.需要下载的js内容少
b.每个页面要使用的数据在渲染之前就会被下载到服务器上

2.搜索引擎优化友好，SEO友好
3.内容导向的合适，例如电商网站、博客等

【劣势】
1.互动性差，因为从一个页面导航到另一个页面，每次都可能要服务器渲染，导致浏览器加载整页

【分类】
1.静态站点生成
HTML是在构建时生成的，一旦完成网站开发，将被导出为静态的HTML、css、js，然后可以将其部署在web服务器上，不是总需要生成。
一般服务器只发送一次页面。

2.动态站点
每次接收到请求，服务器实际生成新的html，所以基本上，他会为每一个用户生成新页面，当底层数据变化时将非常有用
```

**理解**

服务端渲染就是之前那种前后端不分离模式。

一个显著区别,当网速慢时

```sh
【客户端渲染】
一直处在加载状态，页面出不来

【服务端渲染】
页面很早就出来了，但是事件都不会生效，因为js文件还在下载
```



## Next.js

官方称

```sh
属于React的web框架
```

**理解**

```sh
1.是一个构建在React上的框架，所以React的一切，他都可以使用。
2.在Pure React之上添加了路由和数据获取。
3.集成了服务器，方便开发全栈程序。
```

Next有很多固执的规则，可以说是样板，如果团队都采用Next，将使用很方便。

**Next带来的关键要素**

```sh
1.服务器渲染
Next支持静态和动态渲染

2.开箱即用的路由器
基于文件系统的约定定义了路由路线，就如我们要新建一条路由规则，只需要在建立一个文件夹而不需要编写代码
Next有两个路由器，分别是应用程序路由器`App Router`和页面路由器`Pages Router`
【App Router】
React团队想要实现的全栈架构路由器，揉合了React服务器、服务器操作、Streaming等等
2023年推出稳定版
优点：
a.可以使用fetch直接获取数据
b.实现布局、错误处理等十分容易
c.允许高级的路由模式，例如拦截路由、并行路由
d.更好的开发者和用户体验，尤其是服务器组件的存在

缺点：
a.缓存问题令人困惑
b.相比Pages Router学习路线陡峭

【Pages Router】
传统路由器
从Next诞生就在使用，现在官方已不再推荐，但仍然会更新与支持
更容易学习，很多旧项目还在用
缺点：
a.构建布局路由这种基本功能写起来很混乱
b.请求API时要用到很多Next独有的方法，例如getStaticProps和getServerSideProps

3.获取和改变数据
通过React server等组件可以直接从服务器获取和改变数据

4.提供了大量优化技术
针对图片、字体、SEO、预加载等等提供了大量优化技术。
```

### 创建next项目

```sh
npx create-next-app@latest
# 指定版本
npx create-next-app@14
```

然后按照提示一步步操作

```sh
What is your Project named? 项目名称
#左右键选择，enter键确定
Would you like to use TypeScript?是否使用Ts
Would you like to use ESlint?是否使用ESlint
Would you like to use Tailwind css?是否使用Tailwind css
# 一般选否
Would you like to use `src/` directory?
# 现在都推荐
Would you like to use App Router?
# 设置资源根目录
Would you like to customize the default import alias(@/*)?
```

重要目录结构

```sh
【app】
如果我们选Pages Router就会生成一个Pages文件夹
App router就是一个app目录，我们可以删除其下多余的文件，只留下`page.js`，这是根路由

【next.config.mjs】
next配置文件
```

### 路由的创建

例如我们想要访问`localhost:3000/city`，我们不需要写复杂的路由规则，只需要按照以下的步骤

```sh
1.在app下创建city目录
2.在city目录下创建page.js
```

**理解**

```sh
/city 
【/】就是app目录
【/city或/city/】就是app下的city目录

page.js
每个路由的入口组件，我们在里面书写React组件，最后抛出的就是访问路由时呈现在页面上的组件
```

**举例**

```jsx
//组件的名称并不重要，一切以抛出的为主
export default function House() {
  return <div>这里是城市路由</div>;
}
```

**延申**

如果要实现`localhost:3000/city/test`,就继续在city下新建test目录并书写page.js

**技巧**

设置vscode的标签名，设置--`custom label`，针对` Custom Labels: Patterns`添加

```sh
【项key】
**/app/**/page.js

【值value】
Page: ${dirname}
```

这样是为了提高Next目录的可读性。

#### 公共路由组件

有时候我们想要一个组件在多个路由组件之间共用，我们当然可以写在任意地方然后引入，但是为了规范起见，我们一般会在app目录下新建一个components目录，由于这个目录没有`page.js`,所以不会被当作路由访问，我们可以书写一些公共部分，例如

```jsx
import Link from "next/link";

export default function Navigation() {
  return (
    <ul>
      <li>
        <Link href="/">首页</Link>
      </li>
      <li>
        <Link href="/about">关于页</Link>
      </li>
    </ul>
  );
}
```

使用

```jsx
import Navigation from "../components/Navigation";

export default function page() {
  return (
    <div>
      <Navigation />
      这是关于页面
    </div>
  );
}
```



### 路由跳转

**a标签的方式**

和我们之前探讨React Router一样，这种方式能实现跳转，但是会触发页面的重载,需要完全重载速度慢

```jsx
export default function Home() {
  return (
    <div>
      Hello
      <a href="/about">关于页</a>
    </div>
  );
}
```

#### Link

和a标签写法类似，跳转速度快，和React Router的区别，去往哪里使用的是href而不是to

```jsx
import Link from "next/link";

export default function Home() {
  return (
    <div>
      Hello
      <Link href="/about">关于页</Link>
    </div>
  );
}
```

Next在背后进行了一些优化，即使是服务端渲染也可以有和单页面应用一样的体验

### 全局layout

在app目录下和`page.js`同级会默认生成一个`layout.js`,这是Next推荐我们使用根布局，这个文件实际上替代了之前项目中的`index.html`的功能，所以里面声明的组件必须包含`html`和`body`,例如

```jsx
import Navigation from "./components/Navigation";

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        <Navigation />
        <main>{children}</main>
        <footer>Copyright by upsilon</footer>  
      </body>
    </html>
  );
}
```

children就是路由组件出口，它是由Next自动传入的。

我们都知道HTML文件通常包括html、body、head，在Next中head不需要我们在`layout.js`组件中写出,而是声明一个metadata对象，例如

```jsx
import Navigation from "./components/Navigation";

//这个对象是Next规定的，它行使和head标签内meta一样的功能
export const metadata = {
  //修改标题
  title: "my blog",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        <Navigation />
        <main>{children}</main>
        <footer>Copyright by upsilon</footer>
      </body>
    </html>
  );
}
```

### 引入图片

在public下的静态资源文件可以被直接访问，例如

```jsx
export default function Logo(){
    return (
    	<img src="/logo.png" />
    )
}
```

静态资源根路径`/`就是public文件夹

### RSC

全称React Server Components，我们刚才在app里面写的都是这种服务器组件。

**为什么需要RSC**

之前客户端渲染，通过state来控制UI页面，具有响应快、交互好的优点，但是要下载大量js，也可能会请求很多的数据

而服务端渲染，是data来控制UI页面，不需要js、直接操作数据，但是交互差。

所以推出RSC，由state和data共同控制UI页面

**什么是RSC**

```sh
1.用于构建React应用程序的全栈架构
2.服务器作为组件树的一部分
3.只在服务器上运行，不在客户端上运行，用于直接在服务端上获取数据
4.由于只在服务器上运行，所以没有交互性，没有状态，不需要下载js
5.我们像常规客户端组件一样书写后端代码在前端代码旁边
6.默认情况下在React程序中RSC是未激活的，一般需要像Next这样的框架
```

在Next和Redix这种框架中，RSC是默认组件，原本的客户端组件我们可以选择性加入。

在这些框架中如果想要一个组件成为客户端组件，就在文件的顶部加入`use client`

**比较Rsc和客户端组件**

在Next中

```sh
【服务端组件】
1.默认组件
2.不能使用state、hooks
3.没有状态提升，因为没state
4.我们可以通过props连接服务器组件和客户端组件，但是服务器组件的props必须可序列化，不能是函数或类
5.数据获取，直接用async/await就可以
6.可以导入客户端组件和服务端组件，即可以import他们
7.可以渲染客户端组件和服务端组件(Client components and server components)
8.render时机:url更改时

【客户端组件】
1.需要在文件头部加入`use client`
2.可以使用state、hooks
3.有状态提升
4.有props
5.一般推荐用第三方库进行数据获取
6.只能导入客户端组件
7.可以渲染客户端组件和服务端组件传递来的数据(Client components and server components passed as props)
8.render时机：自身状态或父亲状态发生变化时
```

### 获取数据

我们可以直接把服务器组件作为异步函数来获取数据，例如

```jsx
export default async function page() {
  const res = await fetch("https://jsonplaceholder.typicode.com/users");
  const data = await res.json();

  //这实际上是后端获取数据所以打印的位置在后端控制台而不是浏览器的控制台
  console.log(data);
  return <div>这是关于页面</div>;
}
```

**注意**

服务端组件打印的数据将在终端控制台，和nodejs一样

### Next中使用客户端组件

例如我们建造一个计数器组件

```jsx
"use client";
import React, { useState } from "react";

export default function Counter() {
  const [count, setCount] = useState(0);
  return (
    <div>
      <p>{count}</p>
      <button onClick={() => setCount(count + 1)}>加1</button>
    </div>
  );
}
```

在服务器组件中使用

```jsx
import Counter from "../components/Counter";

export default async function Page() {
  const res = await fetch("https://jsonplaceholder.typicode.com/users");
  const data = await res.json();

  //这实际上是后端获取数据所以打印的位置在后端控制台而不是浏览器的控制台
  console.log(data);
  return (
    <div>
      <Counter />
    </div>
  );
}
```

通过研究原理我们会知道客户端组件在浏览器渲染，所以会下载js包，当网速慢时，服务器组件加载完毕，客户端组件的样式也加载完毕（因为无论是客户端组件还是服务端组件的UI都是由服务器渲染好后发给浏览器的），但是事件却要等一段时间才能生效，这是由于`水合反应`需要时间，交互事件则需要浏览器下载好对应的js文件，所以页面早就好了但是事件不生效。

**进一步用porps通信**

我们可以演示在服务器组件中通过props给客户端组件通信，例如

```jsx
import Counter from "../components/Counter";

export default async function Page() {
  const res = await fetch("https://jsonplaceholder.typicode.com/users");
  const data = await res.json();

  //这实际上是后端获取数据所以打印的位置在后端控制台而不是浏览器的控制台
  console.log(data);
  return (
    <div>
      <Counter users={data}/>
    </div>
  );
}
```

客户端组件

```jsx
"use client";
import React, { useState } from "react";

export default function Counter({users}) {
  const [count, setCount] = useState(0);
  console.log(users)
  return (
    <div>
      <p>{count}</p>
      <button onClick={() => setCount(count + 1)}>加1</button>
    </div>
  );
}
```

### 全局loading

在app目录下，与`page.js`同级可以构造一个`loading.js`，这是Next规定的全局加载器

```jsx
export default function Loading() {
  return <div>loading...</div>;
}
```

我们可以书写任意服务器组件内容

## 理解RSC架构底层

How RSC Works Behind the Scenes

RSC vs. SSR How are They Related

P32-017~p32-018

## Suspense

**Suspense是什么**

```sh
1.是React的内置组件，用于捕获或隔离未准备好、未渲染完的组件。
因为这些组件多数时候在执行异步操作，此时我们就称这些组件或其子树都悬空了

在React中他会称这些组件处于暂停状态
```

**是什么导致了组件处于暂停**

```sh
1.异步获取数据，即fetching data
2.延迟加载,即Loading code
```

**基本使用逻辑**

```sh
用Suspense将可能未渲染或未准备好的组件包裹，就是作为他们的兜底父组件
```

**渲染逻辑**

```sh
当React来到未渲染好或者未准备好的组件时，就会向祖先去找Suspense，这就称为边界
当渲染完或准备好之后，组件脱离暂停状态展示自
```