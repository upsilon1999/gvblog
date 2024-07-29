# 单元重点

```sh
1.解释wasted renders，学会优化浪费的呈现和性能
2.使用更先进的React工具，让程序尽可能快
3.更深入的使用useEffect
```

**React开发工具**

React开发工具按照后会给浏览器添加两个选项`Components`和`Profiler`，第一个就是组件相关内容，可以看到组件树，组件内部的各种状态，第二个就是探查器，我们优化时会用到的。

>Profiler可以帮助我们分析render和re-render，我们可以看到哪些组件渲染了，耗时多少

更多详细解读查看p244

## 优化性能和浪费的呈现

当我们要优化React的性能时一般会关注三个领域

```sh
1.prevent wasted render防止浪费渲染
【内置React工具】
a.memo记忆组件，
b.useMemo钩子，可以使用它记忆对象和函数
c.useCallback钩子
d.将组件作为children或者普通props传递到其他组件中


2.提高整个程序的整体速度和响应能力，以确保应用程序时百分之百流畅的,没有延迟
a.useMemo钩子
b.useCallback钩子
c.useTransition钩子

3.Reduce Bundle size减小束的大小
a.在代码中使用更少的第三方包
b.代码拆分和懒加载
```

还有很多优化性能的技巧

```sh
a.不要在React组件中嵌套声明其他React组件(重复渲染问题)
```

### 1.什么是浪费呈现

我们先得理解什么是`wasted render`，我们首先要知道组件实例何时被重新呈现，在React中,组件实例只能在三种不同的情况下重新呈现

```sh
【state changes】
组件的state改变时

【Context changes】
上下文改变时，订阅了上下文的消费者组件都会重新渲染

【Parent re-renders】
父组件重新渲染时，
```

**常见误区**

```sh
误区1：
不要误解为props改变，实际上是的顺序是父组件重新渲染可能导致props改变，有时候父组件重新渲染，props实际上是不变的，但是仍然会让子组件重新渲染

误区2：
渲染组件并不意味着DOM得到了更新，渲染只意味着组件函数被调用，React将创建一个新的虚拟DOM
```

**什么是wasted render**

```sh
"一个没有在DOM中产生任何变化的渲染"，这是一种浪费，因为所有不同的计算逻辑仍然必须运行
在一般情况下，这其实不构成问题，因为React很快，但是当数量庞大或者某些组件响应特别慢时，会使应用程序感到滞后和无响应，例如当用户执行某个操作后更新用户界面的速度不够快
```

## 优化方法1之利用子prop防止组件重新渲染

这不是现代常用的技术，但它提供了一种思路让我们了解React内部是如何工作的，例如我们准备了一个Test组件

```jsx
import { useState } from "react";

//这个组件会生成10万条数据，
function SlowComponent() {
  // If this is too slow on your maching, reduce the `length`
  const words = Array.from({ length: 100_000 }, () => "WORD");
  return (
    <ul>
      {words.map((word, i) => (
        <li key={i}>
          {i}: {word}
        </li>
      ))}
    </ul>
  );
}

export default function Test() {
  const [count, setCount] = useState(0);
  return (
      <div>
          <h1>Slow counter?!?</h1>
          <button onClick={() => setCount((c) => c + 1)}>Increase: {count}</button>
		  {/*我们将会生成10万条数据的组件作为了Test的子组件*/}
          <SlowComponent />
      </div>
    );
}
```

我们将生成10万条数据的SlowComponent组件作为了Test的子组件，会发生以下事件

```sh
1.当Test中的按钮按下触发点击事件，使得count状态更新
2.count状态更新使得Test重新渲染，父组件的重新渲染使得子组件SlowComponent也重新渲染
3.由于生成十万条数据的普通函数会重新执行，然后触发map，重新生成，SlowComponent会大大消耗性能
```

这里产生的问题就是

```sh
1.SlowComponent组件根本没有依赖count状态，但它受到了该状态变化的影响
```

**使用children属性优化**

