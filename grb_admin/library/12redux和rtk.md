# 单元重点

```sh
1.基于useReducer的Redux
2.现代Redux Toolkit
3.用Thunks请求API
```

## Redux介绍

```sh
1.Redux是一个第三方库，用来管理全局状态(global state)
2.Redux是个独立的库，可以与任何框架一起使用，甚至是普通的js项目都可以使用，React一般会结合react-redux使用
3.Redux背后的主要思想是可以存储应用程序中的所有全局状态，在这个可以全局访问的store中，我们可以用actions来对状态进行更新
4.Redux其实很像Context API+useReducer

5.现在有两个版本的Redux:
a)传统Redux b)现代Redux Toolkit
```

**特点**

```sh
global store更新后，所有消费者组件都会re-render
```

### 是否应该学习Redux

现在的高级状态管理库很多，Redux已经不再是唯一选择。我们学习Redux的原因

```sh
1.Redux很难学，难学所以要学:)
2.在旧的项目中可能有很多时候用到了Redux
3.有些应用程序需要类似的Redux的东西
```

### Redux的工作原理

我们需要先回顾useReducer的工作原理

```sh
1.组件中的事件将一个action交给dispatch，这个action通常包含一个type类型和payload有效载荷。
2.dispatch把action派发给reducer函数，该函数根据action对当前状态进行加工，然后得到下一个状态。
3.状态的变化引起re-render
```

redux的工作原理，是在useReducer上做的改造，有两个区别

```sh
【区别1】
1.组件事件将一个action交给dispatch，dispatch将把这个action交给store
2.store是存放有一堆reducer函数的仓库，store会识别派发来的action，让正确的reducer函数来处理它

【区别2】
1.action不再是我们手动写出，而是使用一个Action creator function
2.注意：这是一个可选的特性，不是Redux必须的
```

所以redux的流程

```sh
1.组件中的事件发生，我们调用Action creator function生成一个action
2.然后将action交给dispatch，再被派发到仓库
3.store选择合适的reducer函数来处理action
4.reducer函数利用action和当前状态计算出新状态
5.状态更新，所以的订阅组件(消费者组件)re-render
```

**action的解读**

action实际上可以是任何形式，只是开发规范让我们采取以下形式

```js
{type:"分发事件名",payload:要传递的值}
```

例如

```js
{type:"add",payload:100}
{type:"person",payload:{
	name:"lisi",
    age:18
}}
```

## 传统Redux使用

### 了解redux

>传统的redux已经不再被推荐使用，所以有很多方法都被显示废弃了，但是我们还是得了解他

1.安装传统的Redux

```sh
npm i redux
```

2.书写一个store.js，我们再里面写redux内容

```js
//1.引入创建仓库的方法
import {createStore} from "redux"

const initValue = {
    val:0,
    sum:0
}


//这一部分和useReducer做的差不多，就是设置初始状态和reducer函数
function reducer(state=initValue,action){
    switch(action.type):
        case "count/add":
        	return {...state,val:state.val+action.payload}
    	case "count/dec":
    		return {...state,val:state.val-action.payload}
    	default
    		return initValue
}

//将reducer函数注册为一个仓库
const store = createStore(reducer)


//使用仓库实例上的dispatch方法
store.dispatch({type:"count/add",payload:5})

//获取最新的状态值
store.getState()
```

上述是一个脱离实际的案例，他向我们展示了redux的实质，开发中只要把store暴露出去，然后外部引入后调用dispatch即可，和useReducer使用类似。

```sh
Redux与useReducer的区别是，useReducer维护的状态我们需要层层传递到使用它的地方，而Redux维护的状态则是可以从全局取用。
```

**知识点小结**

```js
【仓库是什么】
实际上就是根据reducer函数创建一个仓库。
const store = createStore(reducer)

【派发】
在需要使用的地方引入仓库，例如
import store from "./xxx"
然后进行使用
store.dispatch(action)

【读取和使用状态】
仓库.getState()
```

**解释一个知识点**

```js
"count/add"
"count/dec"
```

实际上就是普通字符串，事件别名，写成以下形式也完全没问题

```js
"add" "dec"
```

