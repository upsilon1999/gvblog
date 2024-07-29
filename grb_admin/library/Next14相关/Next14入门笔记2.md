# 单元重点

```sh
1.Next项目结构
2.优化字体和图片
3.添加metadata和收藏夹图标
4.实现嵌套布局
```

## tips

**常用图标库**

tailwindcss和react经常使用的一个图标库

```sh
npm i @heroicons/react
```



## Next项目结构

### app目录

App Router路由约束会识别app目录生成Next路由。

我们习惯性的将路由组件中可以复用的部分抽离成组件已到了`app/components`目录下，实际上按照Next的路由约束，会帮他也生成一条路由规则，我们可以通过`localhost:3000/components`访问到该目录下的`page.js`,现在我们想让他只成为一个组件文件夹，不时Next路由，解决方案:

```sh
在目录名前用短横杆开头，这样就会被认为是私有目录，不再受Next路由的约束，例如
`app/_components`
```

不存在src目录时，我们会把所有内容都建立到app目录下，例如

```sh
_components 组件目录
_styles 样式目录
_lib  工具类
```

等等，其实都是按照自己的想法命名而已。

**存在src目录**

在建立Next项目时，会询问是否建立src目录，如果建立了src目录，那么app目录将在src目录下，我们就可以在app目录同级建立`components`目录来存放组件，不用担心Next路由问题，因为App Router约束只识别app目录，同理Pages Router约束识别pages目录

**网站图标**

按照Next约定，我们在app目录下的`icon.png`将被识别为网站图标，和其他项目的`favicon`功能类似。

### public目录

用于存放静态资源文件，例如图片，访问举例

```jsx
<img src="/hello.png"/>
```

资源请求的`/`就是public目录

## Next配置路径联想

在Next的根目录下，有一个`jsconfig.json`，在里面配置

```js
{
  "compilerOptions": {
    "paths": {
      "@/*": ["./*"]
    }
  }
}
```

联想路径根据项目配置。

## Next配置tailwindcss

在Next项目创建时选择使用tailwindcss，然后项目建立后如果有css文件直接在`app/layout.js`引入，如果没有css文件，就从tailwindcss官网抄一份，必须包含以下开头即可。

```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

例如我们的`global.css`

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer components {
  .spinner {
    margin: 3.2rem auto 1.6rem;
    width: 60px;
    aspect-ratio: 1;
    border-radius: 50%;
    border: 8px solid theme("colors.primary.900");
    border-right-color: theme("colors.primary.200");
    animation: rotate 1s infinite linear;
  }

  .spinner-mini {
    margin: 0;
    width: 20px;
    aspect-ratio: 1;
    border-radius: 50%;
    border: 2px solid theme("colors.primary.200");
    border-right-color: transparent;
    animation: rotate 1s infinite linear;
  }

  @keyframes rotate {
    to {
      transform: rotate(1turn);
    }
  }
}

/* For data picker */

.rdp {
  --rdp-cell-size: 32px !important;
  --rdp-accent-color: theme("colors.accent.500") !important;
  --rdp-background-color: theme("colors.accent.600") !important;
  margin: 0 !important;
}
```

在`app/layout.js`引入

```js
import "@/app/_styles/globals.css";
```

## Next的metadata

### 标题

在传统项目中有index.html，我们可以编写head的meta信息，在Next项目中我们可以通过metadata对象实现同样的功能

在`layout.js`可以定制全局生效的meta数据，例如

```js
export const metadata = {
  title: "我是标题",
};
```

当我们要给某个某个路由设置自己标题时可以在该路由对应的`page.js`文件中书写

```js
export const metadata = {
  title: "xx路由",
};
```

**只改一部分标题**

我们想要在`layout.js`中书写标题，然后其他路由只能修改特定的部分，

```js
export const metadata = {
  title:{
  	template:"%s 我是layout",
    default:"welcome 你好"
  },
};
```

`%s`是占位符接收其他路由传入的title，例如

```js
export const metadata = {
  title: "xx路由",
};
```

最终标题为

```sh
xx路由 我是layout
```

### 页面描述

这对SEO搜索引擎优化非常重要，格式如下

