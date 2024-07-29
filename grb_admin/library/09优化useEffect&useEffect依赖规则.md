





# 优化

## 解决useEffect无限请求的问题

来看以下代码

```jsx
function App(){
    //从路由传参中获取id
    const {id} = useParams()
    const [count,setCount] = useState(0)
    
    async function getCity(id){
        const res = await fetch(`http://192.40.10.14:8888/city/${id}`)
        const data = await res.json()
        setCount(count+1)
    }
    
    useEffect(function(){
        gitCity(id)
    },[id,gitCity])
}
```

问题

```sh
会陷入死循环，因为监听了getCity，解析:
1.getCity中修改了状态，导致App重新渲染
2.App重新渲染重新创建getCity，useEffect监听到getCity地址变化，执行回调
3.执行回调内部的getCity，又修改了状态回到了第一步
```

这就是死循环的由来，但如果我们不监听`getCity`,React内置的Eslint就会报错

所以我们可以用useCallBack来解决这个问题

```jsx
function App(){
    //从路由传参中获取id
    const {id} = useParams()
    const [count,setCount] = useState(0)
    
    //使用useCallBack就可以解决这个死循环问题
    //依赖项根据提示添加，看用到了那个变量
    const getCity = useCallBack(async function getCity(id){
        const res = await fetch(`http://192.40.10.14:8888/city/${id}`)
        const data = await res.json()
        setCount(count+1)
    },[])
    
    useEffect(function(){
        gitCity(id)
    },[id,gitCity])
}
```

## 通过代码分割来减少Bundle大小

课程p253

**知识储备**

```sh
每当用户访问应用程序的时候，基本上实在访问一个托管在某个服务器上的网站，一旦用户真的导航到应用程序，服务器将向客户端发回一个巨大的javascript文件，这就是请求，这个发回的文件就是bundle

【bundle】
本质上只是一个js文件，包含应用程序的整个代码，他被称之为束，因为像Vite货webPack这样的工具会将我们所有的开发文件捆绑成一个巨大的捆绑包，它同样包含我们所有的应用程序代码。
这意味着包一旦被下载，一次性加载整个React应用程序，本质上将它变成一个完全在客户机上运行的单页面应用程序，所以每当应用程序中的URL发生变化时，客户机只呈现一个新的React组件，且不用从服务器加载任何新文件，因为所有的js代码已经在bundle中了。

【bundle size】
为了使用应用程序，每个用户需要下载的js代码量。
bundle size越大，下载时间自然越长，所以这是一个优化的重点。

【code splitting】
代码拆分技术就是专门来解决这个问题的。
代码拆分基本上采取把bundle分成多个部分，所以不再是一个巨大的js文件，我们将有一堆小文件，随着时间的推移下载，这种顺序加载代码的过程称为惰性加载。
```

下面我们将探讨如何实现代码拆分以及惰性加载。我们可以将前端项目打包来验证这个巨大的js文件。

### 懒加载组件

最常见的做法是从路由级别进行调整，将路由组件进行懒加载，示例

原来的路由

```jsx
import {BrowserRouter,Routes,Route} from "react-router-dom"
import HomePage from "./HomePage"
import Product from "./Product"
import Price from "./Price"
import AppLayout from "./AppLayout"
import City from "./City"
import Village from "./Village"
import NotFound from "./NotFound"
function App(){
    return (
        <BrowserRouter>
            <Routes>
                <Route to="/" element={<HomePage/>}/>
                <Route to="/product" element={<Product/>}/>
                <Route to="/price" element={<Price/>}/>
                <Route to="/app" element={<AppLayout/>}>
                    <Route to="/city" element={<City/>}/>
                    <Route to="/village" element={<Village/>}/>
                </Route>
                <Route to="*" element={<NotFound/>}></Route>
            </Routes>
        </BrowserRouter>
    	
    )
}
```

懒加载组件的语法

```jsx
//从react引入lazy方法
import {lazy} from "react"

//lazy方法接收一个回调函数
const 组件名 = lazy(()=>import(组件路径))
```

例如我们改造后

```jsx
import {lazy} from "react"
import {BrowserRouter,Routes,Route} from "react-router-dom"

