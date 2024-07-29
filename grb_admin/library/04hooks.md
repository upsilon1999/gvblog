# hooks

本单元重点

```sh
1.hooks学起来容易，掌握起来难
2.hooks的规则
3.深度分析useState
4.useRef
5.自定义钩子
```

## 技巧

### 局部禁用 eslint

这是一段有问题的代码，eslint 会检测报错，例如

```jsx
function App() {
	if (a > 10) [isTop, setIsTop] = useState(true);
}
```

由于没有在顶层使用 hook，所以 eslint 会报错，我们可以用一个注释来解除，例如

```jsx
function App() {
	/* eslint-disable */
	if (a > 10) [isTop, setIsTop] = useState(true);
}
```

## React hooks

**本质**

```sh
1.React钩子本质上时内置在React中的特殊函数，它允许我们连接到React的一些内部机制，或者说React hooks就是公开的一些内部功能API
2.共同点，使用"use"开头，为了与我们其他的函数区别开来
3.我们可以创建自定义钩子，推荐使用use开头，自定义钩子的基本目的是重用非可视化逻辑，这逻辑要与UI无关
```

react 内置了 20 多个 hook，最常用的有

```sh
【useState】
创建并修改状态
【useEffect】
副作用钩子，可以监听属性和实现生命周期
【useReducer】
【useContext】
```

较常用的有

```sh
【useRef】
【useCallback】
【useMemo】
【useTransition】
【useDeferredValue】
```

### 使用 hooks 的规则

一般我们都需要遵守以下两条规则
**规则 1**

```sh
hook只在顶层使用，即不能在if条件语句、循环体、函数体中使用，也不应该提前return
```

原因是:hook 只有在他们总是以完全相同的顺序被调用时才起作用，简单来说，就是只有在组件顶层(即组件全局，不被其他包含)才起作用
深层原因参见 P160，因为 hook 是以链表结构存储的，上一个 hook 指向下一个 hook，如果有 hook 在 if、循环、函数这些中，就可能会造成链表的断裂,我们可以违反规则并制造链表断裂来查看报错信息。

解释一下提前 return

```jsx
function App(){
	if(a>100) return <p>hello</p>
	useEffect(...)
}
```

return 之后还注册有别的钩子，也容易造成列表断裂
**规则 2**

```sh
hook只能在React组件或自定义hook中调用
```

即不能是常规函数，甚至不能是类组件

## useState 的更多细节

### useState 的初始值只与初始渲染有关

useState 的初始值只和初始渲染有关，举例

```jsx
function App() {
	const [isTop, setIsTop] = useState(score > 8);
	console.log(isTop);
}
```

我们预想的是，根据 score 状态的变化，从而更新值 isTop，实际上初始渲染时 score 如果是 undefined，那么 isTop 为 false，即使后面 score 变化了，也不会影响到 isTop，即状态只被初次渲染影响，后续更新状态需要使用对应的 setIsTop，可以结合 useEffect 监听对应属性来刷新 state，例如

```jsx
function App() {
	const [isTop, setIsTop] = useState(score > 8);
	console.log(isTop);
	useEffect(
		function () {
			setIsTop(score > 8);
		},
		[score]
	);
}
```

其实对于这种情况我们更多的时候应该使用派生状态，而不是用 useEffect 去监听某个属性变化，例如
因为 isTop 和状态 score 有关，应该直接将他写成 score 的派生状态，那么 score 的每一次变化都将同步影响到 isTop

```jsx
const isTop = score > 8;
```

### 用回调函数来处理异步的情况

我们知道状态的设置是异步的，实际上事件的处理也是异步的，例如

```jsx
function App() {
	const [count, setCount] = useState("1");
	const [avg, setAvg] = useState(0);
	function add() {
		setAvg(Number(count));
		// 拿到的是同步的count
		setAvg(count * 2);
	}
	return <button onClick={add} />;
}
```

我们想在点击按钮时实现，avg 的值先设定为将 count 数字化，然后再乘以 2,要实现这样连续的效果就需要回调函数，举例

