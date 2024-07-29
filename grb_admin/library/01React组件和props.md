# 组件

## 技巧

在 windows 的 vscode 中，在函数名上用 ctrl 就会跳到函数声明的地方
在声明或定义函数的地方用 ctrl 点击函数名会显示出他在何处被使用

### 节约渲染成本

将不会重新渲染和永远不会改变的变量再组件外部声明和定义，相当于全局变量，举例

```jsx
// 放在组件函数外，防止每次组件重新渲染再生成他们
const containerStyle = {
	display: "flex",
	alignItems: "center",
	gap: "16px",
};

export default function StarRating() {
	return (
		<div style={containerStyle}>
			<div>
				{Array.from({ length: 5 }, (_, i) => (
					<span>S{i + 1}</span>
				))}
			</div>
			<p>10</p>
		</div>
	);
}
```

### 渲染列表的方式

之前我们渲染列表多使用 map 方法，实际上还可以使用数组的各种方式，例如 Array.form,举例

```jsx
{
	Array.from({ length: maxRating }, (_, i) => <Star key={i}></Star>);
}
```

## 了解组件

```sh
1.如何将用户界面拆分成组件
2.什么时候真正创建新组件
```

学会区分组件，两个极端巨大组件和超小组件，拿巨大组件来说：

```sh
1.React组件本身是一个函数，里面包含了太多的功能
2.接收或者需要太多的props，例如10个或20个
3.难以重用
4.包含复杂切交织在一起的代码，使得整个组件难以理解和使用
```

然后考虑超小组件，

```sh
1.页面被拆分成了数百个甚至数千个迷你组件
2.混乱的代码库，因为组件太多了
3.过于抽象，很多组件没有实际意义，本该知识一部分的功能却被抽离成了组件
```

所以合理的方式应该是折中，每一个组件都有较为明确的含义，也不会过于复杂，基本思路:

```sh
1.内容的逻辑分离
2.可重用
3.明确含义
```

思考一下即可，随着编码的增多养成自己的编码风格，这些问题会迎刃而解。

## 成分组成(插槽)

实际上就是利用 children props 来复用组件，来看下面一个例子

```jsx
function Modal() {
	return (
		<div className="modal">
			<Success />
		</div>
	);
}
function Success() {
	return <p>Well Done!</p>;
}
```

实际上 Modal 组件只能用来展示成功的信息，不能用来展示失败或警告的信息，复用度很低，所以我们可以用 children props 进行改造

```jsx
function Modal({ children }) {
	return <div className="modal">{children}</div>;
}
function Success() {
	return <p>Well Done!</p>;
}
```

这样我们就可以在使用的时候根据标签体的组件来复用，例如

```jsx
<Modal>
	<Success />
</Modal>
<Modal>
	<Error />
</Modal>
```

## 用插槽来解决螺旋钻孔问题

什么是螺旋钻孔问题，实际上就形如爷孙组件通信,

```sh
爷爷把数据传递给父亲，父亲再转发给儿子，实现爷爷与孙子的通信
```

实际案例

```jsx
export default function App() {
	const [movies, setMovies] = useState(tempMovieData);
	return (
		<>
			<NavBar movies={movies} />
		</>
	);
}

function NavBar({ movies }) {
	return (
		<nav className="nav-bar">
			<Logo />
			<Search />
			<NumResults movies={movies} />
		</nav>
	);
}

function NumResults({ movies }) {
	return (
		<p className="num-results">
			Found <strong>{movies.length}</strong> results
		</p>
	);
}
```

movies 从 App 到 NavBar，再到真正需要他的 NumResults,可以理解为爷爷见不到孙子，所以需要爸爸来转发，而运用插槽的写法就相当于爸爸把儿子直接带过来了，爷爷可以直接见到孙子，举例:

```jsx
export default function App() {
	const [movies, setMovies] = useState(tempMovieData);
	return (
		<>
			<NavBar>
				<Logo />
				<Search />
				<NumResults movies={movies} />
			</NavBar>
		</>
	);
}

function NavBar({ children }) {
	return <nav className="nav-bar">{children}</nav>;
}

function NumResults({ movies }) {
	return (
		<p className="num-results">
			Found <strong>{movies.length}</strong> results
		</p>
	);
}
```

