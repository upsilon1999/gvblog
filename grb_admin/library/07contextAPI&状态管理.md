## 学习思路

```sh
1.跟随课程步骤记录知识点
2.结课后花两天整理笔记
3.还剩90课，每天30课
```

# 单元重点:上下文API

```sh
1.Context API
2.深入了解state管理，了解一些新类型的状态和处理它们的工具
```

## 什么是context API

* 上下文API解决的问题

  我们需要将一些状态传递到多个深度嵌套的子组件中，我们以前的做法是找到公共祖先组件，然后将一个状态传递下去。这样容易产生一个新问题，因为通过树的多个层次传递props，会变得繁琐和不方便，我们将之称为`prop drilling(螺旋钻孔)`问题。遇到这种情况我们之前的解决方案是重组组件，例如利用标签体,但是这种方法不是每次都可行。

  所以我们需要一个API能实现，直接将祖先组件的状态传递给后代组件，这就是上下文API，上下文API基本上允许树种各处的组件读取状态，一个上下文共享。

* 详述context API

  首先，上下文API是一个在整个应用程序中传递数据的系统,无需手动将props传递到组件树，他本质上允许我们广播全球状态（global state）,因此，声明现在应该对某个上下文的所有子组件可用

  ```sh
  【Provider】
  提供者，是一个特殊的组件，允许所有子组件访问指定的值。
  提供者可以出现在组件树的任何地方，但我们通常把他放在最上面。
  
  【value】
  我们想要提供的数据，通常是state和function
  
  【Consumers】
  消费者，就是读取提供者提供的数据的所有组件，也称为订阅上下文的组件，因为他们能够从上下文中读取值。
  ```

  一个Provider可以拥有任意多个Consumers

* Provider的value更新时会发生什么

  每当value更新时，所有Consumers都会`re-render`,即所有读取value的组件当共享的值以某种方式更新时，Provider将立即通知Consumers,然后Consumers组件更新。

## 使用上下文API

1.创建上下文API

```jsx
//1.引入创建上下文API的钩子
import {createContext} from "react"

//2.创建context API组件
const PostContext = createContext()

function App(){
    return <p>...</p>
}

export default App
```

**注意事项:**createContext方法创建的实际上是一个React组件，所以变量名用大写

2.用Provider包裹要使用上下文的部分

```jsx
import {createContext} from "react"

const PostContext = createContext()

function App(){
    return (
        <div>
            {/* 
            	被Context API的Provider包裹的部分就是能接收到该上下文的区域
            	反之不能接收到上下文
            	
            	不过在实际开发中，我们用Provider把所有部分都包裹进去
            */}
            <PostContext.Provider>
                {/* 处于Provider中，组件本身极其后代都能接收到Context消息 */}
                <p>...</p>
                <Header></Header>
                <Logo/>
                <Main/>
        	</PostContext.Provider>
            {/* 处于Provider外，不能接收到消息*/}
            <Footer/>
        </div>
    )
}
export default App
```

3.传递value

```jsx
import {createContext,useState} from "react"

const PostContext = createContext()

function App(){
    const [count,setCount] = useState(0)
    const Add = ()=>{
        setCount(count+1)
    }
    return (
        <div>
            {/* 
            	value和其他组件属性一样，能传递任何东西，
            	不过我们一般用来传递state或者function
            	
            	下面的例子中我们传递了一个对象，对象中有值、状态、方法
            */}
            <PostContext.Provider value={{
                    id:8456,
                    name:"lisi",
                    count:count,
                    addMethods:Add
                }}>
                <p>...</p>
                <Header></Header>
                <Logo/>
                <Main/>
        	</PostContext.Provider>
        </div>
    )
}

export default App
```

4.消费者读取数据

我们假定City组件是Main的子组件，由于Main组件在上下文中，所以他的后代组件也能读取上下文，所以我们用City来示例

**改造一下App组件**

这次的改造主要是抛出上下文组件，在后续我们会优化，因为上下文组件也算一个React组件，而实际开发中尽量不要在一个文件中创造多个React组件