```jsx
function App() {
	const [count, setCount] = useState("1");
	const [avg, setAvg] = useState(0);
	function add() {
		setAvg(Number(count));
		setAvg((count) => count * 2);
	}
	return <button onClick={add} />;
}
```

### 更新 state

用一个新值更新 state

```jsx
//例如
setCount(1000);
```

使用当前状态更新 state

```jsx
setCount((count) => count + 50);
```

当我们更新的状态时数组或对象时，一定要用一个新的数组或对象去更新，例如

```jsx
const [person, setPerson] = useState({ name: "李四", age: 19 });
// 使用一个新对象去更新
setPerson({ ...person, name: "张三" });
```

## 本地存储 state

需求：

```sh
1.将数据存储到浏览器缓存中
2.每次更新列表数据时更新缓存数据
3.组件挂载时从缓存中读取之前的数据
```

### 将数据存入缓存

这可以在两个地方实现，一个是在事件处理程序中，在事件发生时更新缓存，此时需要注意由于更新是异步执行所以存入时需要存入新数据，二是使用 useEffect，此时监听 watched，传入的数据就是新的 watched
**方式 1**

```jsx
function App() {
	const [watched, setWatched] = useState([]);
	function handleAddWatched(movie) {
		setWatched((watched) => [...watched, movie]);

		function handleAddWatched(movie) {
			setWatched((watched) => [...watched, movie]);

			// 不直接用watched的原因，由于状态异步更新，这里的watched还是旧数据
			// localStorage.setItem("watched",watched)

			// 由于本地存储中只能用字符串存储，所以需要转成字符串
			localStorage.setItem(
				"watched",
				JSON.stringify([...watched, movie])
			);
		}

		// 这里只是简单写了个点击按钮，实际功能中是复杂数据
		return <button onClick={handleAddWatched(movieInfo)} />;
	}
}
```

**方式 2**
使用 useEffect 存储

```jsx
function App() {
	const [watched, setWatched] = useState([]);
	function handleAddWatched(movie) {
		setWatched((watched) => [...watched, movie]);
	}

	// 由于watched监听得到的就是新数据，直接存就行
	useEffect(
		function () {
			localStorage.setItem("watched", JSON.stringify(watched));
		},
		[watched]
	);
	// 这里只是简单写了个点击按钮，实际功能中是复杂数据
	return <button onClick={handleAddWatched(movieInfo)} />;
}
```

> 拓展

useEffect 如何获得旧的值,可以在外部设定一个变量，来接收清理函数中的值，例如

```jsx
function App() {
	let oldData;
	useEffect(
		function () {
			// 初次打印这是初始值，例如undefined
			//当更新时，由于清理函数中我们将上次的值赋予了它
			//所以oldData将是上次的值
			console.log(oldData);

			return function () {
				oldData = query;
			};
		},
		[query]
	);
}
```

### 挂载时从缓存中读回数据

```jsx
function App() {
	const [watched, setWatched] = useState(function () {
		const storedValue = localStorage.getItem("watched");
		// 由于我们存储的时候转成了JSON字符串，所以取的时候要转回来
		return JSON.parse(storedValue);
	});
}
```

我们可以用一个函数来返回复杂的初始值，不要采取以下做法，因为那样可能无效

```jsx
//无效的示范
function App() {
	// 这样可能无效
	const [watched, setWatched] = useState(
		JSON.parse(localStorage.getItem("watched"))
	);
}
```

其次很重要的一点，当我们使用函数来返回初始值时，必须使用纯函数，即没有参数的函数，例如

```jsx
function App() {
	// 使用无参的普通函数
	const [watched, setWatched] = useState(function () {
		const storedValue = localStorage.getItem("watched");
		return JSON.parse(storedValue);
	});

	// 使用无参的箭头函数
	const [watched, setWatched] = useState(() => {
		const storedValue = localStorage.getItem("watched");
		return JSON.parse(storedValue);
	});
}
```

**小结**