```js
export const metadata = {
  title: "xx标题",
  desciption:"这是一条神奇的天路，这个页面描述的是xx"
};
```

## Next字体托管

**例如导入谷歌字体**

我们在`layout.js`中实现

```jsx
//这里引入的实际上是一个函数
import {Josefin_Sans} from "next/font/goole"

//将配置对象返回成一个变量
const josefin = Josefin_Sans({
    //设置字体即，例如拉丁、中文
    subsets:["latin"],
    //字体加载方式
    display:"swap"
})

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      {/*在body中设置字体类，*/}
      <body className=`${josefin.className} bg-blue-100`>
        <Navigation />
        <main>{children}</main>
        <footer>Copyright by upsilon</footer>
      </body>
    </html>
  );
}
```

**注意**

```sh
在需要使用字体的标签上加入类名，格式为
`字体类名.className`
```

很诡异的封装方式

## Next图像优化

图像在很大程度上都是页面加载缓慢的根源。

传统的处理方式

```jsx
function App(){
    return <img src="/hello.png"/>
}
```

### Next中的Image组件

Next中的Image组件使用,就是将img替换为Image

```jsx
import Image from "next/image"
function App(){
    return <Image src="/hello.png" height='80' width="60"/>
}
```

这个组件做了三件事

```sh
1.它将自动提供正确的图像尺寸
2.防止布局位移，他迫使我们要指定准确的高度和宽度
3.自动懒加载图像
```

**注意**

如果src指定的是路径就必须指定宽高，正确示范

```jsx
<Image src="/hello.png" height='80' width="60"/>
```

错误示范

```jsx
<Image src="/hello.png" width="60"/>
```

#### Image组件拓展

如果先导入图片，src指定的是引入的对象，不是具体路径就可以不指定宽高

```jsx
import Image from "next/image"
//1.先导入图片
import logo "@/public/logo.png"

function App(){
    return <Image src={logo}/>
}
```

这样的图片无法获得正确尺寸，因为我们没有给尺寸

### Image组件属性

```jsx
import Image from "next/image"
//1.先导入图片
import logo "@/public/logo.png"

function App(){
    /*
    	src 图片位置
    	width 图像宽度
    	height 图像高度
    	alt 提示信息
    	quality 图像质量，就是引入的资源大小，会影响到清晰度，值是1~100，对应1%~100%
    	fill  是否填充整个父元素
    	placeholder 值为blur时是随着加载由模糊到清晰
    */
    return <Image src={logo}/>
}
```

### 处理Image响应式的问题

场景

```sh
1.我们的图像宽高要响应式变化，所以不能指定Image宽高
2.我们的图像是从数据库获取的，无法使用import导入
```

此刻既想要使用Image组件，又想要使用响应式的解决方案

```sh
第一步：给Image加上fill属性，和自定义类实现图像cover
第二个：给Image一个父元素，父元素要具有宽高，可以是响应式百分比
```

案例

```jsx
import Image from "next/image"

function App(){
    return (
        <div style={{height:"100%",width:"100%"}}>
        	<Image src='/logo.png' fill style={{objectFit: cover}}/>
        </div>
    )
}
```

## Next嵌套布局

所谓嵌套布局实际上就是说在某个路由下设置他自己的layout，例如

```sh
`app/city`
`app/city/cityInfo`
`app/city/cityItem`
```

我们可以在city目录下建立`layout.js`，那样他以及他子组件的`page.js`就会是传入的`children`

```jsx
//children传入的是同级page.js和后代page.js暴露的组件
export default function Layout({ children }) {
  return (
      <div>
      	{children}
      </div>
  );
}
```

**小结**

```-sh
每个目录下的page.js就是路由展示页
同级还可以有一个layout.js，就是该路由页面的布局，通过children属性来传入page.js和后代目录的page.js
```

## Next处理数据

34开头的章节

用各种不同方式处理数据，流式数据和缓存，静态及动态渲染。

```sh
1.Suspense和loading.js文件
2.静态和动态渲染
3.Next如何缓存数据，利用各种缓存机制来提升网站性能
```

### Next服务器数据