但是我们知道一点，我们会在store会处理多个reducer，所以为了让我们传递action时知道自己在干什么，我们需要将事件区分清楚，于是就有了一些开发规范，用第一个单词来划分模块，后面依次写功能

```js
//字符串型,
"count/add"

//常量型
COUNT_ADD
```

#### 使用Actions Creators

动作创建者只不过是返回action的简单函数，即使没有它，Redux也可以很好的工作，例如

```js
//就是写一个函数来返回action
function add(amount){
    return {type:"count/add",payload:amount}
}

//然后在传递action时就调用这个函数
store.dispatch(add(5))
```

实际的用处就是不用每次使用都写action，而是用封装的函数来返回，提升代码可读性，规范代码

```sh
有时手写action容易出错，例如
{type:"count/add",payload:amount}
{type:"count/aDd",payload:amount}
```

封装成函数就可以避免这个问题

**一些可能出现的写法**

1.用一个常量来保存事件字符串

```sh
const COUNT_ADD = "count/add"
```

然后把所有用到该字符串的地方都用这个常量替换，是一种代码风格

#### 多个reducer

我们在之前讲解原理的时候说过

```sh
store是存放有一堆reducer函数的仓库，store会识别派发来的action，让正确的reducer函数来处理它
```

所以现在我们来构造一个有多个reducer的仓库，实际上就是构造一个rootReducer把所有的reducer包裹进去，然后用rootReducer去实例化仓库，例如

```js
import {createStore,combineReducers} from "redux"

const initCount = {
    val:0,
    sum:0
}

function countReducer(state=initCount,action){
    switch(action.type):
        case "count/add":
        	return {...state,val:state.val+action.payload}
    	case "count/dec":
    		return {...state,val:state.val-action.payload}
    	default
    		return initCount
}


const initPerson={
    name:"lisi",
    age:18
}
//action的payload的形式不一定是一样的,因为就是个载体，最终根据事件来派发处理
function personReducer(state=initPerson,action){
    switch(action.type):
        case "person/old":
        	//例如此次action为{type:"person/old",payload:20}
        	return {...state,age:state.age+action.payload}
    	case "person/changeName":
    		//action为{type:"person/old",payload:{name:"sam"}}
    		return {...state,name:action.payload.name}
    	default
    		return initPerson
}

//将多个reducer函数用rootReducer包裹
/*
	combineReducers是redux的一个方法，用于整合多个reducer
	接收一个对象，
	对象键为要给reducer函数的状态别名
	对象的值就是reducer函数
*/
const rootReducer = combineReducers({
    count:countReducer,
    person:personReducer
})

//用rootReducer来注册仓库实例
const store = createStore(rootReducer)

//使用仓库实例上的dispatch方法
//我们不用指定reducer函数，这就是我们说的store会自动识别这是哪个reducer函数的方法然后去执行
store.dispatch({type:"count/add",payload:5})

//获取最新的状态值
/*
	当我们有多个reducer时，这就是刚才combineReducers的内容,一个大对象，例如
	{
		count:{ val:5,sum:0 },
		person:{ name:"lisi",age:18 }
	}
	
*/
store.getState()
```

当reducer非常多时，全写在一个文件中太混乱了，所以我们要进行拆分

#### 拆分store文件

在大型项目中我们应该有明确的代码组织结构，所以在src目录下建立文件夹

```sh
【features】
功能文件夹，根据每个功能在其下建立分门别类的文件夹，然后把对应内容转移进去
```

例如我们目前的store.js文件为

```js
import {createStore,combineReducers} from "redux"

const initCount = {
    val:0,
    sum:0
}

function countReducer(state=initCount,action){
    switch(action.type):
        case "count/add":
        	return {...state,val:state.val+action.payload}
    	case "count/dec":
    		return {...state,val:state.val-action.payload}
    	default
    		return initCount
}


const initPerson={
    name:"lisi",
    age:18
}
//action的payload的形式不一定是一样的,因为就是个载体，最终根据事件来派发处理
function personReducer(state=initPerson,action){
    switch(action.type):
        case "person/old":
        	//例如此次action为{type:"person/old",payload:20}
        	return {...state,age:state.age+action.payload}
    	case "person/changeName":
    		//action为{type:"person/old",payload:{name:"sam"}}
    		return {...state,name:action.payload.name}
    	default
    		return initPerson
}

//将多个reducer函数用rootReducer包裹
const rootReducer = combineReducers({
    count:countReducer,
    person:personReducer
})

//用rootReducer来注册仓库实例
const store = createStore(rootReducer)
```