```jsx
function SlowComponent() {
  // If this is too slow on your maching, reduce the `length`
  const words = Array.from({ length: 100_000 }, () => "WORD");
  return (
    <ul>
      {words.map((word, i) => (
        <li key={i}>
          {i}: {word}
        </li>
      ))}
    </ul>
  );
}

function Counter({ children }) {
  const [count, setCount] = useState(0);
  return (
    <div>
      <h1>Slow counter?!?</h1>
      <button onClick={() => setCount((c) => c + 1)}>Increase: {count}</button>

      {children}
    </div>
  );
}

export default function Test() {
  return (
    <div>
      <h1>Slow counter?!?</h1>
      <Counter>
        <SlowComponent />
      </Counter>
    </div>
  );
}
```

组件树没有产生变化，但是我们用性能分析会发现当count状态发生变化时，SlowComponent组件没有被重新呈现，原因

```sh
这其实和JSX有关系，SlowComponent组件在第一次时被渲染好了，他被作为children属性传递给Counter组件。
Counter组件中count状态变化，只会让Counter组件重新渲染，而SlowComponent组件由Test组件负责渲染，他已经存在了，而且没有发生变化，仍然按之前的children属性进行传递，这也是为什么SlowComponent组件不会重新渲染。
```

即使不用children属性，用组件属性传递方式也是等价的

```jsx
function SlowComponent() {
  // If this is too slow on your maching, reduce the `length`
  const words = Array.from({ length: 100_000 }, () => "WORD");
  return (
    <ul>
      {words.map((word, i) => (
        <li key={i}>
          {i}: {word}
        </li>
      ))}
    </ul>
  );
}

function Counter({ slow }) {
  const [count, setCount] = useState(0);
  return (
    <div>
      <h1>Slow counter?!?</h1>
      <button onClick={() => setCount((c) => c + 1)}>Increase: {count}</button>

      {slow}
    </div>
  );
}

export default function Test() {
  return (
    <div>
      <h1>Slow counter?!?</h1>
      <Counter slow={SlowComponent} />
    </div>
  );
}
```

**原理**

```sh
count状态变化影响不到SlowComponent组件，它处在Test组件下。
Counter组件重新渲染前，SlowComponent组件早就渲染好了，Counter组件变化没有触发Test组件更新也没触发SlowComponent组件更新，所以不受影响。
```

## 优化方法2之记忆化工具

memo函数，useMemo、useCallback这三个方法后的基本概念都是记忆化

**什么是记忆化**

记忆化是一种优化技术，只执行一次纯函数。

>纯函数的一大特性:相同的输入只会获得相同的输出

```sh
假设函数A，他将结果存储在内存中或者说缓存中，
1.当我们以后尝试用相同的输入再次执行相同的函数时，它将简单地返回以前存储在缓存中的结果,因此函数不会再次执行，因为纯函数的特点，再次执行相同的输入是没有意义的
2.如果输入不同，那么函数当然会再次执行
```

基于记忆化我们有了以下三种不同处理方案

```sh
1.用memo函数记忆components
2.用useMemo记忆objects
3.用useCallback记忆 functions
```

这样做有以下两个好处

```sh
1.杜绝了wasted renders
2.提高了整体应用程序速度
```

### memo函数

它也被称作备忘录函数，我们可以使用这个函数创建一个不会重新呈现的组件，

```sh
1.当其父组件重新渲染时，只要传递的props在两次渲染之间保持不变它也不会改变，
原理是将组件函数类比了纯函数，props就是传入的参数，既然参数没变，就不需要重新渲染

2.只受到props的影响:
这里其实是笼统的说法，更完整的说法应该是只有该组件自身state变化、父组件传递的props变化、订阅的context变化三者才会引起组件的重新渲染
```

非Memo函数组件的行为

```sh
1.父组件重新渲染，子组件也会跟着重新渲染，即使props没变
```

**不要滥用memo组件**

尽管memo组件很好用，但是并不是所有组件都使用它，我们只在处理较重(heavy或者说slow rendering)的组件时才会使用它

>因为轻量级组件或者渲染快的情况下没必要浪费内存来记忆

#### memo函数练习

**不使用memo函数**

App父组件

```jsx
function App(){
    const [count,setCount] = useState(0)
    const [name,setName] = useState("lisa")
    return (
    	<div>
          <button onClick={()=>setCount(count+1)}/>
          <SlowBot name={name}>
        </div>
    )
}
```

SlowBot组件