## 具名插槽

所谓的具名插槽实际上就是借用 react 组件属性能传递任何内容的特性。
children props 本质上是组件上的 children 属性，解析

```jsx
function Box({ children }) {
	const [isOpen, setIsOpen] = useState(true);
	return (
		<div className="box">
			<button
				className="btn-toggle"
				onClick={() => setIsOpen((open) => !open)}
			>
				{isOpen ? "–" : "+"}
			</button>
			{isOpen && children}
		</div>
	);
}

export default function App() {
	const [movies, setMovies] = useState(tempMovieData);
	const [watched, setWatched] = useState(tempWatchedData);
	return (
		<>
			<Main>
				<Box>
					<MovieList movies={movies} />
				</Box>
			</Main>
		</>
	);
}
//等价于
export default function App() {
	const [movies, setMovies] = useState(tempMovieData);
	const [watched, setWatched] = useState(tempWatchedData);
	return (
		<>
			<Main>
				<Box children={<MovieList movies={movies}/>}/>
			</Main>
		</>
	);
}
```

所以我们完全可以用自定义属性来传递 jsx 内容,例如

```jsx
function Box({ element }) {
	const [isOpen, setIsOpen] = useState(true);
	return (
		<div className="box">
			<button
				className="btn-toggle"
				onClick={() => setIsOpen((open) => !open)}
			>
				{isOpen ? "–" : "+"}
			</button>
			{isOpen && element}
		</div>
	);
}

export default function App() {
	const [movies, setMovies] = useState(tempMovieData);
	const [watched, setWatched] = useState(tempWatchedData);
	return (
		<>
			<Main>
				<Box element={<MovieList movies={movies} />} />
				{/* 但传入的jsx很长时可以使用 <></>包裹*/}
				<Box
					element={
						<>
							<WatchedSummary watched={watched} />
							<WatchedList watched={watched} />
						</>
					}
				/>
			</Main>
		</>
	);
}
```

之所以说具名插槽，是因为我们完全可以传入多个不同的 jsx，给他们不同的属性名。

## 星星评级组件

几个要点：

```sh
1.通过Array.from遍历生成n个相同的元素
2.svg要给定宽高才会展示，所以一般用span包裹
3.实现评分和星级的绑定关系，一般采用索引，索引加1就是评分
4.空星星和满星星，采用当前分数和索引的比较来动态渲染
```

实现基本的评级组件

```jsx
import { useState } from "react";

// 放在组件函数外，防止每次组件重新渲染再生成他们
const containerStyle = {
	display: "flex",
	alignItems: "center",
	gap: "16px",
};
const starContainerStyle = {
	display: "flex",
};

const textStyle = {
	lineHeight: "1",
	margin: "0",
};
// 解构的时候设置默认值
export default function StarRating({ maxRating = 5 }) {
	const [rating, setRating] = useState(0);
	function handleRating(rating) {
		setRating(rating);
	}

	// 星星评级实现的思路，点击那个星星就用该星星的索引加1去更新分数
	return (
		<div style={containerStyle}>
			<div style={starContainerStyle}>
				{Array.from({ length: maxRating }, (_, i) => (
					<Star
						key={i}
						onRate={() => handleRating(i + 1)}
						full={rating >= i + 1}
					></Star>
				))}
			</div>
			<p style={textStyle}>{rating || ""}</p>
		</div>
	);
}

const starStyle = {
	width: "48px",
	height: "48px",
	display: "block",
	cursor: "pointer",
};
function Star({ onRate, full }) {
	return (
		// svg图片必须给定宽高才会展示
		<span role="button" style={starStyle} onClick={onRate}>
			{/* 通过full来判断展示空星还是满星 */}
			{full ? (
				<svg
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 20 20"
					fill="#000"
					stroke="#000"
				>
					<path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
				</svg>
			) : (
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke="#000"
				>
					<path
						strokeLinecap="round"
						strokeLinejoin="round"
						strokeWidth="{2}"
						d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"
					/>
				</svg>
			)}
		</span>
	);
}
```