```sh
1.函数必须没有参数，即使用纯函数
2.必须有返回值来作为初始值
```

## useRef

在引出这个钩子前,我们先来实现一个需求,选中某个 dom 元素,例如我们要在 search 组件挂载后选中 input 输入框，让他获得焦点,按照传统做法

```jsx
function Search({ query, setQuery }) {
	useEffect(function () {
		// 传统做法，选中输入框并让他在初始加载时获得焦点
		const el = document.querySelector(".search");
		el.focus();
	}, []);
	return (
		<input
			className="search"
			type="text"
			placeholder="Search movies..."
			value={query}
			onChange={(e) => setQuery(e.target.value)}
		/>
	);
}
```

这种做法虽然能实现效果，但是需要有类名、id 或者其他很多多余的东西，也和 React 的风格不统一，所以 React 提供了一个操作 DOM 的钩子

### useRef 基础了解

1.引入钩子，不用说了

```jsx
import { useRef } from "react";
```

2.给定初始值，

```jsx
//myRef自定义的标识名
const myRef = useRef(23);
```

这里的初始值，实际上是`标识.current`的属性值，它是我们给定一个可变值

3.两个大的用处

```sh
1.创建一个变量，并在不同的渲染中保持不变
2.选择并存储一个DOM元素
```

4.使用场景

```sh
1.useRef通常只出现在事件处理程序或效果中，一般不出现在jsx中
2.不允许在呈现逻辑中写入或读取useRef
```

### refs 和 state 的异同

某种程度上 refs 很像 state，只不过能力很小。

```sh
1.他们都是跨render持久化，即组件会记住这些值，即使在重新渲染后，这里的区别就是，更新状态会导致组件重新render，但是更新refs不会
2.状态是不变的，refs不是
3.state是异步更新的，意味着我们不能在更新后立即使用新状态，但是refs不是，我们可以在更新一个新的当前属性后立即读取它
```

所以我们一般用 refs 回到之前记录的状态，并存储设置超时的 ID，即在 render 中间记住某些数据片段，同时每当我们更新 refs 时,不重新渲染用户界面，每当更新 refs 时

### useRef 操作 DOM

分为三个步骤，
步骤 1，使用 useRef 创建标识符，传入 current 的初始值，由于我们要操作 DOM，假定初始为空元素，即

```jsx
function App() {
	// 用refs操作DOM有三个步骤
	// 1.创建ref，传入current属性初始值，当我们使用DOM元素时，一般假设它是空的，所以
	const inputEl = useRef(null);
}
```

步骤 2，给要操作的元素绑定 ref 属性，属性值就是 refs 标识，设置成功后，refs 标识的 current 属性值就变成了该 DOM 元素

```jsx
function Search({ query, setQuery }) {
	// 用refs操作DOM有三个步骤
	// 1.创建ref，传入current属性初始值，当我们使用DOM元素时，一般假设它是空的，所以
	const inputEl = useRef(null);

	return (
		// 2.给我们要操作的DOM设置ref属性，属性值就是ref标识名
		//设置完成后，该标识的current属性值就变成了要操作的DOM元素
		<input
			className="search"
			type="text"
			placeholder="Search movies..."
			value={query}
			onChange={(e) => setQuery(e.target.value)}
			ref={inputEl}
		/>
	);
}
```

步骤 3,我们现在可以通过`标识符.current`来操作对应 DOM 元素,操作一般在 useEffct 或事件处理程序中进行，还有一点值得注意，DOM 必须在组件挂载完后才可访问

```jsx
function Search({ query, setQuery }) {
	// 用refs操作DOM有三个步骤
	// 1.创建ref，传入current属性初始值，当我们使用DOM元素时，一般假设它是空的，所以
	const inputEl = useRef(null);

	//3.一般我们对DOM的操作，搭配内部事件处理程序或者useEffect使用
	useEffect(function () {
		console.log(inputEl.current);
		inputEl.current.focus();
	}, []);

	return (
		// 2.给我们要操作的DOM设置ref属性，属性值就是ref标识名
		//设置完成后，该标识的current属性值就变成了要操作的DOM元素
		<input
			className="search"
			type="text"
			placeholder="Search movies..."
			value={query}
			onChange={(e) => setQuery(e.target.value)}
			ref={inputEl}
		/>
	);
}
```