我们在features目录下分别建立counts和persons文件夹，命名规范就是Xxreducer中的Xx加s，实际上我们应该把项目中操作count(计数)、person(人)的所有文件，只要提供功能的都转入这里。

```sh
"文件目录是为了规范项目，按照项目组的开发风格决定，这里只是告诉大家一种开发风格"
例如现在features目录下有counts文件夹，我们在counts下面再新建`countSlice.js`，xxSlice文件就是存储对应的reducer
```

例如，`features/counts/countSlice.js`

```js
const initCount = {
    val:0,
    sum:0
}

function countReducer(state=initCount,action){
    switch(action.type):
        case "count/add":
        	return {...state,val:state.val+action.payload}
    	case "count/dec":
    		return {...state,val:state.val-action.payload}
    	default
    		return initCount
}

//创造action的函数
//将action函数分别抛出
export function Add(num){
    return {type:"count/add",payload:num}
}

export function Dec(num){
    return {type:"count/dec",payload:num}
}

//默认抛出reducer
export default countReducer
```

例如,`features/person/personSlice.js`

```js
const initPerson={
    name:"lisi",
    age:18
}

export default function personReducer(state=initPerson,action){
    switch(action.type):
        case "person/old":
        	return {...state,age:state.age+action.payload}
    	case "person/changeName":
    		return {...state,name:action.payload.name}
    	default
    		return initPerson
}
```

然后再src下的store.js文件下进行引入

```js
import {createStore,combineReducers} from "redux"
//引入对应的reducer
import countReducer from "./features/counts/countSlice.js"
import personReducer from "./features/counts/personSlice.js"

const rootReducer = combineReducers({
    count:countReducer,
    person:personReducer
})

//用rootReducer来注册仓库实例
const store = createStore(rootReducer)

//导出仓库
export default store
```

我们哪个组件需要使用到store就把store导入

### 了解react-redux

到目前为止我们的redux和React程序是分离的，还是简单的引用文件并使用，

此时store内维护的状态对react还不是全局状态，更起不到全局更新React组件的效果

所以现在让我们将store注入到React中

在入口文件，cra中是`index.js`,vite项目中是`main.js`,未注入store前

```jsx
import React from "react"
import ReactDOM from "react-dom/client"
import App from "./App"

const root = ReactDOM.createRoot(document.getElementById("root"))
root.render(
	<React.StrictMode>
        <App/>
    </React.StrictMode>
)
```

要将redux注入React，我们需要在安装一个包，

```sh
npm i react-redux
```

将store注入React

```jsx
import React from "react"
import ReactDOM from "react-dom/client"

//从react-redux引入提供者
import {Provider} from "react-redux" 

//引入仓库
import myStore from "./store"
import App from "./App"

const root = ReactDOM.createRoot(document.getElementById("root"))
root.render(
	<React.StrictMode>
        {/*
        	类似上下文组件一样将myStore传入
        	这里的store属性是react-redux对ContextAPI的封装
        */}
        <Provider store={store}>
        	<App/>
        </Provider>
    </React.StrictMode>
)
```

**理解**

```sh
通过注入store我们可以发现，这和ContextAPI的原理基本一致，用Provider将App根组件包裹，这样全局状态就发布到了每个想要读取他的组件
```

#### **读取状态**

现在React和Redux两者结合之后，我们就可以使用`react-redux`提供的钩子来操作了。现在让我们来的City组件来使用Redux

```jsx
//引入读取状态的钩子
import {useSelector} from "react-redux"

function City(){
    /*
    	语法为
    	useSelector(当前仓库=>{
    		return 选择对应的仓库状态
    	})
    	
    	返回值就是状态
    */
    const count = useSelector(store=>store.count)
    return <p>{count}</p>
}

export default City;
```

**解读**

让我们回忆以下这个

```js
const rootReducer = combineReducers({
    count:countReducer,
    person:personReducer
})

//用rootReducer来注册仓库实例
const store = createStore(rootReducer)
```