### 附加功能

我们希望鼠标在星条上滑动时会临时评分，例如滑动 10，就显示 10 星和 10 分，所以需要创建一个状态来存储临时值

```sh
鼠标的滑动悬停事件可以用onMouseEnter和onMouseLeave来实现
```

实际实现

```jsx
import { useState } from "react";

// 放在组件函数外，防止每次组件重新渲染再生成他们
const containerStyle = {
	display: "flex",
	alignItems: "center",
	gap: "16px",
};
const starContainerStyle = {
	display: "flex",
};

const textStyle = {
	lineHeight: "1",
	margin: "0",
};
// 解构的时候设置默认值
export default function StarRating({ maxRating = 5 }) {
	const [rating, setRating] = useState(0);
	const [tempRating, setTempRating] = useState(0);
	function handleRating(rating) {
		setRating(rating);
	}

	// 星星评级实现的思路，点击那个星星就用该星星的索引加1去更新分数
	return (
		<div style={containerStyle}>
			{/* 鼠标离开时将临时分数置为0的原因，那样就不会干扰到实际的评分 */}
			<div style={starContainerStyle}>
				{Array.from({ length: maxRating }, (_, i) => (
					<Star
						key={i}
						onRate={() => handleRating(i + 1)}
						full={
							tempRating ? tempRating >= i + 1 : rating >= i + 1
						}
						onHoverIn={() => setTempRating(i + 1)}
						onHoverout={() => setTempRating(0)}
					></Star>
				))}
			</div>
			{/* 鼠标离开后临时分数变为0，所以展示实际分数 */}
			<p style={textStyle}>{tempRating || rating || ""}</p>
		</div>
	);
}

const starStyle = {
	width: "48px",
	height: "48px",
	display: "block",
	cursor: "pointer",
};
function Star({ onRate, full, onHoverIn, onHoverout }) {
	return (
		// svg图片必须给定宽高才会展示
		<span
			role="button"
			style={starStyle}
			onClick={onRate}
			onMouseEnter={onHoverIn}
			onMouseLeave={onHoverout}
		>
			{/* 通过full来判断展示空星还是满星 */}
			{full ? (
				<svg
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 20 20"
					fill="#000"
					stroke="#000"
				>
					<path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
				</svg>
			) : (
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke="#000"
				>
					<path
						strokeLinecap="round"
						strokeLinejoin="round"
						strokeWidth="{2}"
						d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"
					/>
				</svg>
			)}
		</span>
	);
}
```

### 将星星组件变成一个可自定义的组件

例如复用该组件，根据传入的 props 来来定义星星颜色和大小等等，注意默认值的设置以及 svg 的样式

```jsx
import { useState } from "react";

// 放在组件函数外，防止每次组件重新渲染再生成他们
const containerStyle = {
	display: "flex",
	alignItems: "center",
	gap: "16px",
};
const starContainerStyle = {
	display: "flex",
};

// 解构的时候设置默认值
export default function StarRating({
	maxRating = 5,
	color = "#fcc419",
	size = 48,
}) {
	const [rating, setRating] = useState(0);
	const [tempRating, setTempRating] = useState(0);
	function handleRating(rating) {
		setRating(rating);
	}

	const textStyle = {
		lineHeight: "1",
		margin: "0",
		color,
		fontSize: `${size / 1.5}px`,
	};

	// 星星评级实现的思路，点击那个星星就用该星星的索引加1去更新分数
	return (
		<div style={containerStyle}>
			{/* 鼠标离开时将临时分数置为0的原因，那样就不会干扰到实际的评分 */}
			<div style={starContainerStyle}>
				{Array.from({ length: maxRating }, (_, i) => (
					<Star
						key={i}
						onRate={() => handleRating(i + 1)}
						full={
							tempRating ? tempRating >= i + 1 : rating >= i + 1
						}
						onHoverIn={() => setTempRating(i + 1)}
						onHoverout={() => setTempRating(0)}
						color={color}
						size={size}
					></Star>
				))}
			</div>
			{/* 鼠标离开后临时分数变为0，所以展示实际分数 */}
			<p style={textStyle}>{tempRating || rating || ""}</p>
		</div>
	);
}

function Star({ onRate, full, onHoverIn, onHoverout, color, size }) {
	const starStyle = {
		width: `${size}px`,
		height: `${size}px`,
		display: "block",
		cursor: "pointer",
	};
	return (
		// svg图片必须给定宽高才会展示
		<span
			role="button"
			style={starStyle}
			onClick={onRate}
			onMouseEnter={onHoverIn}
			onMouseLeave={onHoverout}
		>
			{/* 通过full来判断展示空星还是满星 */}
			{/* svg的颜色是在svg的属性上定义的，fill是填充色，stroke是边框色 */}
			{full ? (
				<svg
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 20 20"
					fill={color}
					stroke={color}
				>
					<path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
				</svg>
			) : (
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke={color}
				>
					<path
						strokeLinecap="round"
						strokeLinejoin="round"
						strokeWidth="{2}"
						d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"
					/>
				</svg>
			)}
		</span>
	);
}
```