```jsx
import {createContext,useState} from "react"

//抛出Context组件，因为消费者需要使用这个组件
export const PostContext = createContext()

function App(){
    const [count,setCount] = useState(0)
    const Add = ()=>{
        setCount(count+1)
    }
    return (
        <div>
            {/* 
            	value和其他组件属性一样，能传递任何东西，
            	不过我们一般用来传递state或者function
            	
            	下面的例子中我们传递了一个对象，对象中有值、状态、方法
            */}
            <PostContext.Provider value={{
                    id:8456,
                    name:"lisi",
                    count:count,
                    addMethods:Add
                }}>
                <p>...</p>
                <Header></Header>
                <Logo/>
                <Main/>
        	</PostContext.Provider>
        </div>
    )
}

export default App
```

**City组件消费者**

```jsx
//1.引入useContext钩子，顾名思义，使用上下文
import {useContext} from "react"
//1.5 引入Context组件
import {PostContext} from "./App.jsx"
function City(){
    /*
    	useContext(上下文组件)
    	返回值就是Provider提供的value
    */
    const x = useContext(PostContext)
    
    return (
        <div>
            {/*更推荐的方式是先解构*/}
            <p>城市id是:{x.id}</p>
            <p>城主名称是:{x.name}</p>
            <p>数量是:{x.count}</p>
            <button onClick={x.addMethods}>点我增加</button>
        </div>
    ) 
}
```

## 优化ContextAPI的使用

我们希望实现一个自定义提供程序组件以及自定义钩子来消费数据，目的就是创建一个外置上下文组件，可以从该组件中提取Provider和使用封装后的消费者，即抽离ContextAPI逻辑，让他更独立。

### 新建一个PostContext组件

#### 命名规范

规范不是强制的，但希望遵守

```sh
1.文件名称就是Context组件的名称，例如我们的Context组件是PostContext
2.文件内部的组件名使用XxProvider，因为该组件就是抽离的Provider成分
3.最终需要抛出的部分XxProvider，useXx，一个是Provider，一个是消费者方法
```

#### 组件改造1--构筑Provider

1.观察原来的App组件

```jsx
import {createContext,useState} from "react"
export const PostContext = createContext()

function App(){
    //与Provider有关，需要抽离到Provider或者传递到Provider
    const [count,setCount] = useState(0)
    const Add = ()=>{
        setCount(count+1)
    }
    
    //与Provider无关保留
    const [age,setAge] = useAge(20)
    
    return (
        <div>
            <PostContext.Provider value={{
                    id:8456,
                    name:"lisi",
                    count:count,
                    addMethods:Add
                }}>
                {/* Provider的标签体内容，可以使用children属性传递*/}
                <p>...</p>
                <Header></Header>
                <Logo/>
                <Main/>
        	</PostContext.Provider>
        </div>
    )
}

export default App
```

2.将原先和Provider相关的都迁移过来

```jsx
import {createContext} from "react"
//1.创建context组件
const PostContext = createContext()
function PostProvider({children}){
    const [count,setCount] = useState(0)
    const Add = ()=>{
        setCount(count+1)
    }
    
    return (
        <div>
            <PostContext.Provider value={{
                    id:8456,
                    name:"lisi",
                    count:count,
                    addMethods:Add
                }}>
                {children}
        	</PostContext.Provider>
        </div>
    )
}

export {PostProvider}
```

改造后的App组件

```jsx
import {createContext,useState} from "react"
import {PostProvider} from "./PostContext"

function App(){
    //与Provider无关保留
    const [age,setAge] = useAge(20)
    
    return (
        <div>
            <PostProvider>
                <p>...</p>
                <Header></Header>
                <Logo/>
                <Main/>
        	</PostProvider>
        </div>
    )
}

export default App
```

3.传递版

有时候有些状态不仅上下文要用，其他组件可能也要用(如果用上下文组件包含了所有子组件没有这个问题，但有时候有些项目他们会有如下形式)，例如