const HomePage = lazy(()=>import("./HomePage"))
const Product = lazy(()=>import("./Product"))
const Price = lazy(()=>import("./Price"))

import AppLayout from "./AppLayout"
import City from "./City"
import Village from "./Village"
import NotFound from "./NotFound"
function App(){
    return (
        <BrowserRouter>
            <Routes>
                <Route to="/" element={<HomePage/>}/>
                <Route to="/product" element={<Product/>}/>
                <Route to="/price" element={<Price/>}/>
                <Route to="/app" element={<AppLayout/>}>
                    <Route to="/city" element={<City/>}/>
                    <Route to="/village" element={<Village/>}/>
                </Route>
                <Route to="*" element={<NotFound/>}></Route>
            </Routes>
        </BrowserRouter>
    	
    )
}
```

**Suspense组件**

这里还有一个问题，懒加载会将组件挂起，在组件还没抵达的时候我们的页面可能一片空白甚至报错，所以我们需要留一条退路，这个退路就是Suspense组件，它的基本功能就是当路由还没跳转成功时指定一个呈现组件来填补空白，未来我们会详解这个组件，现在只要知道，我们一般用它包裹整个Routes，然后指定一个全局加载组件，例如

```jsx
import {lazy,Suspense} from "react"
import {BrowserRouter,Routes,Route} from "react-router-dom"

const HomePage = lazy(()=>import("./HomePage"))
const Product = lazy(()=>import("./Product"))
const Price = lazy(()=>import("./Price"))

import AppLayout from "./AppLayout"
import City from "./City"
import Village from "./Village"
import NotFound from "./NotFound"

//全局加载组件
import SpinnerFullPage from "./SpinnerFullPage"
function App(){
    return (
        <BrowserRouter>
            <Suspense fallback={<SpinnerFullPage/>}>
                <Routes>
                    <Route to="/" element={<HomePage/>}/>
                    <Route to="/product" element={<Product/>}/>
                    <Route to="/price" element={<Price/>}/>
                    <Route to="/app" element={<AppLayout/>}>
                        <Route to="/city" element={<City/>}/>
                        <Route to="/village" element={<Village/>}/>
                    </Route>
                    <Route to="*" element={<NotFound/>}></Route>
                </Routes>
            </Suspense>
            
        </BrowserRouter>
    	
    )
}
```

## 优化课程总结

p254

不该做的事情

```sh
1.不要过早的优化
2.如果没必要优化那就尽量不要优化任何东西
3.不要把所有组件都用memo包裹，useMemo和useCallBack同理
4.如果ContextAPI不是很慢或者消费者不是特别多，就不要优化他
(因为记忆化会对性能产生影响)
```

该做的事情

```sh
1.利用React开发工具提供的Profile探查器找到性能瓶颈或者更好的识别问题。当我们看到一个滞后或者缓慢的UI时，在进行处理。
2.用记忆化处理那些昂贵的re-render，例如大数据量的单次计算。
```

# useEffect规则

## useEffect依赖数组规则

```sh
1.每一个在effect中使用的state、prop以及Context value都必须在依赖数组中
2.所有的`reactive value（反应值）`必须包含在依赖数组中，
任何function或者变量，只要他们与state、prop、contextVlaue产生了关系，他们就是反应值。
只要反应值出现在了effect中，也必须包含在依赖中。
3."不应该将对象或者数组用作依赖项"，因为在实际运行中，每次重新渲染都会把数组和对象重置为新的引用，即使对象和数组的内容保持不变，这一切会导致effect再次执行，也可能会陷入死循环
```

**理解反应值**

```jsx
const [number,setNumber] = useState(5)
const [duration,setDuration] = useState(0)
//反应值
const mins = Math.floor(duration)
//反应值
const secs = (duration - mins)*60
//反应值
const formatDur = function(){
    return `${mins}:${secs<10?"0":""}${secs}`
}