对于服务器组件，直接从服务器请求数据渲染，无需使用state管理

```jsx
export default async function Page() {
  //直接从服务器获取数据
  const cabins = await getCabins();
    
  return (
    <div>
      <h1 className="text-4xl mb-5 text-accent-400 font-medium">
        Our Luxury Cabins
      </h1>
      <p className="text-primary-200 text-lg mb-10">
        Cozy yet luxurious cabins, located right in the heart of the Italian
        Dolomites. Imagine waking up to beautiful mountain views, spending your
        days exploring the dark forests around, or just relaxing in your private
        hot tub under the stars. Enjoy nature&apos;s beauty in your own little
        home away from home. The perfect spot for a peaceful, calm vacation.
        Welcome to paradise.
      </p>

        <!--实际上只有这一部分需要在数据加载时被挂起-->
      	<div className="grid sm:grid-cols-1 md:grid-cols-2 gap-8 lg:gap-12 xl:gap-14">
          {cabins.map((cabin) => (
            <CabinCard cabin={cabin} key={cabin.id} />
          ))}
    	</div>
    </div>
  );
}
```

### Next设置环境变量

在Next的根目录下设置`.env.local`

```sh
# 格式大写，多个单词采用下划线分割
SUPABASE_URL = https://zgry1949.fun
```

使用环境变量，格式`process.env.环境变量名`

```js
console.log(process.env.SUPABASE_URL)
```

可以供浏览器使用的全局变量

```sh
# 带了NEXT_PUBLIC_前缀的全局变量就可以供所有人使用
NEXT_PUBLIC_SOME_VAR = 23
#没有带该前缀就只能让服务器使用，浏览器无法访问
SUPABASE_URL = https://zgry1949.fun
```

### Next的loading约束

在和page.js同级建立的loading.js文件就是当page.js尚未加载好时的loading组件。例如

```jsx
export default function Loading() {
  return (
      <div>
      	loading data...
      </div>
  );
}
```

### 了解React中的Suspense

34-005

**基本概念**

```sh
Suspense时React的内置组件，允许我们捕获或隔离未准备好的组件，就像try-catch中的catch一样。

当组件未渲染完成时，就处于一种suspending状态。
```

**什么导致了组件suspending**

```sh
1.异步获取数据，fetching data
2.懒加载时，即loading code
```

以往我们都是用三元判断或者isLoading来处理正在加载的组件，有了Suspense之后，我们只需要用该组件包裹即可

**注意**

实际上手动实现Suspense的功能是很困难的，所以一般就交给第三库或者框架集成了，例如

```sh
React Query、Next、Redix
```

#### 使用Suspense

我们之前使用loading.js，这样的话整个组件都将被阻塞，但我们想要的结果是只有请求数据的地方被挂起，例如

```jsx
import CabinCard from "@/app/_components/CabinCard";
import { getCabins } from "../_lib/data-service";
import CabinList from "../_components/CabinList";

export const revalidate = 3600;
// export const revalidate = 15;

export const metadata = {
  title: "Cabins",
};



export default async function Page() {
  const cabins = await getCabins();
    
  return (
    <div>
      <h1 className="text-4xl mb-5 text-accent-400 font-medium">
        Our Luxury Cabins
      </h1>
      <p className="text-primary-200 text-lg mb-10">
        Cozy yet luxurious cabins, located right in the heart of the Italian
        Dolomites. Imagine waking up to beautiful mountain views, spending your
        days exploring the dark forests around, or just relaxing in your private
        hot tub under the stars. Enjoy nature&apos;s beauty in your own little
        home away from home. The perfect spot for a peaceful, calm vacation.
        Welcome to paradise.
      </p>

        <!--实际上只有这一部分需要在数据加载时被挂起-->
      	<div className="grid sm:grid-cols-1 md:grid-cols-2 gap-8 lg:gap-12 xl:gap-14">
          {cabins.map((cabin) => (
            <CabinCard cabin={cabin} key={cabin.id} />
          ))}
    	</div>
    </div>
  );
}
```

处理思路

```sh
1.将需要被挂起的部分单独封装一个组件
2.然后将该封装的组件引入原位置并用Suspense包裹
```

