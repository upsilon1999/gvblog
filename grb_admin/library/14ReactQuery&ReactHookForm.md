## 单元重点

```sh
1.样式组件
2.React Query
3.React Hook Form一个表单管理库
4.Supabase 用于构筑后台，存储远程数据
```

## 技巧

一个日期函数库

```sh
npm i date-fns
```

一个提示组件，在我们不想使用UI组件库，又想要漂亮提示是可以使用

```sh
npm i react-hot-toast
```

**获取children上的属性**

```jsx
function App(){
    return (
    	<div>
        	<City>
                <CityItem id="543">
            <City>
        </div>
    )
}
```

对于City组件，CityItem组件是children，如果我们的

## 客户端渲染和服务端渲染

英文单词是Client-side rendering(CSR)和Server-side rendering(SSR)

```sh
【CSR使用纯净react】
1.通常用于构建单页面应用，百分百地呈现在客户端上
2.所有的html都是在用户地浏览器中生成
3.在程序运行之前所有js文件都要下载好，对performance不利，当用户设备差和网不好时影响性能
4.适用于内部使用的不需要搜索优化的

【SSR使用freamework】
1.通常用于构造多页面应用
2.一些HTML在服务器上渲染
3.少量的js需要下载
4.搜索优化的推荐
```

## 样式组件

查看文档

```sh
https://styled-components.com/
```

### 基本使用

**简介**

用js书写样式，返回一个React组件，这个组件包含我们定义的样式，这些样式时有随机类名使他们只在声明该样式组件的地方生效，这样就不会污染全局样式。

**安装**

```sh
npm i styled-components
```

**使用**

基本语法

```js
//1.引入样式组件
import styled from "styled-components";

const React组件名 = styled.html标签名 `模板字符串用于书写样式`
```

样式组件就是基于一个html标签，给他加上特定样式，变成一个独一无二的React组件，例如

```jsx
import styled from "styled-components";

const StyledAppLayout = styled.div`
  display: grid;
  grid-template-columns: 26rem 1fr;
  grid-template-rows: auto 1fr;
  height: 100vh;
`;

function App(){
    return (
        {/*StyledAppLayout本质上还是一个div，只不过加上了一个随机类名，该类包含我们定义的样式 */}
    	<StyledAppLayout>
            <p>...</p>
        </StyledAppLayout>
    )
}
```

**理解**

```sh
1.所谓的样式组件，就是将css与html元素绑定后生成一个可复用的React样式组件
2.用于解决全局css问题
```

**插件**

```sh
为了让代码阅读更方便，推荐安装vscode插件：
vscode-styled-components
```

### 全局样式化组件

就如css模块化的样式可以通过`:global()`提升到全局一样，样式化组件也具备相应能力

**构造全局样式组件**

一般会把全局样式文件放在一起，例如再src目录下新建styles目录，然后建立全局样式组件，命名为`GlobalStyles.js`

```js
//1.引入全局样式组件构建器
import { createGlobalStyle } from "styled-components";

//2.构建全局样式组件
//就是将之前index.css做的全局样式重置书写进去
const GlobalStyles = createGlobalStyle`
:root {
  /* Indigo */
  --color-brand-50: #eef2ff;
  --color-brand-100: #e0e7ff;
  --color-brand-200: #c7d2fe;
  --color-brand-500: #6366f1;
  --color-brand-600: #4f46e5;
  --color-brand-700: #4338ca;
  --color-brand-800: #3730a3;
  --color-brand-900: #312e81;

  /* Grey */
  --color-grey-0: #fff;
  --color-grey-50: #f9fafb;
  --color-grey-100: #f3f4f6;
  --color-grey-200: #e5e7eb;
  --color-grey-300: #d1d5db;
  --color-grey-400: #9ca3af;
  --color-grey-500: #6b7280;
  --color-grey-600: #4b5563;
  --color-grey-700: #374151;
  --color-grey-800: #1f2937;
  --color-grey-900: #111827;

  --color-blue-100: #e0f2fe;
  --color-blue-700: #0369a1;
  --color-green-100: #dcfce7;
  --color-green-700: #15803d;
  --color-yellow-100: #fef9c3;
  --color-yellow-700: #a16207;
  --color-silver-100: #e5e7eb;
  --color-silver-700: #374151;
  --color-indigo-100: #e0e7ff;
  --color-indigo-700: #4338ca;

  --color-red-100: #fee2e2;
  --color-red-700: #b91c1c;
  --color-red-800: #991b1b;

  --backdrop-color: rgba(255, 255, 255, 0.1);

  --shadow-sm: 0 1px 2px rgba(0, 0, 0, 0.04);
  --shadow-md: 0px 0.6rem 2.4rem rgba(0, 0, 0, 0.06);
  --shadow-lg: 0 2.4rem 3.2rem rgba(0, 0, 0, 0.12);

  --border-radius-tiny: 3px;
  --border-radius-sm: 5px;
  --border-radius-md: 7px;
  --border-radius-lg: 9px;

  /* For dark mode */
  --image-grayscale: 0;
  --image-opacity: 100%;
}