然后我们的状态就会是

```js
//大对象中的键是根据combineReducers中键得来
{
    count:{
        val:xx,
        sum:xx
    },
    person:{
        name:xx,
        age:xx
    }
}
```

我们未注入时会这样来访问状态

```js
//获取状态大对象
store.getState()

//获取到对应的count
store.getState().count

store.getState().count.sum
```

当我们注入后就可以使用钩子获取对应状态或状态值

```js
const count = useSelector(store=>store.count)

const sum = useSelector(store=>store.count.sum)
```

>使用react-redux还有一层原因，它内部对ContextAPI的访问做了优化

#### **分派任务**

```js
//引入读取状态的钩子
import {useSelector} from "react-redux"
//获得分派任务的钩子
import {useDispatch} from "react-redux"

function City(){
    const count = useSelector(store=>store.count)
    
    //得到分派函数
    /*
    	语法：
    	const 分派函数别名 = useDispatch()
    */
    const dispatch = useDispatch()
    
    function Add(){
        dispatch({type:"count/add",payload:5})
    }
    
    return <p>{count}</p>
}

export default City;
```

让我们回顾一下纯redux使用dispatch，

```js
//引入store
import store from "./xx"

//分派任务
store.dispatch(action)
```

所以React-redux只是做了一下封装，让我们从钩子获取这个派发函数

**注意事项**

我们经常把useReducer的第二个参数也命名为dispatch，所以为了不产生混淆，应该见名知意，我的方案是

```jsx
//对于useReducer的dispatch和reducer都带上状态名前缀，可读性高
const [account,accountDispatch] = useReducer(accountReducer,0)

//对于redux分发就用store当前缀
const storeDispatch = useDispatch()
```

具体的看项目组的规范，对于别人的项目只能结合内容来理解

#### 使用Action creator

实际上这个没有什么好说的，因为它没有任何封装，就是从文件中引入并执行而已，例如我们的countSlice.js

```js
const initCount = {
    val:0,
    sum:0
}

function countReducer(state=initCount,action){
    switch(action.type):
        case "count/add":
        	return {...state,val:state.val+action.payload}
    	case "count/dec":
    		return {...state,val:state.val-action.payload}
    	default
    		return initCount
}

//创造action的函数
//将action函数分别抛出
export function add(num){
    return {type:"count/add",payload:num}
}

export function dec(num){
    return {type:"count/dec",payload:num}
}

//默认抛出reducer
export default countReducer
```

我们在组件中使用action创造器，也就是调用那两个抛出的函数

```jsx
//引入函数
import {add} from "./features/counts/countSlice.js"
function City(){
    const dispatch = useDispatch()
    dispatch(add(5))
    
    return <p>...</p>
}
```

#### 糟糕的connect

connect是react-redux提供的一个API，我们可能会在老项目中看到他，这个是我们在有useSelector钩子之前获取状态的一种方式

```jsx
import {connect} from "react-redux"

function City({sum}){
    return <p>...</p>
}

//从store中读取状态
function mapStateProps(store){
    return {
        sum:store.count.sum
    }
}


//这是对City组件进行封装返回一个新组件
//connect(要传给组件的props)(要接收的组件)
export default connect(mapStateProps)(City)
```

实际上就是从仓库中取出值并把值作为props传递给组件，然后组件就可以使用仓库里的值

了解曾经有这种做法就可以了。

### redux中间件

我们可以通过redux中间件来扩展redux的功能，

我们来看一下需要中间件的原因，设想一下如果我们想对一些API进行异步调用，我们该在redux中如何实现，

>因为reducer函数是不能有副作用和异步操作的
>
>而Redux store本身也只有分派操作和更新状态的功能

我们传统的做法是把异步请求操作放到组件中，然后在需要修改状态的时候再触发派发操作，例如

```jsx
function App(){
    useEffect(function（){
        async function test(){
            const res = await fetch("xxxx")
            ...
            //在请求成功或失败要更新状态时,派发操作
            dispatch(xxx)
        }
        test()
    },[])
    
    return <p>...</p>
}
```