封装组件

```jsx
import CabinCard from "@/app/_components/CabinCard";
import { getCabins } from "../_lib/data-service";

async function CabinList() {
  // noStore();

  const cabins = await getCabins();

  if (!cabins.length) return null;

  return (
    <div className="grid sm:grid-cols-1 md:grid-cols-2 gap-8 lg:gap-12 xl:gap-14">
      {cabins.map((cabin) => (
        <CabinCard cabin={cabin} key={cabin.id} />
      ))}
    </div>
  );
}

export default CabinList;
```

引入并用Suspense包裹，

```jsx
//引入Suspense
import { Suspense } from "react";
import CabinList from "../_components/CabinList";
//在加载时呈现的组件
import Spinner from "../_components/Spinner";

export const revalidate = 3600;

export const metadata = {
  title: "Cabins",
};

export default function Page() {
  return (
    <div>
      <h1 className="text-4xl mb-5 text-accent-400 font-medium">
        Our Luxury Cabins
      </h1>
      <p className="text-primary-200 text-lg mb-10">
        Cozy yet luxurious cabins, located right in the heart of the Italian
        Dolomites. Imagine waking up to beautiful mountain views, spending your
        days exploring the dark forests around, or just relaxing in your private
        hot tub under the stars. Enjoy nature&apos;s beauty in your own little
        home away from home. The perfect spot for a peaceful, calm vacation.
        Welcome to paradise.
      </p>

      <!--fallback 当内部组件处于挂起状态时用于兜底渲染的组件-->
      <Suspense fallback={<Spinner />}>
        <CabinList />
      </Suspense>
    </div>
  );
}
```

### Next动态路由

我们想要实现

```sh
/api/cabin/:cabinId
```

这样的路由规则效果，实现方式，在cabin目录下新建`[cabinId]`目录，然后再这个目录里再建`page.js`，然后我们可以在函数的params参数中获取path参数，即

```jsx
export default function Page({params}){
    //此处的params就是/api/cabin/:cabinId中的cabinId值
    //形式是一个对象，例如{cabinId:90}
    console.log(params)
}
```

**理解**

```sh
1.占位符的键就是[键]
2.对应的值可以通过page.js中的函数参数解构params获得
```

### Next动态设置metaData

当我们使用动态路由的时候，我们希望浏览器的标题也能跟随动态路由变化，此时可以使用异步的generateMetadata方法，他的返回值就是新的metaData，他的参数就是路由，所以也可以从参数上解构出params，例如

```jsx
export async function generateMetadata({params}){
    console.log(params)
    //我们可以在该函数中发起异步请求
    const {name} = await getCabin(params.cabinId)
    return {
        title:`Cabin ${name}`
    }
}
```

### 错误处理

#### 全局错误边界

主要用于处理以下几个问题

```sh
1.访问了不存在的界面，例如/api/cabin/900，出发了notfound报警
2.读取不存在的内容，常常是某个对象上不存在某属性，即常见的undefined问题
```

在app的根目录下设置一个`error.js`，这也是Next的约束，用于捕获全局报错，注意事项

```sh
1.Error.js必须是客户端组件
```

例如

```jsx
"use client"

//函数命名是任意的
//可以解构出两个参数，error 错误信息 reset重置方法
export default function MyError({error,reset}){
    return (
    	<>
        <p>{error.message}</p>
        <button onClick={reset}>重新加载</button>
        </>
    )
}
```

reset方法是内置的，就是重新加载该页面，在很多时候我们会做其他处理，例如重定向到首页

和其他的约束一样，我们可以给某个目录和他的子目录加入自己的`error.js`错误捕获

**拓展**

```sh
1.如果有layout.js,那么error.js会成为layout.js的子组件
2.而global-error.js则会覆盖layout.js
```

### Not Found的处理

虽然`error.js`能处理这个错误，但是有时候我们会想给他一个单独的展示，Next中也提供了类似的约束。

在app根目录下新建`not-found.js`,该文件所设定的组件就是专门用来捕获404的，例如

```jsx
export default NotFound(){
    return (
    	<>
        	<p>不好了，页面走丢了</p>
        </>
    )
}
```