*,
*::before,
*::after {
  box-sizing: border-box;
  padding: 0;
  margin: 0;

  /* Creating animations for dark mode */
  transition: background-color 0.3s, border 0.3s;
}

html {
  font-size: 62.5%;
}

body {
  font-family: "Poppins", sans-serif;
  color: var(--color-grey-700);

  transition: color 0.3s, background-color 0.3s;
  min-height: 100vh;
  line-height: 1.5;
  font-size: 1.6rem;
}

input,
button,
textarea,
select {
  font: inherit;
  color: inherit;
}

button {
  cursor: pointer;
}

*:disabled {
  cursor: not-allowed;
}

select:disabled,
input:disabled {
  background-color: var(--color-grey-200);
  color: var(--color-grey-500);
}

input:focus,
button:focus,
textarea:focus,
select:focus {
  outline: 2px solid var(--color-brand-600);
  outline-offset: -1px;
}

/* Parent selector, finally 😃 */
button:has(svg) {
  line-height: 0;
}

a {
  color: inherit;
  text-decoration: none;
}

ul {
  list-style: none;
}

p,
h1,
h2,
h3,
h4,
h5,
h6 {
  overflow-wrap: break-word;
  hyphens: auto;
}

img {
  max-width: 100%;

  /* For dark mode */
  filter: grayscale(var(--image-grayscale)) opacity(var(--image-opacity));
}

`;

//暴露全局样式组件
export default GlobalStyles;
```

**全局样式组件的使用**

要满足以下几点要求

```sh
1.要是其他所有组件的兄弟组件
2.不能是子组件
3.必须是自闭合组件
```

使用举例

```jsx
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";

import GlobalStyles from "./styles/GlobalStyles";
import Dashboard from "./pages/Dashboard";
import Bookings from "./pages/Bookings";
import Cabins from "./pages/Cabins";
import Users from "./pages/Users";
import Settings from "./pages/Settings";
import Account from "./pages/Account";
import Login from "./pages/Login";
import PageNotFound from "./pages/PageNotFound";
import AppLayout from "./ui/AppLayout";

function App() {
  return (
    <>
      <GlobalStyles />
      <BrowserRouter>
        <Routes>
          <Route element={<AppLayout />}>
            <Route index element={<Navigate replace to="dashboard" />} />
            <Route path="dashboard" element={<Dashboard />} />
            <Route path="bookings" element={<Bookings />} />
            <Route path="cabins" element={<Cabins />} />
            <Route path="users" element={<Users />} />
            <Route path="settings" element={<Settings />} />
            <Route path="account" element={<Account />} />
          </Route>

          <Route path="login" element={<Login />} />
          <Route path="*" element={<PageNotFound />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}
```

所以出现位置是App组件，即根组件，然后用`<></>`包裹

### 样式组件的优势

我们需要注意到样式组件实际上是js，书写css的地方是模板字符串，所以我们可以书写js，例如

```js
import styled from "styled-components";

const MyDiv = styled.div`
  width:${10>5?"30px":"20px"};
  height:"40px";
`;

export default MyDiv;
```

#### **使用变量**

```js
import styled from "styled-components";

const test = `text-align:center;`

const MyDiv = styled.div`
  width:${10>5?"30px":"20px"};
  height:"40px";
  ${test}
`;

export default MyDiv;
```

有时也会写成

```js
import styled from "styled-components";

//1.这里这个css前缀是为了让vscode-styled-component能够识别并高亮
//2.当模板字符串中有复杂逻辑，例如判断或其他变量引用时必须要css，否则可能失灵
const test = css`text-align:center;`

const MyDiv = styled.div`
  width:${10>5?"30px":"20px"};
  height:"40px";
  ${test}
`;

export default MyDiv;
```

更多写法

```js
import styled from "styled-components";

//通过截断或三元来动态呈现样式
//
const test = css`
	text-align:center;
	${5>3&&"background-color:red;"}