有时候可能会让用户传入自定义类名来修改某些样式，就将接收的类名字符串渲染到指定位置，例如,一般来说默认值设定为空字符串

```jsx
import { useState } from "react";

// 放在组件函数外，防止每次组件重新渲染再生成他们
const containerStyle = {
	display: "flex",
	alignItems: "center",
	gap: "16px",
};
const starContainerStyle = {
	display: "flex",
};

// 解构的时候设置默认值
export default function StarRating({
	maxRating = 5,
	color = "#fcc419",
	size = 48,
	className = "",
}) {
	const [rating, setRating] = useState(0);
	const [tempRating, setTempRating] = useState(0);
	function handleRating(rating) {
		setRating(rating);
	}

	const textStyle = {
		lineHeight: "1",
		margin: "0",
		color,
		fontSize: `${size / 1.5}px`,
	};
	return (
		<div style={containerStyle} className={className}>
			<div style={starContainerStyle}>
				{Array.from({ length: maxRating }, (_, i) => (
					<Star
						key={i}
						onRate={() => handleRating(i + 1)}
						full={
							tempRating ? tempRating >= i + 1 : rating >= i + 1
						}
						onHoverIn={() => setTempRating(i + 1)}
						onHoverout={() => setTempRating(0)}
						color={color}
						size={size}
					></Star>
				))}
			</div>
			<p style={textStyle}>{tempRating || rating || ""}</p>
		</div>
	);
}
```

有时候需求显示的不是评分而是根据评分来展示不同评价，一般会传入一个数组，来通过索引展示不同的值

```jsx
import { useState } from "react";

// 放在组件函数外，防止每次组件重新渲染再生成他们
const containerStyle = {
	display: "flex",
	alignItems: "center",
	gap: "16px",
};
const starContainerStyle = {
	display: "flex",
};

// 解构的时候设置默认值
export default function StarRating({
	maxRating = 5,
	color = "#fcc419",
	size = 48,
	className = "",
	messages = ["terrible", "Okay", "Good", "Son", "News"],
}) {
	const [rating, setRating] = useState(0);
	const [tempRating, setTempRating] = useState(0);
	function handleRating(rating) {
		setRating(rating);
	}

	const textStyle = {
		lineHeight: "1",
		margin: "0",
		color,
		fontSize: `${size / 1.5}px`,
	};

	// 星星评级实现的思路，点击那个星星就用该星星的索引加1去更新分数
	return (
		<div style={containerStyle} className={className}>
			{/* 鼠标离开时将临时分数置为0的原因，那样就不会干扰到实际的评分 */}
			<div style={starContainerStyle}>
				{Array.from({ length: maxRating }, (_, i) => (
					<Star
						key={i}
						onRate={() => handleRating(i + 1)}
						full={
							tempRating ? tempRating >= i + 1 : rating >= i + 1
						}
						onHoverIn={() => setTempRating(i + 1)}
						onHoverout={() => setTempRating(0)}
						color={color}
						size={size}
					></Star>
				))}
			</div>
			{/* 鼠标离开后临时分数变为0，所以展示实际分数 */}
			{/* <p style={textStyle}>{tempRating || rating || ""}</p> */}
			<p style={textStyle}>
				{messages.length === maxRating
					? messages[tempRating ? tempRating - 1 : rating - 1]
					: tempRating || rating || ""}
			</p>
		</div>
	);
}
```