```jsx
import {createContext,useState} from "react"
export const PostContext = createContext()

function App(){
    //与Provider有关，同时Footer组件也用到了且Footer不在上下文中
    const [count,setCount] = useState(0)
    const Add = ()=>{
        setCount(count+1)
    }
    
    //与Provider无关保留
    const [age,setAge] = useAge(20)
    
    return (
        <div>
            <PostContext.Provider value={{
                    id:8456,
                    name:"lisi",
                    count:count,
                    addMethods:Add
                }}>
                {/* Provider的标签体内容，可以使用children属性传递*/}
                <p>...</p>
                <Header></Header>
                <Logo/>
                <Main/>
        	</PostContext.Provider>
            {/*不在上下文中*/}
            <Footer fcount={count}/>
        </div>
    )
}

export default App
```

所以采用传递的形式

```jsx
import {createContext} from "react"
//1.创建context组件
const PostContext = createContext()
//接收公共状态
function PostProvider({children,count,Add}){
    return (
        <div>
            <PostContext.Provider value={{
                    id:8456,
                    name:"lisi",
                    count:count,
                    addMethods:Add
                }}>
                {children}
        	</PostContext.Provider>
        </div>
    )
}

export {PostProvider}
```

改造后的App组件

```jsx
import {createContext,useState} from "react"
import {PostProvider} from "./PostContext"

function App(){
    //与Provider有关，同时Footer组件也用到了且Footer不在上下文中
    const [count,setCount] = useState(0)
    const Add = ()=>{
        setCount(count+1)
    }
    
    //与Provider无关保留
    const [age,setAge] = useAge(20)
    
    return (
        <div>
            <PostProvider count={count} Add={Add}>
                {/* Provider的标签体内容，可以使用children属性传递*/}
                <p>...</p>
                <Header></Header>
                <Logo/>
                <Main/>
        	</PostProvider>
            {/*不在上下文中*/}
            <Footer fcount={count}/>
        </div>
    )
}

export default App
```

#### 组件改造2--使用者钩子

现在的开发规范是将和上下文有关的组件或者自定义钩子都放在一个文件中。

```jsx
import {createContext} from "react"
//1.创建context组件
const PostContext = createContext()
function PostProvider({children}){
    const [count,setCount] = useState(0)
    const Add = ()=>{
        setCount(count+1)
    }
    
    return (
        <div>
            <PostContext.Provider value={{
                    id:8456,
                    name:"lisi",
                    count:count,
                    addMethods:Add
                }}>
                {children}
        	</PostContext.Provider>
        </div>
    )
}

//2.封装一个使用上下文的钩子，我们可以称为消费者钩子
//目的，不用在每次传递时再引入Context组件，在有多个Context组件时可读性更强
function usePosts(){
    const context = useContext(PostContext)
    return context
}

export {PostProvider,usePosts}
```

例如City组件，原来不使用消费者钩子,

```jsx
//1.引入useContext钩子，顾名思义，使用上下文
import {useContext} from "react"
//1.5 引入Context组件
import {PostContext} from "./App.jsx"
function City(){
    const x = useContext(PostContext)
    return (
        <div>
            {/*更推荐的方式是先解构*/}
            <p>城市id是:{x.id}</p>
            <p>城主名称是:{x.name}</p>
            <p>数量是:{x.count}</p>
            <button onClick={x.addMethods}>点我增加</button>
        </div>
    ) 
}
```

有如下缺点

```sh
1.每次都要引入useContext钩子
2.要引入对应Context组件并针对性使用，当有多个Context组件时，可读性差
```

改造后

```jsx
//1.引入消费者钩子
import {usePosts} from "./PostContext.jsx"
function City(){
    const x = usePosts()
    return (
        <div>
            {/*更推荐的方式是先解构*/}
            <p>城市id是:{x.id}</p>
            <p>城主名称是:{x.name}</p>
            <p>数量是:{x.count}</p>
            <button onClick={x.addMethods}>点我增加</button>
        </div>
    ) 
}
```