这没有问题，但有时候我们想保持组件的清洁和无数据获取，我们还希望将获取重要数据的逻辑封装起来，让它再项目中通用，而不是分布在各个组件。这样的话，我们更希望在store中能存储这些逻辑。

**中间件**

中间件`middleware`就是一个处于dispatch和store之间的函数，我们可以将异步操作放在这里。

以前的流程是:

```sh
dispatch后，action被推送到store，然后store选择合适的reducer执行，获得新状态
```

中间件加入后

```sh
dispatch后，将action等先交给`middleware`处理，然后再将处理结果交给store，store选择合适的reducer执行，获得新状态
```

所以中间件是我们处理异步操作、API请求的好地方。

我们可以自己编写中间件函数，但是很多时候redux开发者会选择第三方包来实现，redux中最流行的第三方库是`Redux Thunks`

#### Redux Thunks

基本上分为三步

```sh
1.安装所需包
2.将中间件应用到store
3.在操作创建者函数中使用中间件
```

安装react-thunk

```sh
npm i react-thunk
```

将中间件应用到store,来到`store.js`

```jsx
import {createStore,combineReducers,applyMiddleware} from "redux"

//引入thunk
import thunk from "react-thunk"


//引入对应的reducer
import countReducer from "./features/counts/countSlice.js"
import personReducer from "./features/counts/personSlice.js"

const rootReducer = combineReducers({
    count:countReducer,
    person:personReducer
})

/*
	createStore的第一个参数是reducer
	第二个参数是中间件
	
	
	applyMiddleware()，这是redux提供的一个方法，用于应用中间件
*/
const store = createStore(rootReducer，applyMiddleware(thunk))

//导出仓库
export default store
```

**在操作创建者函数中使用中间件**

这一步很重要，曾经无作用的Action creator总算要发挥左右了，

```sh
1.如果返回一个非函数，就会被认为是action
2.如果返回一个函数，就会被thunk识别为中间件函数，该函数第一个参数是派发操作，第二个参数用于获取当前状态
```

例如在`./features/counts/countSlice.js`中

```js
const initCount = {
    val:0,
    sum:0
}

function countReducer(state=initCount,action){
    switch(action.type):
        case "count/add":
        	return {...state,val:state.val+action.payload}
    	case "count/dec":
    		return {...state,val:state.val-action.payload}
    	default
    		return initCount
}

//此处的current是从外部传进来的，是当前状态
export function add(num，current){
    //根据条件来决定是直接返回action还是进入中间件
    //例如，如果当前数字小于10就直接返回，大于10就进入中间件
    if(current<10){
       return {type:"count/add",payload:num}
    }
   
    return async function(dispatch,getState){
        //异步操作或者副作用
        const res = await fetch(`xxx`)
        const data = await res.json()
        let newNum = num+data
        //再次派发action
        dispatch({type:"count/add",payload:newNum})
    }
}

export function dec(num){
    return {type:"count/dec",payload:num}
}

export default countReducer
```

**current的来源**

我们刚才的当前状态实际上是从外部传入的，例如

```jsx
import {add} from "./features/counts/countSlice.js"
function City(){
    const dispatch = useDispatch()
    const currentNum = useSelector(store=>store.count.val)
    
    dispatch(add(5,currentNum))
    return <p>...</p>
}
```

**异步请求没完成前的处理**

既然是异步请求就可能面临一个问题，请求耗时长此时我们应该有个等待的过程，解决思路，增加一个事件，来控制isLoading，例如

```js
const initCount = {
    val:0,
    sum:0,
    //组件根据这个值来知道请求是否完成
    isLoading:false
}

function countReducer(state=initCount,action){
    switch(action.type):
        case "count/add":
        	return {...state,val:state.val+action.payload,isLoading:false}
    	case "count/dec":
    		return {...state,val:state.val-action.payload,isLoading:false}
    	case "count/loading":
    		return {...state,isLoading:true}
    	default
    		return initCount
}

export function add(num，current){
    if(current<10){
       return {type:"count/add",payload:num}
    }
   
    return async function(dispatch,getState){
        //异步请求开始，进入loading状态
        dispatch({type:"count/loading"})
        
        const res = await fetch(`xxx`)
        const data = await res.json()
        let newNum = num+data
        //请求结束，再次派发更新操作
        dispatch({type:"count/add",payload:newNum})
    }
}

export function dec(num){
    return {type:"count/dec",payload:num}
}

export default countReducer
```