```jsx
function SlowBot({name}){
    const words = Array.from({ length: 100_000 }, () => "WORD");
    return (
        <div>
        	你的名字:{name}
            <ul>
                {words.map((word, i) => (
                    <li key={i}>
                        {i}: {word}
                    </li>
                ))}
            </ul>
        </div>
  );
}
```

缺点:当count状态变化后，App会重新渲染，从而会触发子组件渲染

**使用memo函数**

基本语法

```jsx
const 新组件名 = memo(原组件)
```

示例

```jsx
import {memo} from "react"

const SlowBot = memo(function SlowBot({name}){
    const words = Array.from({ length: 100_000 }, () => "WORD");
    return (
        <div>
        	你的名字:{name}
            <ul>
                {words.map((word, i) => (
                    <li key={i}>
                        {i}: {word}
                    </li>
                ))}
            </ul>
        </div>
  );
})
```

规范

```sh
1.memo返回的组件名尽量与之前的组件名一样，这样不用改之前使用的地方
```

**解释**

```sh
实际上就是将要被记忆化的组件作为参数传入memo函数中
```

**简化写法1**

由于传入memo函数后，最终使用的组件名是按memo返回值来的，所以参数完全可以使用匿名函数

1.箭头函数形式

```jsx
import {memo} from "react"

//即将memo函数的参数改为匿名函数即可
const SlowBot = memo(({name})=>{
    const words = Array.from({ length: 100_000 }, () => "WORD");
    return (
        <div>
        	你的名字:{name}
            <ul>
                {words.map((word, i) => (
                    <li key={i}>
                        {i}: {word}
                    </li>
                ))}
            </ul>
        </div>
  );
})
```

2.无名函数模式

```jsx
import {memo} from "react"

//即将memo函数的参数改为匿名函数即可
const SlowBot = memo(function ({name}){
    const words = Array.from({ length: 100_000 }, () => "WORD");
    return (
        <div>
        	你的名字:{name}
            <ul>
                {words.map((word, i) => (
                    <li key={i}>
                        {i}: {word}
                    </li>
                ))}
            </ul>
        </div>
  );
})
```

### 理解useMemo和useCallback

现在让我们来看一个问题

```jsx
function App(){
    const [count,setCount] = useState(0)
    const option = {
        name:"lisa",
        flag:true
    }
    return (
    	<div>
          <button onClick={()=>setCount(count+1)}/>
          <SlowBot option={option}>
        </div>
    )
}
```

即使我们记忆了

```jsx
import {memo} from "react"

//即将memo函数的参数改为匿名函数即可
const SlowBot = memo(function ({option}){
    const words = Array.from({ length: 100_000 }, () => "WORD");
    return (
        <div>
        	你的名字:{option.name}
            <ul>
                {words.map((word, i) => (
                    <li key={i}>
                        {i}: {word}
                    </li>
                ))}
            </ul>
        </div>
  );
})
```

我们发现App重新渲染时，SlowBot还是被重新渲染了，原因是

```sh
option是一个对象，在App渲染时，这个对象被重新创建了，所以是个新对象，也就是说传入的prop变了
(真实原因是地址变了)
```

#### **原理解读**

每当组件实例重新渲染时，内部的一切都是重新渲染的，所以所有的值总是被再次创建，这包含组件中定义的对象和函数，所以一个新的渲染会得到新的函数和新的对象，即使它们和以前一模一样。

在JS中，两个看起来相同的对象或函数，看起来代码一样，实际上不同的唯一对象，这里经典的例子是，一个空对象不同于另一个空对象，即

```js
//a与b不一样
let a = {}
let b = {}
```

**推论**

如果我们将函数或对象作为prop传递给子组件,每当有重新渲染时，子组件总是将它们视为新的prop。

又由于重新渲染的prop不一样，memo函数就不管用了。

```sh
如果组件收到对象或者函数作为prop，在重新渲染时它们将会是新的prop，即使它们看起来一模一样
```

**解惑**

```sh
为什么对象、函数有这个问题，而state不会，原因:
在React重新渲染时，如果state没有发生变化，它就是个不变的值，这在之前我们就知道了。


其实很好理解，因为state的变化会触发重新渲染，如果每次重新渲染也能重创state，那就会陷入死循环。
```