以上是自动触发的，我们也可以在出现错误时手动导向未存在组件，例如

```jsx
import notFound from "next/navigation"

export async function getCabin(id){
    const {data,error} = await 请求
    
    if(error){
        console.log(error)
        //他将调用最近的not-found.js
        notFound()
    }
}
```

### Next的静态和动态渲染

34-011

在Next应用中将由Next来自动选择页面是该静态渲染还是动态渲染，当我们bulid我们的项目时，会得到一个列表，其中对我们的每一个路由都做了解析和划分，前面画圈的代表静态渲染，Lambda的就是动态渲染。

### 动态渲染案例

对于`/api/cabin/[cabinId]`这样的请求，由于`[cabinId]`是不确定的，所以Next采取了动态渲染策略。

我们可以提供一个值列表，告诉Next有哪些可能值，Next就能采用静态渲染。例如在`[cainId]/page.js`

```js
//该方法的作用就是告诉Next，对于该动态路线该页面实际存在哪些参数
export async function generateStaticParams(){
    //我们不建议写一个固定的数据，因为未来一切是可变的
    
    
    //相当于从接口知道我们会有哪些参数
    const cabins = await getCabins();
    //注意事项:规定对象的值是string类型
    const ids = cabins.map(cabin=>{cabinId:String(cabin.id)})
    
    //最终结果得到一个[{cabinId:xx},{cabinId:xx}]的形式
    return ids
}
```

## Next静态站点生成

一般的Next项目构建

```sh
next build
```

如果我们要实现静态站点生成，还需要配置Next的配置项，在`next.config.mjs`下配置

```js
/** @type {import('next').NextConfig} */
const nextConfig = {
   //要实现静态站点生成，需要加入这一项
  output: "export",
};

export default nextConfig;
```

然后再执行`next build`,将在根目录下生成一个`out`目录，我们可以拿这个目录去部署，就和其他项目打包生成的dist目录功能一样

### 预渲染

介于动态渲染和动态渲染之间，就是先提供静态渲染，然后对要变化的部分进行动态渲染，(目前还在研发阶段)

## Next缓存数据

**什么是缓存**

```sh
对已获取或计算好的数据进行临时存储，当我们再次请求同样的数据时，直接从缓存中读取，不需要重新获取或计算。
```

**Next的缓存**

```sh
1.Next在用户浏览器上缓存非常积极，基本上获取的数据、访问过的路由都被缓存了
2.Next提供API重新验证不同的缓存，重新验证只是从缓存中删除所有数据，并用新数据更新他，即重新从来源处再次请求和计算数据。
```

**激进缓存的危害**

缓存本身是个好东西，但是Next滥用缓存导致页面数据陈旧，有些Next缓存甚至无法关闭，这种激进的缓存方式让Next饱受诟病。

由于有太多的API影响缓存，所以控制Next的缓存是个困难的事情。

### Next的四大缓存机制

其中三个缓存，以便他们将数据存储在服务器上，这些被称为请求记忆(request memoization)、数据缓存(data cache)、完整路由缓存(full route cache)

还有一个缓存在客户端，被称为路由器缓存(router cache),也称为客户端缓存。

```sh
【请求记忆(request memoization)】
储存数据:缓存已获取数据的技术，例如Data fetched with similar GET requests(same url and options in fetch function)

多久: One page request(one render,one user),即在一名用户的一项请求的生命周期内，换而言之，数据缓存仅在一页渲染期间。

举例:我们有五个组件请求同一个接口获取相同数据(请求url和参数都相同)，那么只会请求一次

注意:他只适用于React组件
```

```sh
【数据缓存(data cache)】
数据:储存在特定路由或单个fetch请求中获取的所有数据

多久:在我们重新验证前永久存在

举例:即使重新部署也存在，他经常被用于静态页面渲染。
```

```sh
【完整路由缓存(full route cache)】
数据:存储整个html静态页面和rec载荷

多久:直到数据缓存失效或者app重新加载


作用:基本上充当存储机制
```