`

const MyDiv = styled.div`
  width:${10>5?"30px":"20px"};
  height:"40px";
  ${test}
`;

export default MyDiv;
```

#### **样式组件接收props**

既然他是一个react组件，就能接收props，不过这里的props有自定义和as两种

```jsx
function App(){
    return (
    	<MyDiv type="danger"></MyDiv>
    	<MyDiv type="info"></MyDiv>
        <MyDiv type="warnning"></MyDiv>
    )
}
```

在样式组件的模板字符串中可以用回调函数来接收props，例如

```js
import styled from "styled-components";


const MyDiv = styled.div`
  ${props=>
	props.type ==="danger"&&
    css`
    	color:"red";
    	font-size:"20px";
    `
  }
  
  ${props=>
	props.type ==="info"&&
    css`
    	color:"gray";
    	font-size:"16px";
    `
  }
  
  ${props=>
	props.type ==="warnning"&&
    css`
    	color:"orange";
    	font-size:"16px";
    `
  }
`;

export default MyDiv;
```

注意事项

```sh
1.对于这种复杂的选择逻辑必须使用css前缀包裹要呈现的字符串
2.模板字符串内不能有注释，有注释样式组件就会失效
```

**自定义和as的区别**

自定义是设置一个类名，as则是官方提供的一个特定属性用于变更html元素，

```jsx
function App(){
    return (
        {/*MyDiv的底层html从div变成h2 */}
    	<MyDiv type="danger" sizes="hello" as="h2"></MyDiv>
    )
}
```

**声明props的默认值**

```jsx
import styled from "styled-components";


const MyDiv = styled.div`
  ${props=>
	props.type ==="danger"&&
    css`
    	color:"red";
    	font-size:"20px";
    `
  }
  
  ${props=>
	props.type ==="info"&&
    css`
    	color:"gray";
    	font-size:"16px";
    `
  }
  
  ${props=>
	props.type ==="warnning"&&
    css`
    	color:"orange";
    	font-size:"16px";
    `
  }
`;


//声明props的默认值
MyDiv.defaultProps={
    type:"info",
}

export default MyDiv;
```

### 小结

styled-components是一种管理css的方式，了解即可，遇到不懂查文档。

## Supabase

p337-345

使用该库构建免费的后端服务

```sh
1.明确应用程序会用到哪些数据
2.创建关系表
3.使用Supabase API加载数据
```

### 什么是Supabase

```sh
1.Supabase是一项技术，允许开发人员快速构建后端，拥有完整的Postgres数据库
2.Supabase会自动创建一个数据库和一些匹配的api
3.Supabase还可以用于用户验证和文件存储
```

实际上就是一个后端微服务供应商，当你想快速搭建实验性后端时可以选择他们家的服务，免费服务是2个项目，必须保证周活跃，其他时间要付费。

### 使用Supabase

1.创建账户，登录`supabase.com`

推荐实际需要使用时查文档，因为与时俱进。

## React Query

```sh
英文文档
https://tanstack.com/query/v3

中文文档
https://cangsdarm.github.io/react-query-web-i18n/react/
```

作用

```sh
1.Remote state管理
2.用于处理数据的请求与存储
```

### 简介

React query本质是一个非常强大的管理远程状态的库，很多人也将他称为数据获取库。

```sh
1.允许我们编写更少的代码从API中获取数据，同时管理所有的数据
2.所有remote state都被缓存，这意味着数据将被存储方便在app上重用。举例，A组件向API发起请求拿到了城市数据，该数据将被缓存，当B组件要使用该城市数据时就不用再发请求，而是直接使用缓存中的城市数据，这样加快了响应速度。
3.React query会自动给出所有加载和错误状态
4.React Query在某些情况下会自动重新获取数据，例如某个超时之后，我们离开浏览器窗口再回来时，目的是为了保证remote state和app保持同步，例如用别的应用程序改变了远程状态，react query也将帮助同步。
5.预取数据，获取我们知道以后会变得重要的数据，经典例子就是分页。不仅可以为当前页获取数据，还可以为下一页获取。这样用户移动到下一页时就可以从缓存中读取。
6.很容易更新remote state
7.离线支持，由于数据已经缓存了，离线时仍然可以使用缓存的数据
```

remote state通常都是异步的，与UI state不同。

### React Query简单使用

**安装**

```sh
npm i @tanstack/react-query@4
```

**引入**

在App.jsx中引入，使用方式和ContextAPI类似

