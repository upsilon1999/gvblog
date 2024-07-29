## 组件实例的生命周期

### 1.组件挂载阶段

组件实例的生命周期第一阶段是组件挂载阶段(Mount),也被称为初始渲染(Initial render)，在这个阶段

```sh
1.组件第一次渲染
2.state和props被创建
```

这个阶段可以称为组件实例的诞生阶段

### 2.组件实例更新阶段

组件实例更新阶段，也可以叫做重复渲染(Re-render)阶段，在这个阶段，组件实例可以无限次地重新渲染，重新渲染的发生时机:

```sh
1.State更新
2.Props更新
3.父组件重新渲染
4.Context改变
```

这个阶段不是组件必须的，它是可选的，因为有些组件只有安装和卸载，不与 state、props 等接触

### 3.组件实例卸载阶段

在这一步中，将发生

```sh
1.组件实例将被完全销毁并从屏幕中删除
2.组件的state和props也将被销毁
```

## useEffect

### 前言

在接触这个钩子前我们先来看看直接在组件中发起请求会有什么副作用，例如

```jsx
function App() {
	fetch(`http://www.omdbapi.com/?apikey=${KEY}&s=interstellar`)
		.then((res) => res.json)
		.then((data) => console.log(data.Search));
}
```

这样的后果就是，每次 App 组件重新渲染都会发送请求，例如第一次挂载和每次状态更新,而如果我们在里面有状态更新的操作，则会陷入死循环，例如

```jsx
function App() {
	const [movies, setMovies] = useState([]);
	// 每次状态更新会发送请求，而发送请求成功后又更新了状态，从而陷入了死循环
	//这也就是我们常说的副作用
	fetch(`http://www.omdbapi.com/?apikey=${KEY}&s=interstellar`)
		.then((res) => res.json)
		.then((data) => setMovies(data.Search));
}
```

### useEffect 钩子的基本使用

引入钩子

```jsx
import useEffect from "react";
```

useEffect 的基本格式

```jsx
useEffect(效应函数, 依赖数组);
```

效应函数一般是个匿名函数或箭头函数，我们在效应函数中书写副作用，依赖数组则负责控制效应函数的执行时机

依赖数组为`[]`表示只在挂载的时候运行

```jsx
function App() {
	useEffect(function () {
		fetch(`http://www.omdbapi.com/?apikey=${KEY}&s=interstellar`)
			.then((res) => res.json)
			.then((data) => console.log(data.Search));
	}, []);
}
```

### useEffect 的详细解释

#### 什么是副作用

副作用是 React 组件与该组件之外的世界之间的任何相互作用，例如

```sh
从某些API获取数据，这时组件从外部世界获取数据，自然是副作用
```

副作用不应该出现在渲染逻辑中，在 React 中只有两个地方允许产生辅作用，

```sh
1.第一个是内部事件处理程序，我们应该时刻牢记，事件处理程序只是被触发的函数，即回调函数
2.useEffect，因为有时我们需要某项功能在组件挂载、组件销毁或者状态更新时触发，即在组件实例生命周期的不同时刻运行
```

#### 事件处理程序和 useEffect

例如获取电影数据，如果采用事件形式

```jsx
function App() {
	const [movies, setMovies] = useState([]);
	function handleClick() {
		fetch(`http://www.omdbapi.com/?apikey=${KEY}&s=interstellar`)
			.then((res) => res.json)
			.then((data) => setMovies(data.Search));
	}
	return <button onClick={handleClick}></button>;
}
```

如果我们想在组件实例的生命周期中获取，例如组件挂载后获取

```jsx
function App() {
	useEffect(function () {
		fetch(`http://www.omdbapi.com/?apikey=${KEY}&s=interstellar`)
			.then((res) => res.json)
			.then((data) => setMovies(data.Search));
	}, []);
}
```

useEffect 存在的真正原因，实际上不仅仅是生命周期，更重要的是保持组件与外部系统的同步
**使用原则**
内部事件处理程序是我们处理副作用的首选，基本上所有可以在事件处理程序中处理的事情都应该在那里处理，不要滥用 useEffect

#### 用 async 和 await 改写我们的请求副作用

原来的请求书写，

```jsx
function App() {
	useEffect(function () {
		fetch(`http://www.omdbapi.com/?apikey=${KEY}&s=interstellar`)
			.then((res) => res.json)
			.then((data) => setMovies(data.Search));
	}, []);
}
```

用 async 改写

```jsx
function App() {
	const [movies, setMovies] = useState([]);
	useEffect(function () {
		//构造异步请求函数
		async function fetchMovies() {
			const res = await fetch(
				`http://www.omdbapi.com/?apikey=${KEY}&s=interstellar`
			);
			const data = await res.json();
			setMovies(data.Search);
			// 一个小知识点，由于异步的原因,我们可以在这拿到旧的值
			//此时获取到的movies就是旧值
			console.log(movies);
		}
		//调用函数
		fetchMovies();
	}, []);
}
```

### 依赖数组

#### 组件生命时间线

```sh
【mount】
组件实例挂载或初始化阶段
【commit】
提交到DOM
【Brower Paint】
浏览器进行绘制
【effect】
浏览器绘制完后，effect才起作用，所以说effect是异步的，在渲染已经绘制到屏幕后才起作用，究其原因是effect可能包含长时间运行的进程，例如获取数据，在那种情况下如果React将在浏览器绘制新屏幕之前执行该效果,就会造成阻塞。但是effect异步呈现的后果就是，如果effect设置state，就需要第二个额外的渲染器来正确地显示UI，这也是为什么不要滥用effect的原因之一
【状态改变】
state或props改变
【re-render】
重新渲染，提交DOM，浏览器进行重绘
【effect】
如果该状态在依赖数组中，此次重绘将触发对应的effect
【unmount】
卸载组件
```

#### 依赖数组的三种形态

依赖数组时 useEffect 的第二个参数，有三种形态:

```jsx
// 不使用依赖数组，组件实例挂载以及每次重新更新组件都会触发效应函数
useEffect(fn);