我们在页面上只需要使用`count.isLoading`即可，因为根据订阅的机制，它的变化会自动同步到消费者，并且触发re-render，

>别忘了在请求成功后还原isLoading为false

### Redux DevTools

安装redux开发者工具也需要三个步骤

```sh
1.安装谷歌拓展--Redux DevTools
2.在项目中装包：npm i redux-devtools-extension
3.在store文件中使用
```

在store文件中使用redux开发者工具

```js
import {createStore,combineReducers,applyMiddleware} from "redux"
import thunk from "react-thunk"

import {composeWithDevTools} from "redux-devtools-extension" 
import countReducer from "./features/counts/countSlice.js"
import personReducer from "./features/counts/personSlice.js"

const rootReducer = combineReducers({
    count:countReducer,
    person:personReducer
})

//用开发者工具提供的函数再把中间件方法包起来
const store = createStore(rootReducer，composeWithDevTools(applyMiddleware(thunk)))

export default store
```

再次打开浏览器我们会发现和Network同级有一个redux目录，这就是redux开发者工具

## Redux Toolkit

现代Redux，简称RTK

>RTK是脱胎于传统redux，所以它兼容传统redux的写法

RTK相对于传统redux的优势

```sh
1.首先说明一点RTK是redux社区最佳实践的整合体
2.于传统redux相比，RTK允许我们编写更少的代码来实现和redux一样的效果
3.RTK内置了中间件和开发工具，让我们无需额外安装
```

三个亮点:

```sh
1.We can write code that "mutates" state inside reducers(will be converted to immutable logic behind the scenes by "immer" library)

2.Action creators are automatically created

3.Automatic setup of thunk middleware and DevTools
```

由于RTK和传统redux兼容，所以我们基于传统redux项目来渐进升级。

### store的重新书写

安装RTK

```sh
npm i @reduxjs/tookit
```

我们先来看看传统redux书写的store.js

```js
import {createStore,combineReducers,applyMiddleware} from "redux"
import thunk from "react-thunk"

import {composeWithDevTools} from "redux-devtools-extension" 
import countReducer from "./features/counts/countSlice.js"
import personReducer from "./features/counts/personSlice.js"

const rootReducer = combineReducers({
    count:countReducer,
    person:personReducer
})


const store = createStore(
    //设置根reducer
    rootReducer，
    //设置开发者工具和中间件
    composeWithDevTools(applyMiddleware(thunk))
)

export default store
```

在新的RTK中

```js
import {configureStore} from "@reduxjs/toolkit"

import countReducer from "./features/counts/countSlice.js"
import personReducer from "./features/counts/personSlice.js"

//只要传入一个对象就可以，不需要调用别的钩子构造根路由
//thunk和redux devtools也被内置了
const store = configureStore({
 	reducer:{
        count:countReducer,
    	person:personReducer
    }   
})

export default store
```

### reducer的重新书写

由于我们的命名是xxSlice,所以很多时候会被称为xx切片，我们来看传统redux书写的形式

```js
//初始值
const initCount = {
    value:0,
    sum:0,
    isLoading:false
}

//reducer函数
function countReducer(state=initCount,action){
    switch(action.type):
        case "count/remote":
        	return {...state,value:state.value+action.payload,isLoading:false}
    	case "count/dec":
    		return {...state,value:state.value-action.payload,isLoading:false}
    	case "count/multi":
    		return {...state,value:state.value*action.payload,isloading:false}
    	case "count/division":
    		return {...state,value:state.value/action.payload,isloading:false}
    	case "count/loading":
    		return {...state,isLoading:true}
    	case "count/stay":
    		return {...state}
    	default
    		return initCount
}

//action构造函数
//含有异步操作
export function remote(num，current){
    if(current<10){
       return {type:"count/remote",payload:num}
    }
   
    return async function(dispatch,getState){
        dispatch({type:"count/loading"})
        
        const res = await fetch(`xxx`)
        const data = await res.json()
        let newNum = num+data
        dispatch({type:"count/remote",payload:newNum})
    }
}
//action构造函数
export function dec(num){
    return {type:"count/dec",payload:num}
}

export function multi(num){
    return {type:"count/multi",payload:num}
}

export function division(num){
    //如果除数为0，保持状态不变
    if(num===0){
        return {type:"count/stay"}
    }
    return {type:"count/division",payload:num}
}
export default countReducer
```