```jsx
import {QueryClient,QueryClientProvider} from "@tanstack/react-query"

//QueryClient()传入一个配置对象
const queryClient = new QueryClient({
    //默认配置项
    defaultOptions:{
        queries:{
            //缓存中的数据保存时间，单位毫秒
            staleTime:60*1000，
        }
    }
})

//用QueryClientProvider包裹所有组件，包括路由组件，client属性就是QueryClient实例
function App(){
    return (
    	<QueryClientProvider client={queryClient}>
          <GlobalStyles />
          <BrowserRouter>
            <Routes>
              <Route element={<AppLayout />}>
                <Route index element={<Navigate replace to="dashboard" />} />
                <Route path="dashboard" element={<Dashboard />} />
                <Route path="bookings" element={<Bookings />} />
                <Route path="cabins" element={<Cabins />} />
                <Route path="users" element={<Users />} />
                <Route path="settings" element={<Settings />} />
                <Route path="account" element={<Account />} />
              </Route>

              <Route path="login" element={<Login />} />
              <Route path="*" element={<PageNotFound />} />
            </Routes>
          </BrowserRouter>
        </QueryClientProvider>
    )
}
```

**安装react query开发工具**

只是一个npm包，不需要安装浏览器插件

```sh
npm i @tanstack/react-query-devtools
```

使用开发工具，作为QueryClientProvider的第一个子组件

```jsx
import {QueryClient,QueryClientProvider} from "@tanstack/react-query"
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
//QueryClient()传入一个配置对象
const queryClient = new QueryClient({
    //默认配置项
    defaultOptions:{
        queries:{
            //缓存中的数据保存时间，单位毫秒
            staleTime:60*1000，
        }
    }
})

//用QueryClientProvider包裹所有组件，包括路由组件，client属性就是QueryClient实例
function App(){
    return (
    	<QueryClientProvider client={queryClient}>
          {/*ReactQuery开发工具*/}
          <ReactQueryDevtools initialIsOpen={false} />
          <GlobalStyles />
          <BrowserRouter>
            <Routes>
              <Route element={<AppLayout />}>
                <Route index element={<Navigate replace to="dashboard" />} />
                <Route path="dashboard" element={<Dashboard />} />
                <Route path="bookings" element={<Bookings />} />
                <Route path="cabins" element={<Cabins />} />
                <Route path="users" element={<Users />} />
                <Route path="settings" element={<Settings />} />
                <Route path="account" element={<Account />} />
              </Route>

              <Route path="login" element={<Login />} />
              <Route path="*" element={<PageNotFound />} />
            </Routes>
          </BrowserRouter>
        </QueryClientProvider>
    )
}
```

引入后会在项目的左下角有个按钮，可以打开ReactQuery监测面板

### 获取数据

以前获取数据的方式

```jsx
import getCity from "./api/city"
function City(){
    useEffect(function(){
        getCity().then((data)=>{
            console.log(data)
        })
    })
    
    return <p>{{city}}</p>
}
```

使用React Query获取

```jsx
import getCity from "./api/city"
function City(){
    //返回值是一个查询对象
    const x = useQuery({
        //唯一标识要查询的数据，可能是一个复杂数组或者带有字符串的数组
        queryKey:["city"]
        //实际查询函数，负责查询，从API获取数据,就是一个异步函数
        queryFn:getCity
    })
    
    console.log(x)
    
    return <p>{{city}}</p>
}
```

我们经常使用的是从查询对象上解析出以下几个属性

```jsx
import getCity from "./api/city"
function City(){
  	/*
  		isLoading 是否在查询中，是个布尔值
  		data 查询到的数据，
  		error 错误信息
  	*/
    const {isLoading,data:cities,error} = useQuery({
        queryKey:["city"]
        queryFn:getCity
    })
    
    //如果正在加载就使用加载组件
    if(isLoading) return <Loading/>
    
    //如果不再加载中就可能是获取到数据了，后续再讨论报错问题
    return (
        <div>
            {cities.map(item=>{
                <p key={item.id}>{item.name}</p>
            })}
        </div>
        )
}
```

#### 缓存在数据中的作用

我们设置了一个过期时间

```js
const queryClient = new QueryClient({
    //默认配置项
    defaultOptions:{
        queries:{
            staleTime:60*1000，
        }
    }
})
```

他的作用就是每隔特定时间自动去请求接口更新数据，例如数据库的数据变化了，按找我们的设定，从缓存中读取，还将是陈旧的数据，1min后将更新为新数据，所以想要数据时时保持最新，可以使用