useEffect(
	function(){
        document.title = `${number}workout${formatDur()}`
    },
    //依赖数组也要包括用到的反应值
    [number,formatDur]
)
```

### **移除不需要的依赖项**

按照规则我们必须在依赖列表中包含所有的`reactive value（反应值）`,但在某些情况下，这会导致effect运行的太频繁，引入新问题

>不要省略掉依赖，那是不正确的解决方案

#### **移除函数依赖**

**解决方案1**

将函数移入effect，如果函数在effect中，那就不再是effect的依赖了。例如

```jsx
function App(){
    //从路由传参中获取id
    const {id} = useParams()
    const [count,setCount] = useState(0)
    
    async function getCity(id){
        const res = await fetch(`http://192.40.10.14:8888/city/${id}`)
        const data = await res.json()
        setCount(count+1)
    }
    
    useEffect(function(){
        gitCity(id)
    },[id,gitCity])
}
```

优化方案

```jsx
function App(){
    //从路由传参中获取id
    const {id} = useParams()
    const [count,setCount] = useState(0)
    
    useEffect(function(){
        //将函数直接移入effect，那么就不再是依赖了
        async function getCity(id){
            const res = await fetch(`http://192.40.10.14:8888/city/${id}`)
            const data = await res.json()
            setCount(count+1)
        }
        gitCity(id)
    },[id])
}
```

这种方案是最简方案，但是会存在另一个问题，如果这个函数在其他需要被使用就无法采用这个方案

**解决方案2**

使用useCallBack记忆函数，当有很多地方需要这个函数时就应该采用这种方案，例如

```jsx
function App(){
    //从路由传参中获取id
    const {id} = useParams()
    const [count,setCount] = useState(0)
    
    async function getCity(id){
        const res = await fetch(`http://192.40.10.14:8888/city/${id}`)
        const data = await res.json()
        setCount(count+1)
    }
    
    useEffect(function(){
        gitCity(id)
    },[id,gitCity])
}
```

优化解决

```jsx
function App(){
    //从路由传参中获取id
    const {id} = useParams()
    const [count,setCount] = useState(0)
    
	//将函数记忆
    const getCity = useCallBack(async function getCity(id){
        const res = await fetch(`http://192.40.10.14:8888/city/${id}`)
        const data = await res.json()
        setCount(count+1)
    },[])
    
    useEffect(function(){
        gitCity(id)
    },[id,gitCity])
}
```

**解决方案3**

如果函数本身不是反应值，可以将该函数移出组件，因为移出组件他就不会再受重新渲染的影响，例如

```jsx
function App(){
    //从路由传参中获取id
    const {id} = useParams()
    const [count,setCount] = useState(0)
    
    //这个函数不是反应值，完全可以被在组件内
    async function getCity(id){
        const res = await fetch(`http://192.40.10.14:8888/city/${id}`)
        const data = await res.json()
        return data
    }
    
    useEffect(function(){
        gitCity(id)
    },[id,gitCity])
}
```

解决方案

```jsx
//这个函数不是反应值，完全可以被在组件内
async function getCity(id){
    const res = await fetch(`http://192.40.10.14:8888/city/${id}`)
    const data = await res.json()
    return data
}

function App(){
    //从路由传参中获取id
    const {id} = useParams()
    const [count,setCount] = useState(0)
   
    useEffect(function(){
        gitCity(id)
    },[id,gitCity])
}
```

其实这种情况它就不属于真正的依赖

#### 移除对象依赖

**解决方案1**

不包括整个对象，而是只包括我们要用到的那个对象属性,只要这些属性是简单类型，那么就会很好的生效

```jsx
function App(){
    const [count,setCount] = useState(0)
   	const option ={
        count:count,
        name:"lisa"
    }
    useEffect(function(){
        option.count = count
    },[count,option.count])
}
```

**解决方案2**

使用useMemo包裹我们想要作为依赖的对象。

#### 终极方案

使用useReducer来管理状态

## 什么时候不要使用useEffect

```sh
1.不要用useEffect响应用户事件，用户事件应该尽可能使用事件程序函数来处理，而不是useEffect，因为事件处理函数也是能处理副作用的
2.在组件挂载时请求数据，在小型项目中我们的确经常在挂载时请求数据，但是在中大型项目中应该使用更专业的数据获取库，例如React Query
3.同步两个状态之间的变化，这意味着useEffect来监听状态的更改来同步另一个state，这样做的问题是会触发多次render，实际项目中如果两个状态存在同步关系，应该把其中一个状态变为另一个的派生状态
```

当然以上都只是建议，因为实际上useEffect做这些理论上没问题，只是说不要滥用useEffect。