//空数组依赖，组件实例挂载会触发效应函数
useEffect(fn, []);

//依赖数组内储存的是state和props，当挂载以及这些state或props更新时会触发效应函数
useEffect(fn, [x, y, z]);
```

### 多个 useEffect 的执行顺序

需要明确以下几点:

```sh
1.useEffect是异步的
2.同等条件下，多个useEffect按顺序执行
```

举例:

```jsx
function App() {
	useEffect(function () {
		console.log("A");
	}, []);
	useEffect(function () {
		console.log("B");
	});
	console.log("C");
}
```

初次挂载打印结果为

```sh
# 同步的先打印
C
# 按照顺序结构执行
A
B
```

每次状态更新时，C、B 都会打印，A 不会，因为 A 的依赖数组使其只在挂载时生效

### useEffect 清理函数

中文常译为清除副作用，英文名为 useEffect cleanup function，他是 useEffect 的第一个参数效应函数的返回值，该值要为函数，它是可选的，在 effect 再次执行前和组件卸载时执行，来补充生命周期图

```sh
【mount】
组件实例挂载或初始化阶段
【commit】
提交到DOM
【Brower Paint】
浏览器进行绘制
【effect】
浏览器绘制完后，effect才起作用，所以说effect是异步的，在渲染已经绘制到屏幕后才起作用，究其原因是effect可能包含长时间运行的进程，例如获取数据，在那种情况下如果React将在浏览器绘制新屏幕之前执行该效果,就会造成阻塞。但是effect异步呈现的后果就是，如果effect设置state，就需要第二个额外的渲染器来正确地显示UI，这也是为什么不要滥用effect的原因之一
【状态改变】
state或props改变
【cleanUP】
清除上一次的effect
【re-render】
重新渲染，提交DOM，浏览器进行重绘
【effect】
如果该状态在依赖数组中，此次重绘将触发对应的effect
【cleanUP】
执行效应函数的返回函数
【unmount】
卸载组件
```

**cleanUP 函数特点**

```sh
1.该函数时useEffect的第一个参数效应函数的返回值，是可选的
2.他在两个时机执行:
a.在效果再次执行之前运行,为了清理之前副作用的结果
b.在组件实例卸载后立即执行，是为了给我们机会重置我们制造的副作用
```

实际上每个 useEffect 都应该有一个 cleanUP 返回，以便在重新呈现或卸载组件后清理副作用，例如

```sh
如果我们有一个发起http请求的副作用，如果在重新渲染后，第一个请求仍在运行，但第二个请求也会发出，就可能产生一个叫做争用条件的bug，因此在清理函数中取消请求是个好主意
```

**清理函数常做的事情**

```sh
1.针对http请求，取消请求
2.订阅API后，取消API
3.开启计时器后，关闭计时器
4.添加事件侦听器（event listener）后，移除事件监听(listener)
```

结合一个实际案例来理解，我们希望在电影详情组件卸载后还原电影标题，所以

```jsx
function MovieDetails(){
	...
	//最好每个useEffect只实现一个副作用
	useEffect(function(){
		// 我们可以通过document拿到html页面元素
		console.log(document)
		// 当没有标题时不修改标题，使用默认标题
		if(!title) return
		// 通过依赖监听来即时改变标题
		document.title = `movie|${title}`
		// 清理函数，在重新effect和组件实例卸载后执行
		return function(){
			document.title = "usePopcorn"
		}
	},[title])
}
```

#### 闭包思维

看下一段代码

```jsx
function MovieDetails(){
	...
	//最好每个useEffect只实现一个副作用
	useEffect(function(){
		// 我们可以通过document拿到html页面元素
		console.log(document)
		// 当没有标题时不修改标题，使用默认标题
		if(!title) return
		// 通过依赖监听来即时改变标题
		document.title = `movie|${title}`
		// 清理函数，在重新effect和组件实例卸载后执行
		return function(){
			document.title = "usePopcorn"
			console.log("清理函数",title)
		}
	},[title])
}
```

我们发现清理函数中 title 会被打印，它是最后一次的 title 值，这时可能会产生一个疑问，清理函数是卸载后执行的，此时组件的状态、porps 等都被销毁了，这个 title 是怎么读取到的，这其实是 js 闭包的特性，js 的闭包意味着函数将永远记住在函数创建的事件和地点存在的所有变量，所以在清理函数中可以获得上一次渲染后或销毁前的数据，

#### 清理 http 请求

如果不清理 http 请求会存在争用的问题，上一个请求还在运行时，下一个请求已经发送，当网速慢时，可能会积累数个请求，最终会下载过多的数据。还有一种情况，假设我们同时发起了 6 个请求，实际我们想要的就是第六个，但是第 3 个最后下载完，那么我们最终获取到的实际上是第 3 个的数据，这就是争用导致的错误。
所以我们应该做到

```sh
当新的相同请求发出时，上一个请求应该停止，即取消
```

要实现这个效果，我们需要创造一个终止控制器，按照如下步骤

```jsx
function App() {
	//1.创造一个控制器的变量
	// AbortController这是浏览器的API，专门用来终止请求
	const controller = new AbortController();
	useEffect(
		function () {
			async function fetchMovies() {
				try {
					setIsLoading(true);
					setError("");
					//2.标记此次请求，axios的写法可能和fetch不一样，自行了解
					//在fetch请求中添加第二个参数，{ sigal: controller.sigal }
					const res = await fetch(
						`http://www.omdbapi.com/?apikey=${KEY}&s=${query}`,
						{ sigal: controller.sigal }
					);
					if (!res.ok) {
						throw new Error("someting went wrong!");
					}

					const data = await res.json();
					if (data.Response === "False") {
						throw new Error("Movie not found");
					}
					setMovies(data.Search);
				} catch (err) {
					setError(err.message);
				} finally {
					setIsLoading(false);
				}
			}
			if (!query.length) {
				setMovies([]);
				setError("");
				return;
			}
			fetchMovies();
			//3.在重新渲染或组件卸载后终止该次标记的请求
			// 在清理函数中调用控制器的abort()方法
			return function () {
				controller.abort();
			};
		},
		[query]
	);
}
```

**补充**
当我们使用 try...catch 时 abort 会被判断成错误，我们希望不报错，所以处理如下

```jsx
function App() {
	const controller = new AbortController();

	useEffect(
		function () {
			async function fetchMovies() {
				try {
					setIsLoading(true);
					setError("");

					const res = await fetch(
						`http://www.omdbapi.com/?apikey=${KEY}&s=${query}`,
						{ signal: controller.signal }
					);
					if (!res.ok) {
						throw new Error("someting went wrong!");
					}

					const data = await res.json();
					if (data.Response === "False") {
						throw new Error("Movie not found");
					}
					setMovies(data.Search);
					setError("");
				} catch (err) {
					// abort会被捕获，因为我们主动终止了请求，所以就会触发一个AbortError
					//而我们知道这不是我们想抛出的错误，所以可以不抓取他
					if (err.name !== "AbortError") {
						setError(err.message);
					}
				} finally {
					setIsLoading(false);
				}
			}
			if (!query.length) {
				setMovies([]);
				setError("");
				return;
			}
			fetchMovies();

			return function () {
				constroller.abort();
			};
		},
		[query]
	);
}
```

### useEffect 的重要规则

每个 useEffect 应该只有一个 effect，如果想在组件中有多个 effect，那么就使用多个 useEffect，这样不仅每个效果容易理解，清理功能也更容易清理

### 对 keypress 的监听

需求:在全局范围内监听按键事件，这显然是一个副作用，所以

```jsx
function App(){
	...
	useEffect(function(){
		// 因为是全局监听，所以给document增加监听事件
		document.addEventListener("keydown",function(e){
			if(e.code=== "Escape"){
				// 和返回按钮会触发的事件一样，就是跳转
				handleCloseMovie()
			}
		})
	},[])
}
```

如果我们不加入清理事件，那么事件监听器将无限累加，这将造成严重的内存问题。为了移除事件监听器我们必须使用具名函数，所以改造为

```jsx
function App() {
	useEffect(function () {
		function callback(e) {
			if (e.code === "Escape") {
				// 和返回按钮会触发的事件一样，就是跳转
				handleCloseMovie();
			}
		}
		// 因为是全局监听，所以给document增加监听事件
		// 为了能够移除监听事件，必须使用具名函数来注册
		document.addEventListener("keydown", callback);

		return function () {
			// 移除注册的回调事件
			document.removeEventListener("keydown", callback);
		};
	}, []);
}
```

## 测试网速较慢的情况

在很多时候我们想重现网速慢导致页面加载不出来的情况,chrome 可以使用

```sh
控制面板(F12)-->网络(network)-->已停用节流模式(点击可以展开下拉列表，修改预设为3G)
```

当页面加载缓慢时就会有一段空白周期，一般都会加入一个 loading 状态，例如

```jsx
function App() {
	const [isLoading, setIsLoading] = useState(false);
	const [movies, setMovies] = useState([]);
	useEffect(function () {
		async function fetchMovies() {
			// 常常在异步请求发送前将loading启用，请求成功或失败后关闭
			setIsLoading(true);
			const res = await fetch(
				`http://www.omdbapi.com/?apikey=${KEY}&s=interstellar`
			);
			const data = await res.json();
			setMovies(data.Search);
			// 在请求的结束关闭loading
			setIsLoading(false);
		}
		fetchMovies();
	}, []);

	// 然后在需要动态展示的组件处使用loading
	return <>{isLoading ? <Loading /> : <MovieList />}</>;
}
```

## 错误捕获

当我们做异步处理时都应该要假设会出现错误

```jsx
function ErrorMessage({ message }) {
	return (
		<p className="error">
			<span>⛔</span>
			{message}
		</p>
	);
}
function App() {
	// 此状态控制加载
	const [isLoading, setIsLoading] = useState(false);
	//此状态控制错误
	const [error, setError] = useState("");
	const [movies, setMovies] = useState([]);
	useEffect(function () {
		async function fetchMovies() {
			// 捕获错误
			try {
				setIsLoading(true);
				const res = await fetch(
					`http://www.omdbapi.com/?apikey=${KEY}&s=interstellar`
				);
				// 在请求出错时抛出错误
				if (!res.ok) {
					throw new Error("someting went wrong!");
				}

				const data = await res.json();
				// 当查询的电影列表不存在时
				if (data.Response === "False") {
					throw new Error("Movie not found");
				}
				setMovies(data.Search);
				// 在请求的结束关闭loading
				// 由于throw后不再往下执行，所以他不会走到这
				// setIsLoading(false);
			} catch (err) {
				// 这里抓到的就是try中抛出的错误对象
				console.log(err.message);
				setError(err.message);
			} finally {
				// 无论请求成功还是失败都应该关闭loading，所以写在try...finally里面
				setIsLoading(false);
			}
		}
		fetchMovies();
	}, []);

	return (
		<>
			{/*正在加载  */}
			{isLoading && <Loading />}
			{/* 没有加载但是没报错 */}
			{!isLoading && !error && <MovieList />}
			{/* 报错了 */}
			{error && <ErrorMessage message={error} />}
		</>
	);
}
```

当然呈现的时候也可以使用三元来判断，这里分开写只是为了提高可读性

## 控制浏览器的网页标题

我们都知道对于 React 程序来说，浏览器的网页标题默认由`public/index.html`来控制，例如

```html
<title>React App</title>
```

现在我们想实现通过点击电影列表，让浏览器的网页标题变为电影列表点击项的电影名，在应用程序的浏览器中更改页面标题是一个副作用，因为这显然是与外界进行交互.

```jsx
function MovieDetails(){
	...
	//最好每个useEffect只实现一个副作用
	useEffect(function(){
		// 我们可以通过document拿到html页面元素
		console.log(document)
		// 当没有标题时不修改标题，使用默认标题
		if(!title) return
		// 通过依赖监听来即时改变标题
		document.title = `movie|${title}`
	},[title])
}
```

### 使用清理函数

当我们点击返回按钮时，应该重置标题，实际上当我们点击返回按钮时，右边组件已经卸载，所以需求就是在卸载时重置标题

```jsx
function MovieDetails(){
	...
	//最好每个useEffect只实现一个副作用
	useEffect(function(){
		// 我们可以通过document拿到html页面元素
		console.log(document)
		// 当没有标题时不修改标题，使用默认标题
		if(!title) return
		// 通过依赖监听来即时改变标题
		document.title = `movie|${title}`
		return function(){
			document.title = "usePopcorn"
		}
	},[title])
}
```

课后任务，根据 149-157 把代码补充完整，以及构造一个电影 API 后台，只要能提供基本数据和电影详情即可。