##### 防止滥用消费者钩子

如果在上下文外面使用消费者钩子将会是undefined，例如

```jsx
import {createContext,useState} from "react"
import {PostProvider,usePosts} from "./PostContext"

function App(){
    //在这里使用消费者钩子会是undefined
    const x = usePosts()
    console.log(x)//undefined
    
    
    return (
        <div>
            <PostProvider>
                <p>...</p>
                <Header></Header>
                <Logo/>
                <Main/>
        	</PostProvider>
        </div>
    )
}

export default App
```

原因,`useContext`只在消费者组件中获取到Provider提供的值，所以我们可以优化一下钩子,在滥用时抛出错误

```jsx
function usePosts() {
  const context = useContext(PostContext);
  //在context为undefined时抛出错误
  if (context === undefined){
      throw new Error("PostContext was used outside of the PostProvider");
  }
   
  return context;
}
```

# 单元重点:state的管理

## state的类型

1.根据state的可访问性

```sh
【本地状态(local state)】
1.被一个或者更少的组件使用
2.只能在定义它的组件或者其子组件中访问

【全球状态(global state)】
1.可能被非常多的组件使用
2.app内所有的组件都可访问
```

2.根据state的使用区域(domain)

```sh
【远程状态(remote state)】
1.远程状态基本上是从远程服务器加载所有的应用程序数据，通常使用API，它基本上是存在于服务器上的状态，但是可以加载到我们的应用程序中
2.通常异步获取，
3.数据可能需要经常重新获取和更新，因此在大规模应用中，数据缓存更新和重新验证等等，我们都需要一些专门的工具

【UI状态(UI state)】
1.除远程状态外的其他一切，例如当前选择的主题、渲染列表、表单数据等等，即不是从api获取的核心应用程序数据
2.通常是同步获取，并存储在应用程序中，并且根本不与服务器进行交互，这意味着UI状态可以用已有的工具来直接处理，例如useState和useReducer等等

"注意"
由于远程状态与UI状态不同，我们通常要异步获取远程状态
```

## state的放置

我们应该在项目的什么位置放置state，每当我们有一个新state，大约有六种不同的选择可以把他放在哪里，例如

```sh
【情形1:local component】
存放地点:本地组件或者译为当前组件
采用工具:useState、useReducer或者 useRef
何时使用:局部状态(local state)，用于该组件及其后代组件

【情形2:Parent component】
存放地点:父组件或者说公共父组件
采用工具:useState、useReducer或者 useRef
何时使用:状态提升，多用于兄弟通信或者同祖先的后代通信

【情形3:Context】
存放地点:上下文
采用工具:Context API+useState或useReducer
何时使用:全球状态,我们通常用它来包裹根组件实现全球状态，当然也可以包裹某个祖先让他的所有后代通信
注意事项:Context API适合管理UI状态，不一定适合远程状态，尤其是在构建更大的应用程序时

【情形4：状态管理库】
存放地点:第三方状态库
采用工具:Redux、React Query、SWR、Zustand等等
何时使用:全球状态，远程和UI都可以

【情形5:URL】
存放地点:URL中
采用工具:React Router
何时使用:全局状态，用于页面之间通信

【情形6:Browser】
存放地点:浏览器
采用工具:Local storage,session storage等等
何时使用:在用户浏览器存放数据
```

## state的管理工具选择

因为按照可访问性分为了`local state`和`global state`,按照使用区域可分为`UI state`和`Remote state`,所以我们做出一张交叉表

```sh
【local state&UI state】
工具为:
1.useState
2.useReducer
3.useRef

【local state&remote state】
1.fetch+useEffect+useState/useReducer

【global state&UI state】
1.Context API+useState/useReducer
2.Redux,Zustand,Recoil等等
3.React Router

【global state&remote state】
1.Context API+useState/useReducer
2.Redux,Zustand,Recoil等等
(以下三种内置了缓存和自动重取机制，应对了异步问题)
3.React Query
4.SWR
5.RTK Query
```