> 拓展一个更复杂的功能，当我们在按下回车键时，搜索框获得焦点并且清空之前的搜索内容

```jsx
function Search({ query, setQuery }) {
	const inputEl = useRef(null);

	useEffect(function () {
		function callback(e) {
			if (e.code === "Enter") {
				inputEl.current.focus();
				//清空之前的搜索内容，
				//1.我们跳转过来不想获得之前输入的内容
				//2.执行搜索后，应该清空本次输入
				setQuery("");
			}
		}
		// 更复杂的功能，希望按下回车键就可以进行聚焦搜索框
		document.addEventListener("keydown", callback);
		return () => {
			document.removeEventListener("keydown", callback);
		};
	}, []);

	// 更复杂的功能，希望按下回车键就可以进行搜索

	return (
		<input
			className="search"
			type="text"
			placeholder="Search movies..."
			value={query}
			onChange={(e) => setQuery(e.target.value)}
			ref={inputEl}
		/>
	);
}
```

> 优化，如果已经回车获取焦点的事情已经触发，不应该重复触发

这里可以用到 document 的 activeElement 属性，如果现在激活的 DOM 是输入框，则不应该再次聚焦，例如

```jsx
function Search({ query, setQuery }) {
	const inputEl = useRef(null);

	useEffect(function () {
		function callback(e) {
			// 如果当前元素已激活就不再重复触发事件，该功能刷新页面后生效
			if (document.activeElement === inputEl.current) return;

			if (e.code === "Enter") {
				inputEl.current.focus();
				//清空之前的搜索内容，
				//1.我们跳转过来不想获得之前输入的内容
				//2.执行搜索后，应该清空本次输入
				setQuery("");
			}
		}
		// 更复杂的功能，希望按下回车键就可以进行聚焦搜索框
		document.addEventListener("keydown", callback);
		return () => {
			document.removeEventListener("keydown", callback);
		};
	}, []);

	// 更复杂的功能，希望按下回车键就可以进行搜索

	return (
		<input
			className="search"
			type="text"
			placeholder="Search movies..."
			value={query}
			onChange={(e) => setQuery(e.target.value)}
			ref={inputEl}
		/>
	);
}
```

### useRef 使变量持久化

给定一个变量，让他在 render 之间持久化，
步骤 1：传入变量初始值

```jsx
function MovieDetails() {
	const countRef = useRef(0);
}
```

结合实际案例，评分条是可以多次点击的，只有点了提交之后评分才会确定
我们想让用户在每次评分后更新 countRef,记录评分次数

```jsx
function MovieDetails() {
	const countRef = useRef(0);
	useEffect(
		function () {
			if (userRating) {
				countRef.current += 1;
			}
		},
		[userRating]
	);
}
```

通过 React 工具，我们可以看到选择了三次评分后，countRef 变成了 3
**为什么不用变量**

```jsx
function MovieDetails() {
	let count = 0;
	useEffect(
		function () {
			if (userRating) {
				count += 1;
			}
		},
		[userRating]
	);
}
```

选择了三次评分后提交，count 为 1，原因:每次评分都会改变 userRating 的状态值，引起重新渲染，count 的值重置为 0，只有最后一次提交 count 变成了 1
**理解**
通过 refs 和普通变量的对比我们知道了，refs 不受重新渲染的影响，从而实现了数据的持久化，不用 state 的原因，state 的改变会再次触发渲染，从而陷入死循环

## 自定义 hook

一切都是由重用性引出的，对项目来说，我们一般重用的就两种:UI 和逻辑

UI 的重用通过组件来实现，逻辑的重现该如何实现，一般采用外置函数。

如果这个外置函数是一个常规函数，即不包含任何 React 内置钩子。如果包含，则应该创建一个自定义 hook