对于传统redux而言，他对切片实际上没有什么约束，我们也不需要用到钩子，而RTK中它提供了一个钩子来做这件事，这个钩子就是`createSlice`,该钩子有三大好处

```sh
1.他会从我们的reducer中自动创建Action creator
2.它使得编写reducer变得容易，因为我们不用再写switch语句，默认情况下也会自动处理
3.我们可以变异reducer内的状态，因为RTK幕后会有一个immer库把我们的逻辑换回不可变逻辑
```

现代RTK的书写,待会儿再看异步处理的情况

```js
import {createSlice} from "@reduxjs/toolkit"

//初始值没有发生变化
const initCount = {
    value:0,
    sum:0,
    isLoading:false
}

/*
下面写法大变样，
1.首先我们要构造一个XxSlice
2.createSlice要接收一个配置对象
*/
const countSlice = createSlice({
    //这里是切片别名，也就是 "count/dec"的第一部分
    //这是一个约定，方便待会生成事件
    name:"count",
    //初始状态值
    initialState:initCount,
    //reducers，之所以是复数，是因为内部就是action函数变种
    reducers:{
        //第一个参数是当前状态，第二个就是派发的action
        dec(state,action){
            //这就是刚才说的状态变异(翻译问题)
            //就是说我们可以直接操作状态
            state.value-=action.payload
            state.isLoading = false
        },
        multi(state,action){
            state.value = state.value*action.payload
            state.isLoading = false
        },
        division(state,action){
            //return就是状态不变
            //这要比之前我们单独写一个事件来保证状态不变简单太多了
            if(action.payload===0) return 
            state.value = state.value/action.payload
        }
    }
})

/*
	看看countSlice长啥样
	{
		name:"count",
		actions:{
			dec:f(),
			multi:f(),
			division:f()
		},
		reducer:f(),
		getInitialState:f()
		caseReducers:{...}
	}
*/
console.log(countSlice)

//分别暴露action函数
export const {dec,multi,division} = countSlice.actions

//默认暴露reducer
export default countSlice.reducer
```

现在让我们来消化一下变化。

```sh
【不用手动再写事件名】
在RTK中我们不用手动再写形如"count/dec"这样的事件名了，他会通过name和reducer的方法名自动拼接得到事件名

【不用在拓展整个状态】
之前的方法中需要 "{...state,value:state.value+action.payload}"这种形式来修改状态
现在只需要直接操作，让RTK底层去做刚才那种变化的事情

【reducer内部写法的变化】
reducer内部现在是大变样，返回一个个方法，这些方法都接收两个参数，第一个是当前state，第二个是派发的action
```

#### RTK中reducers的局限性

让我们来看一个案例，传统redux实现的一个效果

```js
const initCount = {
    value:0,
    sum:0,
    isLoading:false
}

//reducer函数
function countReducer(state=initCount,action){
    switch(action.type):
        case "count/twoParams":
			return {...state,value:action.payload.value,sum:action.payload.sum}
    	default
    		return initCount
}

export twoParams(val,sum){
    return {
        type:"count/twoParams",
        payload:{
            value:val,
            sum:sum
        }
    }
}

export default countReducer
```

所以我们当初使用的时候是这样使用的

```jsx
import {twoParams} from "./features/counts/countSlice.js"
function City(){
    const dispatch = useDispatch()
    //即我们传了两个参数
    dispatch(twoParams(10,100))
    return <p>...</p>
}
```

当我们用RTK改造后

```js
import {createSlice} from "@reduxjs/toolkit"

const initCount = {
    value:0,
    sum:0,
    isLoading:false
}

const countSlice = createSlice({
    name:"count",
    initialState:initCount,
    reducers:{
        //我们增加的部分
        twoParmas(state,action){
            //我们接收到的是 dispatch(twoParams(10,100))传递的第一个参数
            //即action为{payload:10}
            console.log(action)
        }
    }
})

export const {twoParmas} = countSlice.actions


export default countSlice.reducer
```