注意事项，传入的数组长度应该和评分值由对应关系，例如数组长度为 5，评分值为 1~5，否则会由于无法对应而报错，一个惊艳的地方

```jsx
{
	messages.length === maxRating
		? messages[tempRating ? tempRating - 1 : rating - 1]
		: tempRating || rating || "";
}
```

数组内部索引也是可以用三元表达式的

### 让组件内部评分可以被外部读取

实际上的做法是在外部维护一个状态，将更改状态的方法传入星星组件，当然星星组件内部得接收这个方法，并且将自己的评分绑定给该方法，实际上就是父子组件通信，我们的组件应该预留通信的通道。

```jsx
import { useState } from "react";

// 放在组件函数外，防止每次组件重新渲染再生成他们
const containerStyle = {
	display: "flex",
	alignItems: "center",
	gap: "16px",
};
const starContainerStyle = {
	display: "flex",
};

// 解构的时候设置默认值
export default function StarRating({
	maxRating = 5,
	color = "#fcc419",
	size = 48,
	onSetRating,
}) {
	const [rating, setRating] = useState(0);
	const [tempRating, setTempRating] = useState(0);
	function handleRating(rating) {
		setRating(rating);
		// 向外部暴露评分
		onSetRating(rating);
	}

	const textStyle = {
		lineHeight: "1",
		margin: "0",
		color,
		fontSize: `${size / 1.5}px`,
	};

	// 星星评级实现的思路，点击那个星星就用该星星的索引加1去更新分数
	return (
		<div style={containerStyle}>
			{/* 鼠标离开时将临时分数置为0的原因，那样就不会干扰到实际的评分 */}
			<div style={starContainerStyle}>
				{Array.from({ length: maxRating }, (_, i) => (
					<Star
						key={i}
						onRate={() => handleRating(i + 1)}
						full={
							tempRating ? tempRating >= i + 1 : rating >= i + 1
						}
						onHoverIn={() => setTempRating(i + 1)}
						onHoverout={() => setTempRating(0)}
						color={color}
						size={size}
					></Star>
				))}
			</div>
			{/* 鼠标离开后临时分数变为0，所以展示实际分数 */}
			<p style={textStyle}>{tempRating || rating || ""}</p>
		</div>
	);
}
```

## porps

### props 设置默认值

直接在解构赋值时设置默认值即可

```jsx
function Son({ name = "Lisa" }) {
	return <span>{name}</span>;
}
```

### props 规定类型

值得指出的是，现在的 react 程序一般不再限定 props 的类型，因为如果在意类型，完全可以使用 TypeScript
**步骤 1**
导入 prop-types 包,现在这个包已经不需要额外安装,因为 create-react-app 预装了这个包

```jsx
import PropTypes from "prop-types";
```

**步骤 2**
使用的方式是给要进行 props 校验的组件增加一个 propTypes 属性，

```jsx
function Son({ name, age }) {
	return <span>{name}</span>;
}

//给Son增加一个propTypes属性
Son.propTypes = {
	name: PropTypes.string,
	age: PropTypes.number,
};
```

**注意事项**

```sh
1.引入的校验器一般推荐命名为首字母大写的PropTypes
2.给组件添加的propTypes属性的首字母是小写的
3.校验规则的形式是:【属性名:PropTypes.类型】，这里的类型可以是【number,string,array,bool,func,object】
```

### props 限制必要性

也需要借助于 PropTypes,语法为

```jsx
// 属性名:PropTypes.isrequired
Son.propTypes = {
	name: PropTypes.isrequired,
};
```

也可以同时限制类型和必要性，采用链式语法,例如

```jsx
// 属性名:PropTypes.类型.isrequired
Son.propTypes = {
	name: PropTypes.string.isrequired,
};
```
