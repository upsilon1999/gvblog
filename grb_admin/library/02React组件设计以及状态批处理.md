# React 高级

对应课程 121-

## 组件

我们可以通过打印来得到函数式组件的本质

```jsx
function App() {
	return <p>"hello"</p>;
}

// 可以从控制态看到，他被解析为一个react元素
console.log(<App />);
console.log(<App test={12} />);

// 直接调用，得到的不是react元素
// 虽然组件也会得到展示，但是React不再将他视为一个组件实例
// 后果就是无法再维护状态，以及很多hook无法使用
console.log(App());
```

### key 的作用

当有多个相同类型的组件使用时，推荐使用 key，这会提高效率，key 是一个默认属性，React 会识别他,将它当作一个新的组件实例，即使我们不展示它了他也没有被立即销毁，而是在 DOM 的某处保存，等待再需要它时被启用，例如

```jsx
<ul>
	<Tab key={1} />
	<Tab key={2} />
	<Tab key={3} />
</ul>
```

设计 key 的另一个好处,会记录自身基本状态，而不是沿用兄弟的，举例：

```sh
假如Tab中有一个折叠文本的功能,在不设置key的情况下,如果我们在Tab1中折叠了文本,那么Tab2会沿用这个折叠状态，虽然文本内容不同了(因为文本内容由state控制),但是控制折叠的DOM没变。
而加上key后，上一个组件的DOM就不会影响到下一个。
```

**使用场景**

```sh
1.列表中多个相同元素一定要用key
2.如果组件状态发生了改变但是DOM没变，或者组件状态只改变了一部分，推荐用key，(常见于组件设定了极其复杂的状态或者接收了复杂的props)
```

key 可以设置的很复杂，例如几百个字符的字符串都是可行的，只要它唯一即可。

## 组件设计规则

参考 134

## 状态批处理

这是 18 即之后会有的功能,即更新不是一个一个的而是一次性提交,例如

```jsx
// 在空白处打印,用于检查渲染次数
console.log("RENDER");
// 三个更新统一处理，只进行一次提交，一次渲染，大大提高了渲染效率
function change() {
	setCount(1);
	setStr("");
	setAnswer(true);
}
```

直接更新是异步的，所以无法同步操作只能拿到旧的值，因为 react 还没更新

```jsx
/*
	旧的值：
	count 5
*/
function change() {
	setCount(1);
	// 虽然我们更新了state状态，但是由于是异步的，所以这里获取的还是旧的值
	console.log(count); //5
	// 这也是我们能使用	setCount(count+1) 这种形式的原因
	setStr("");
	setAnswer(true);
}
```

原因是状态要 React 重新渲染后才更改，是异步的，所以这样的同步操作只能拿到旧值。这也解释了如下操作为什么不是加 2

```jsx
function change() {
	// 在这种情况下只执行最后一个
	setCount(count + 1);
	setCount(count + 5);
	setCount(count + 2);
}
```

按照我们的预想，应该执行三次，可实际并不是,他实际上只执行了 setCount(count + 2)，即只执行了最后一个,下面列出各种情况

```jsx
// 只执行最后一个
function change() {
	setCount(count + 1);
	setCount(count + 5);
	setCount(count + 2);
}
//只执行最后一个
function change() {
	setCount((count) => count + 1);
	setCount(count + 5);
	setCount(count + 2);
}
// 只执行最后一个
function change() {
	setCount(count + 1);
	setCount((count) => count + 5);
	setCount(count + 2);
}
// 只执行最后一个
function change() {
	setCount((count) => count + 1);
	setCount((count) => count + 5);
	setCount(count + 2);
}
// 三个都执行
function change() {
	setCount(count + 1);
	setCount(count + 5);
	setCount((count) => count + 2);
}
```

> 总结:如果最后一个是函数形式则执行所有，否则都只执行最后一个

所以如果我们要想实现多次更新推荐全使用函数形式

```jsx
function change() {
	setCount((count) => count + 1);
	setCount((count) => count + 5);
	setCount((count) => count + 2);
}
```

**深度解析**

```jsx
setCount((当前值) => 操作);
```

函数式里面可以获得当前值，且会从上一次继承，

```jsx
setCount((当前值) => {
	return 当前值 + 1;
});
setCount((上一次操作的结果) => 操作);
```

回到我们刚才的案例,来一一解读，

```jsx
function handleDo() {
	//当前count值是0
	setCount((count) => {
		// 打印的是当前值0
		console.log(count); //0
		return count + 1;
	});
	// 同步获取得到的是当前值0,
	console.log(count); //0

	// 此刻传入的状态值是上一次setCount的结果，即1
	setCount((count) => {
		// 所以打印得到1
		console.log(count); //1
		return count + 2;
	});

	// 按顺序执行到这，count是当前值0
	// 所以最终结果是5
	setCount(count + 5);
}
```

再看一个案例

```jsx
// 当前值count 3
function handleDo() {
	setCount(count + 3);
	// 同步获取结果为0
	console.log(count);

	//传入的count实际上是上一次setCount的结果3
	setCount((count) => {
		console.log(count); //3
		return count + 6;
	});

	// 按照顺序执行到这里，count使用的是当前值0
	//所以最终结果是5
	setCount(count + 5);
}
```

**小结**
在回调函数中我们能获得上一次操作的结果，所以如果想要实现多层次状态更新应该注意:

```sh
1.推荐全部使用回调函数
主要是由于只有回调才能接收上一次修改的值
2.其次注意执行顺序
```

**拓展 1**
关于为什么要使用回调函数？
来看一个实际案例，我们将更新状态封装到了一个函数，例如

```jsx
function handleAdd() {
	setCount(count + 1);
}
```

我们在调用时想增加 2，就可能写成以下这种情况

```jsx
function addTwo() {
	handleAdd();
	handleAdd();
}
```

根据前面的知识我们可以知道，这最终只会加 1，所以尽量在改变状态时使用回调函数形式。这样的话方便方法的复用

```jsx
function handleAdd() {
	setCount((count) => count + 1);
}
```

**拓展 2**
如果状态没有变化，是不会触发重新渲染的，举例

```jsx
/*
	当前值：
	count 1
	answer true
*/
function change() {
	setCount(1);
	setAnswer(true);
}
```

因为执行 change 没有更新状态值，所以不会触发渲染，即该函数会正常执行，但是组件实例不会被影响

## React 事件

p137
了解了 DOM 的原理，以及 React 事件的底层实现。React 的事件是合成事件。
**一些基础了解**

```sh
1.React中将原生事件进行了封装，React提供的事件都是驼峰形式，例如原生为onclick，React为onClick
2.在React中阻止事件默认行为的唯一方法是"事件对象.preventDefault()"
3.一个特殊的捕获事件，onClickCapture
```

## 组件不要嵌套

为什么不要在组件 A 中嵌套声明组件 B，
**前提**
组件重新渲染的时机:初始化或者状态更新时
**原因**
如果组件 B 在组件 A 中声明，那么当 A 组件状态更新时，都是把 B 当成一个新组件重新刷新，重置其状态

## jsx 不应该有什么

产生 jsx 的逻辑是呈现逻辑，所以不允许有副作用，呈现逻辑不应该有以下内容：

```sh
1.API调用
2.状态更新
3.定时器
4.对象或可变突变
```

唯一允许副作用的地方是事件处理程序内部和 useEffect

## React 是一个库不是一个框架

React 和 Vue 不一样，React 只是一个库，不是一个框架，他需要安装外部的包来完善自己，所以有很多基于 react 的框架，例如 Next、Remix 等