```js
const queryClient = new QueryClient({
    //默认配置项
    defaultOptions:{
        queries:{
            staleTime:0，
        }
    }
})
```

在以往的项目中我们可能需要使用轮询或者websocket来保证数据同步，而React Query只要设置这个过期时间即可。

### 操作异步数据

例如删除一个城市数据，传统做法

```jsx
import deleCity from "./api/city"
function City(){
    let id = 18
    const handleDelete = (id)=>{
        deleCity(id).then(res=>{
            console.log(res)
        })
    }
    return <button onClick={()=>handleDelete(id)}>删除</button>
}
```

使用React Query

```jsx
import deleCity from "./api/city"
function City(){
    let id = 18
   	/*
   		useMutation传入一个配置对象，获得一个操作实例
   	*/
    const x = useMutation({
        //异步操作，是一个箭头函数
        mutationFn:(id)=>deleCity(id)
    })
    return <button>删除</button>
}
```

我们一般从操作实例中解构出两个东西

```jsx
import deleCity from "./api/city"
function City(){
    let id = 18
   	/*
   		useMutation传入一个配置对象，获得一个操作实例
   		isLoading 是否操作完成，是个布尔值
   		mutate 封装好的操作函数，就是mutationFn
   	*/
    const {isLoading,mutate} = useMutation({
        //异步操作，是一个箭头函数
        mutationFn:(id)=>deleCity(id)
    })
    return <button onClick={()=>mutate(id)} disabled={isLoading}>删除</button>
}
```

#### 成功的回调

以上还没什么特殊，实际上我们可以设定操作成功的回调，例如

```jsx
import deleCity from "./api/city"
function City(){
    let id = 18
    const {isLoading,mutate} = useMutation({
        //异步操作，是一个箭头函数
        mutationFn:(id)=>deleCity(id)
        //操作成功之后执行的回调，例如删除成功后获取列表
        onSuccess:()=>{
        
    	}
    })
    return <button onClick={()=>mutate(id)} disabled={isLoading}>删除</button>
}
```

传统做法

```sh
删除成功后再次调用请求列表的接口，用请求到的数据更新列表。
```

React Query的做法

```sh
1.获取queryClient
2.删除成功后，让原列表数据"失效"，就是更新原列表数据，因为原列表数据被queryClient维护，使用唯一键queryKey
```

```jsx
import deleCity from "./api/city"
function City(){
    let id = 18
    
    //1.获取queryClient
    const queryClient = useQueryClient()
    
    const {isLoading,mutate} = useMutation({
        //异步操作，是一个箭头函数
        mutationFn:(id)=>deleCity(id)
        //操作成功之后执行的回调，例如删除成功后获取列表
        onSuccess:()=>{
        	//让原来维护的列表数据到期，借助queryKey
        	queryClient.invalidateQueries({
                queryKey:["city"]
            })
    	}
    })
    return <button onClick={()=>mutate(id)} disabled={isLoading}>删除</button>
}
```

#### 错误的回调

```jsx
import deleCity from "./api/city"
function City(){
    let id = 18
    
    const queryClient = useQueryClient()
    
    const {isLoading,mutate} = useMutation({
        //这里也可以直接写异步操作函数
        mutationFn:deleCity
        onSuccess:()=>{
        	queryClient.invalidateQueries({
                queryKey:["city"]
            })
    	},
        //错误的回调,这里的err是由mutationFn回调中的deleCity返回的
        onError:err=>{
            alert(err.message)
        }
    })
    return <button onClick={()=>mutate(id)} disabled={isLoading}>删除</button>
}
```

错误来源

```js
export async function deleteCity(id) {
  const { data, error } = await supabase.from("cabins").delete().eq("id", id);

  if (error) {
    console.error(error);
    throw new Error("Cabin could not be deleted");
  }

  return data;
}
```

### 抽离React Query

我们会把React Query内容抽离成useXx钩子，就是自定义hook，区别

```sh
hooks 下存放在几个features通用的hooks
React Query抽离的hook一般在各个功能点下，例如`features\city`下的`useDeleteCity.js`,起名风格一般是`use+请求方法/功能`
```



## React Hook Form

这是一个专门处理表单提交和错误的库。p352

**安装**

```sh
npm i react-hook-form@7
```

具体使用可以查询文档，

```sh
https://react-hook-form.com/

#中文文档
https://react-hook-form.nodejs.cn/
```

**对应课程部分**

```sh
p353 表单提交和reset
p354 表单验证
p355 图片上传
```