```sh
【路由器缓存(router cache)】
作用:存储在浏览器中的所有预加载页面以及用户访问过的所有页面
内在逻辑:将所有页面存储在内存中，允许即时或者几乎即时导航

缺点:当用户再次回到这个路由时，不会向服务器发起请求，实际上如果路由时动态的，那么会缓存30秒，如果时静态的，会缓存5min，无法重新验证缓存
```

### 如何更新验证或者退出

这里不会提到具体的代码，具体内容看文档

```sh
【请求记忆(request memoization)】
"重新验证更新"
无法重新验证缓存

"退出缓存"
可以通过fetch函数的终止控制器，即AbortController
```

```sh
【数据缓存(data cache)】
"重新验证更新"
1.可是设置自动过期时间
2.请求特定API
3.手动请求

"退出缓存"
将这个页面排除在缓存之外，同归在page.js中加入revalidate
```

```sh
【路由器缓存(router cache)】
"重新验证更新"
1.revalidatePath or revalidateTag SA
2.通过router.refresh强制刷新
3.重新设置cookie或者删除cookie
```

### 在项目中配置缓存

由于缓存无法在开发模式下生效，所以我们可以采用如下命令来启动Next应用

```sh
npm run build&&npm run start
```

对应命令

```sh
"start":"next start"
"build":"next build"
```

或者我们也可以加一个指令

```sh
"pord":"next build&&next start"
```

**更新或退出数据缓存**

在需要更新退出数据缓存的`page.js`中加入revalidate

```jsx
//引入Suspense
import { Suspense } from "react";
import CabinList from "../_components/CabinList";
//在加载时呈现的组件
import Spinner from "../_components/Spinner";

//设置数据缓存有效期
//值不能是计算得到的，必须是直接写明的
//单位是秒
export const revalidate = 3600;

export const metadata = {
  title: "Cabins",
};

export default function Page() {
  return (
    <div>
      <h1 className="text-4xl mb-5 text-accent-400 font-medium">
        Our Luxury Cabins
      </h1>
      <p className="text-primary-200 text-lg mb-10">
        Cozy yet luxurious cabins, located right in the heart of the Italian
        Dolomites. Imagine waking up to beautiful mountain views, spending your
        days exploring the dark forests around, or just relaxing in your private
        hot tub under the stars. Enjoy nature&apos;s beauty in your own little
        home away from home. The perfect spot for a peaceful, calm vacation.
        Welcome to paradise.
      </p>

      <!--fallback 当内部组件处于挂起状态时用于兜底渲染的组件-->
      <Suspense fallback={<Spinner />}>
        <CabinList />
      </Suspense>
    </div>
  );
}
```

如果我们设置

```jsx
//这就代表不进行数据缓存，也就相当于退出
export const revalidate = 0;
```

**组件级别使缓存失效**

以上是页面级别的，如果我们要在组件级别，这种方式是不生效的，组件级别的处理

```jsx
//1.引入noStore
import { unstable_noStore as noStore } from "next/cache";

import CabinCard from "@/app/_components/CabinCard";
import { getCabins } from "../_lib/data-service";

async function CabinList({ filter }) {
    //2.在组件开头使用noStore方法
  noStore();

  const cabins = await getCabins();

  if (!cabins.length) return null;

  let displayedCabins;
  if (filter === "all") displayedCabins = cabins;
  if (filter === "small")
    displayedCabins = cabins.filter((cabin) => cabin.maxCapacity <= 3);
  if (filter === "medium")
    displayedCabins = cabins.filter(
      (cabin) => cabin.maxCapacity >= 4 && cabin.maxCapacity <= 7
    );
  if (filter === "large")
    displayedCabins = cabins.filter((cabin) => cabin.maxCapacity >= 8);

  return (
    <div className="grid sm:grid-cols-1 md:grid-cols-2 gap-8 lg:gap-12 xl:gap-14">
      {displayedCabins.map((cabin) => (
        <CabinCard cabin={cabin} key={cabin.id} />
      ))}
    </div>
  );
}

export default CabinList;
```

组件级别的处理方式

```sh
1.从Next中引入noStore方法
import { unstable_noStore as noStore } from "next/cache";

2.在组件头部使用nostore方法
```