这就是RTK的局限性，它规定的形式，但是我们是有处理方式的，

```js
import {createSlice} from "@reduxjs/toolkit"

const initCount = {
    value:0,
    sum:0,
    isLoading:false
}

const countSlice = createSlice({
    name:"count",
    initialState:initCount,
    reducers:{
        twoParmas：{
            prepare(val,sum){
   				return {
             		payload:{
                 		value:val,
                 		sum:sum
             		}
         		}
			},
            reducer(state,action){
    			state.value = action.payload.value
    			state.sum = action.payload.sum
			}
        }
    }
})

export const {twoParmas} = countSlice.actions
export default countSlice.reducer
```

**解读**

```js
 twoParmas：{
     //对action的payload进行重构，接收外部传入的多个参数
     prepare(val,sum){
         return {
             payload:{
                 value:val,
                 sum:sum
             }
         }
     },
     //reducer函数部分
     reducer(state,action){
         state.value = action.payload.value
         state.sum = action.payload.sum
     }
 }
```

### RTK中异步的实现

RTK中提供了自己的异步方法，不过它实现起来比较复杂，所以我们更希望直接混用传统的异步实现，例如

```js
import {createSlice} from "@reduxjs/toolkit"

const initCount = {
    value:0,
    sum:0,
    isLoading:false
}

const countSlice = createSlice({
    name:"count",
    initialState:initCount,
    reducers:{
        dec(state,action){
            state.value-=action.payload
            state.isLoading = false
        },
        multi(state,action){
            state.value = state.value*action.payload
            state.isLoading = false
        },
        division(state,action){
            if(action.payload===0) return 
            state.value = state.value/action.payload
            state.isLoading = false
        },
        loading(state){
            state.loading = true 
        },
        remote(state,action){
            state.value +=action.payload
            state.isLoading = false
        }
    }
})

//混用传统的异步处理
//这里面的事件和传参都不用改，他会去reducers中寻找对应的事件
export function remote(num，current){
    if(current<10){
       //他会自动去寻找reducers中的remote方法
       return {type:"count/remote",payload:num}
    }
   
    return async function(dispatch,getState){
        //他会自动去寻找reducers中的loading方法
        dispatch({type:"count/loading"})
        
        const res = await fetch(`xxx`)
        const data = await res.json()
        let newNum = num+data
        dispatch({type:"count/remote",payload:newNum})
    }
}

export const {dec,multi,division} = countSlice.actions

export default countSlice.reducer
```

未来我们会研究RTK提供的异步处理，因为那样的代码具有一致性

## 回顾Redux和RTK

Redux是自由的，RTK实际上就是对Redux的封装，加了一些预制条件，使得代码更简单，但是灵活性其实变低了，属于用效率换了灵活性。

## Redux与ContextAPI

这两者经常被人拿来比较，例如

```sh
【ContextAPI+useReducer】
这个组合常被用来作为Redux的替代品

优势:
1.已经内置在react中，不需要额外安装
2.设置单个context非常容易

缺点:
1.当我们需要很多个context Provider时，按照我们的封装方式，每个都需要暴露一个probider和一个自定义钩子，这被戏称为"Provider 地狱"
2.没有内置机制来实现异步操作
3.优化ContextAPI和reducer需要我们做一些工作
```

```sh
【Redux】
优点:
1.一旦完成初始化，设置不同的Slice将变得很容易
2.支持中间件的使用，我们可以用中间件来处理异步
3.提前进行了性能优化
缺点:
1.需要额外安装包，使得bundle体积变大
2.需要更多的工作来完成初始化
```

**使用建议**

```sh
【ContextAPI+useReducer】
1.用于在小型应用程序中管理global state
2.如果是共享一个不常改变的值，推荐使用这个组合，例如颜色主题，用户首选语言，已验证用户等等，因为在这种情况下没必要优化
3.解决简单的多层props传递
4.管理local sub-tree的时候，就是祖先往子代通信时
```

```sh
【Redux】
1.在大型应用中使用它来管理global state
2.许多需要被频繁更新的内容
3.当state是复杂的对象或者数组时使用RTK是非常方便的
```