> 自定义 hook 就是允许我们在多个组件中允许我们复用多个有 React hook 的逻辑

规范：
一个自定义 hook 应该只做一件特定的、定义良好的事情，所以说这个想法不是简单地把组件地所有狗子变成一个特制的钩子，还应该使它便携可重用，甚至可以在完全不同的项目中使用

**一个自定义钩子案例**

```jsx
function useFetch(url) {
	const [data, setData] = useState([]);
	const [isLoading, setIsLoading] = useState(false);

	useEffect(function () {
		fetch(url)
			.then((res) => res.json())
			.then((res) => setData(res));
	}, []);

	return [data, isLoading];
}
```

首先自定义 hook 应该是一个 js 函数，它可以接收与返回与此自定义钩子相关的任何数据，事实上，从自定义钩子返回对象或数组是非常常见的

```sh
与组件的差异:组件也是一个函数，但是组件只能接收props，且组件总要返回一些JSX
与常规函数的区别：自定义hook需要使用一个或多个React hook，其次自定义hook应该以use开头
```

### 案例 1：useLocalStorageState

我们要创建一个自定义 hook，一般有两种思路:

```sh
1.我们想重用我们的非可视化逻辑的某些部分
2.我们想提取我们很大的组件的一部分
```

示例:

```jsx
import { useEffect, useState } from "react";
export function useLocalStorageState(initialState, key) {
	const [value, setValue] = useState(function () {
		const storedValue = localStorage.getItem(key);
		// 由于我们存储的时候转成了JSON字符串，所以取的时候要转回来
		return JSON.parse(storedValue) || initialState;
	});

	useEffect(
		function () {
			localStorage.setItem(key, JSON.stringify(value));
		},
		[value, key]
	);

	return [value, setValue];
}
```

使用示范

```jsx
function App() {
	const [watched, setWatched] = useLocalStorageState([], "watched");
}
```

### 案例 2：useKey

```jsx
import { useEffect } from "react";
export function useKey(key, action) {
	useEffect(
		function () {
			function callback(e) {
				// 比较的时候把字符串全转成小写或大写
				if (e.code.toLowerCase() === key.toLowerCase()) {
					action();
				}
			}

			document.addEventListener("keydown", callback);

			return function () {
				document.removeEventListener("keydown", callback);
			};
		},
		[action, key]
	);
}
```

改造示范,从这个案例我们可以看出，事件可以传入很复杂的内容

```jsx
//原1
function App() {
	const inputEl = useRef(null);
	useEffect(
		function () {
			function callback(e) {
				if (e.code === "Enter") {
					// 如果当前元素已激活就不再重复触发事件，该功能刷新页面后生效
					if (document.activeElement === inputEl.current) return;
					inputEl.current.focus();
					//清空之前的搜索内容，
					//1.我们跳转过来不想获得之前输入的内容
					//2.执行搜索后，应该清空本次输入
					setQuery("");
				}
			}
			// 更复杂的功能，希望按下回车键就可以进行聚焦搜索框
			document.addEventListener("keydown", callback);
			return () => {
				document.removeEventListener("keydown", callback);
			};
		},
		[setQuery]
	);
}
//改造后
function App() {
	const inputEl = useRef(null);
	useKey("Enter", function () {
		// 如果当前元素已激活就不再重复触发事件，该功能刷新页面后生效
		if (document.activeElement === inputEl.current) return;
		inputEl.current.focus();
		//清空之前的搜索内容，
		//1.我们跳转过来不想获得之前输入的内容
		//2.执行搜索后，应该清空本次输入
		setQuery("");
	});
}
```

改造案例 2

```jsx
//原2
function App() {
	useEffect(
		function () {
			function callback(e) {
				if (e.code === "Escape") {
					onCloseMovie();
				}
			}

			document.addEventListener("keydown", callback);

			return function () {
				document.removeEventListener("keydown", callback);
			};
		},
		[onCloseMovie]
	);
}
//改造2
function App() {
	useKey("Escape", onCloseMovie);
}
```