**解决方案**

```sh
让函数和对象在两次渲染之间保持不变，除非它们的内容发生变化，也就是在两次渲染之间记住它们。
记住对象用useMemo
记住函数用useCallback
```

**小结**

```sh
1.在两次渲染之中，我们可以用useMemo来记住任何值以及用useCallback记住函数
2.无论我们传递什么值，useMemo和useCallback都会把它们记录在内存(或叫缓存)中，缓存的值将在以后的重新渲染中返回，因此只要保持输入不变，它们就将在re-render中保持不变
3.useMemo和useCallback也有一个依赖数组,当其中一个依赖项发生变化，存储的值就会re-created
```

**使用场景**

```sh
1.我们用memo函数记录一个重或者慢的组件，以后称备忘录组件，当我们要给这个备忘录组件传递对象或者函数时就需要用到useMemo和useCallBack
2.避免在每次渲染上进行昂贵的重新计算，例如某个我们生成的对象或者函数，我们不想让他被重新渲染影响。就是可以用useMemo和useCallBack实现数据持久化，让他不受组件重新渲染的影响。
3.记忆的值被另一个hook的依赖数组所使用(Memoizing values that are used in dependency array of another hook),例如避免useEffect无限循环
```

#### 使用useMemo

基本语法

```jsx
const 对象名 = useMemo(回调函数,依赖数组)
```

整体和useEffect很类似，例如

```jsx
function App(){
    const [count,setCount] = useState(0)
    const option = useMemo(()=>{
        return {
            name:"lisa",
            flag:true,
        }
    },[])
    return (
    	<div>
          <button onClick={()=>setCount(count+1)}/>
          <SlowBot option={option}>
        </div>
    )
}
```

**依赖数组详解**

```sh
1.如果依赖数组为[],则只在加载完成后生成一次，即使以后内容改变也不会改变对象
2.如果想要对象保持更新，应该对内部使用到的变量进行监听，当监听的变量更新时，对象也会对应更新
```

例如用到了某个状态，

```jsx
function App(){
    const [count,setCount] = useState(0)
    //用到了count状态，应该要监听count状态
    //如果不监听count，那么即使count状态发生了变化，option对象也不会改变
    const option = useMemo(()=>{
        return {
            name:"lisa",
            flag:true,
            count:count
        }
    },[count])
    return (
    	<div>
          <button onClick={()=>setCount(count+1)}/>
          <SlowBot option={option}>
        </div>
    )
}
```

尽量监听具体的内容，即使用到什么变量就监听谁

```jsx
function App(){
    const [count,setCount] = useState(0)
    const [posts,setPosts] = useState([])
    //对象内部用到的变量是posts.length,我们就应该监听它，而不是监听posts状态本身
    const option = useMemo(()=>{
        return {
            name:"lisa",
            flag:true,
            len:`邮寄长度是${posts.length}`
        }
    },[posts.length])
    return (
    	<div>
          <button onClick={()=>setCount(count+1)}/>
          <SlowBot option={option}>
        </div>
    )
}
```

#### 使用useCallBack

先来看背景条件，App组件和SlowBot组件

```jsx
function App(){
    const [count,setCount] = useState(0)
    const [posts,setPosts] = useState([])

    const handleAdd = (postItem)=>{
        setPosts([...posts,postItem])
    }
    return (
    	<div>
          <button onClick={()=>setCount(count+1)}/>
           {posts.map((item,index)=>{
                return <p key={index}>{item}</p>
            })}
          <SlowBot onhandleAdd={handleAdd}>
        </div>
    )
}
```

```jsx
import {memo} from "react"

const SlowBot = memo(function ({onhandleAdd}){
    const words = Array.from({ length: 100_000 }, () => "WORD");
    return (
        <div>
        	
            <ul>
                {words.map((word, i) => (
                    <li key={i}>
                        {i}: {word}
                        <button onClick={()=>onhandleAdd(word)}></button>
                    </li>
                ))}
            </ul>
        </div>
  );
})
```

>功能与问题

```sh
【功能】
当按钮点击时在App列表中加入一条数据，

【问题】
App中count更新导致组件渲染，函数handleAdd被重新创建，SlowBot识别到prop变化也重新渲染了。
```

