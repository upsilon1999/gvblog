## 单元重点

```sh
1.Tailwind CSS入门
```

## Tailwind CSS

Tailwind CSS是一个css框架，带有许多预先定义好的实用样式，通过将这些预先样式的组合来实现网站。

### Tailwind CSS的优缺点

```sh
【优点】
1.不用自己设计类名
2.不必一直在组件和样式文件中跳转
3.因为项目使用了统一的样式，团队合作变得容易
【缺点】
1.html和jsx变得贴别难看和难读，因为为了设计一个元素，可能需要用到10个甚至20个class
2.必须学习所有你要用到或可能用到的Tailwind CSS样式
3.需要在每个项目上安装和配置Tailwind CSS
4.可能会想要放弃原生css
```

### 安装Tailwind CSS

遵循Tailwind CSS官网的步骤安装Tailwind CSS。根据自己的项目来配置，

```sh
1.安装-->Framework Guides-->vite-->Using React
```

根据官网流程一步步操作，以v3.4.1为例

**构建react项目**

```sh
npm create vite@latest my-project -- --template react
cd my-project
```

**为项目安装Tailwind CSS**

```sh
#第一步 安装插件
npm install -D tailwindcss postcss autoprefixer

#第二步 生成相关配置文件
npx tailwindcss init -p
```

操作完后，会有两个新文件

```sh
tailwind.config.js
postcss.config.js
```

**配置位置**

就是告诉Tailwind CSS我们将在哪里用到Tailwind CSS，在`tailwind.config.js`中配置

```js
/** @type {import('tailwindcss').Config} */
export default {
  //指定使用Tailwind CSS的文件
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
```

**重定向index.css**

将以下内容加入到`./src/index.css`的顶部

```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

### 安装Tailwind CSS拓展

#### 安装Vscode中的Tailwind CSS拓展

```sh
【Tailwind CSS IntelliSense】
1.将向我们显示Tailwind CSS 的类的底层原生css
2.自动补全
```

#### 给项目安装一个拓展

```sh
【tailwind prettier extension】
1.他会自动帮tailwind css类名排序，方便我们查找
```

这是个github项目，搜索得到安装方法

```sh
npm install -D prettier prettier-plugin-tailwindcss
```

安装好后书写配置文件,`prettier.config.cjs`

ps：这个配置文件不是专门为tailwindcss设计的，指示说tailwindcss作为一个插件被引入

```js
module.exports = {
  plugins:[require("prettier-plugin-tailwindcss")]
}
```

**注意事项**

```sh
由于我们安装了prettier插件，事先带有`prettier.config.js`,所以为了避免冲突我们新建的配置文件命名为`prettier.config.cjs`
```

### 熟悉文档

以下内容只是说明那些文档在项目中很常用，实际上在开发中我们就是直接查文档，因为文档写的很详细。

#### 如何处理颜色

记住多看文档，

颜色是

```sh
https://www.tailwindcss.cn/docs/customizing-colors
颜色名-数字
```

文本颜色

```sh
https://www.tailwindcss.cn/docs/text-color
text-颜色名-数字
```

背景色

```sh
https://www.tailwindcss.cn/docs/background-color
bg-限定色
bg-颜色名-数字
```

文本样式

```sh
【Font Size】
字体大小

https://www.tailwindcss.cn/docs/font-size

【Font Weight】
字体权重
https://www.tailwindcss.cn/docs/font-weight

【Letter Spacing】
文字间距
https://www.tailwindcss.cn/docs/letter-spacing
```

**拓展**

实际上我们是可以自己调的，例如

```sh
text-xl 这是tailwindcss的提供的样式，设置字体大小

text-[100px] 我们可以自定义字体大小为100px
```

盒模型边距和边框、显示

```sh
【Margin】
边距

https://www.tailwindcss.cn/docs/margin
```

```sh
【Padding】
https://www.tailwindcss.cn/docs/padding

【Space Between】
网格化设置间距
https://www.tailwindcss.cn/docs/space

【Display】
https://www.tailwindcss.cn/docs/display
```

#### 响应式设计

```sh
【responsive-design】

https://www.tailwindcss.cn/docs/responsive-design#overview
```

根据前缀在不同情况下进行生效

#### flexbox

```sh
【flex】
https://www.tailwindcss.cn/docs/flex
```

#### Grid

```sh
【Grid Template Rows】
https://www.tailwindcss.cn/docs/grid-template-rows
```

#### 元素状态和过度

所谓元素状态，就是形如按钮hover一类的

```jsx
<p className="hover:类名1 hover:类名2"></p>
```

文档

```sh
搜索hover跳到此文档
https://www.tailwindcss.cn/docs/hover-focus-and-other-states

搜索transition，这是动画相关文档
https://www.tailwindcss.cn/docs/transition-property
```

#### 绝对定位和z-index

```sh
【Position】
https://www.tailwindcss.cn/docs/position

【Z-Index】
https://www.tailwindcss.cn/docs/z-index
```



#### @apply

有时候我们想把一些tailwindcss的组合重命名，例如

```jsx
<input className="rounded-full border border-stone-200 px-4 py-2 text-sm transition-all duration-300 placeholder:text-stone-400 focus:outline-none focus:ring focus:ring-yellow-400 md:px-6 md:py-3;"/>
```

如果这个样式组合要使用多次，我们就可以考虑将他设为一个样式组件，在`index.css`中

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

//设置样式组件
@layer components {
    //将样式组合设置为我们自定义类名
    .my-input {
        @apply rounded-full border border-stone-200 px-4 py-2 text-sm transition-all focus:ring-yellow-400
        duration-300 placeholder:text-stone-400 focus:outline-none focus:ring  md:px-6 md:py-3;
    }
}
```

使用自定义样式组件

```jsx
<input className="my-input"/>
```

当我们要包含的是普通css就不需要用@apply引用，例如

```css
@layer components {
    //就是将原生css包裹，让tailwind css可识别
    //将样式设置为我们自定义类名
    .loader {
        width: 45px;
        aspect-ratio: 0.75;
        --c: no-repeat linear-gradient(theme(colors.stone.800) 0 0);
        background: var(--c) 0% 50%, var(--c) 50% 50%, var(--c) 100% 50%;
        background-size: 20% 50%;
        animation: loading 1s infinite linear;
    }
     @keyframes loading {
        20% {
          background-position: 0% 0%, 50% 50%, 100% 50%;
        }
        40% {
          background-position: 0% 100%, 50% 0%, 100% 50%;
        }
        60% {
          background-position: 0% 50%, 50% 100%, 100% 0%;
        }
        80% {
          background-position: 0% 50%, 50% 50%, 100% 100%;
        }
      }
}
```

这样的话我们就可以将index.css中需要保留的原生css全交给tailwind css管理，有利于项目后续维护

#### 动态使用tailwind css的案例

```jsx
import { Link } from 'react-router-dom';

function Button({ children, disabled, to, type }) {
  const base =
    'inline-block text-sm rounded-full bg-yellow-400 font-semibold uppercase tracking-wide text-stone-800 transition-colors duration-300 hover:bg-yellow-300 focus:bg-yellow-300 focus:outline-none focus:ring focus:ring-yellow-300 focus:ring-offset-2 disabled:cursor-not-allowed';

  const styles = {
    primary: base + ' px-4 py-3 md:px-6 md:py-4',
    small: base + ' px-4 py-2 md:px-5 md:py-2.5 text-xs',
    secondary:
      'inline-block text-sm rounded-full border-2 border-stone-300 font-semibold uppercase tracking-wide text-stone-400 transition-colors duration-300 hover:bg-stone-300 hover:text-stone-800 focus:bg-stone-300 focus:text-stone-800 focus:outline-none focus:ring focus:ring-stone-200 focus:ring-offset-2 disabled:cursor-not-allowed px-4 py-2.5 md:px-6 md:py-3.5',
  };

  if (to){
    return (
      <Link to={to} className={styles[type]}>
        {children}
      </Link>
    );
  }
    

  return (
    <button disabled={disabled} className={styles[type]}>
      {children}
    </button>
  );
}

export default Button;
```

因为tailwindcss实际上和普通css使用差不多，最后都是提供一个字符串。

以上的也给我们提供了一个封装UI组件的思路，通过传入的prop动态控制样式。

#### 粗浅讨论tailwind css配置

来到`tailwind.config.js`

```js
/** @type {import('tailwindcss').Config} */
//eslint-disable-next-line
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
```

`//eslint-disable-next-line`是我们加的注释用于去除eslint的警告，可以通过

```sh
https://www.tailwindcss.cn/docs/configuration
(可以参考和他同一分类的文档来定制更多)
```

也就是定制选项来配置tailwind css，注意覆盖与增加

**覆盖的案例**

```js
/** @type {import('tailwindcss').Config} */
//eslint-disable-next-line
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    //写在这里将覆盖掉theme的colors对象
    colors:{
		pizza:"#123456"
    }
    extend: {},
  },
  plugins: [],
}
```

**增加的案例**

```js
/** @type {import('tailwindcss').Config} */
//eslint-disable-next-line
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: { 
    extend: {
        //写在extend就是扩展theme的colors对象
        colors:{
        	pizza:"#123456"
        }
  	},
  },
  plugins: [],
}
```

**拓展**

我们常用vh来获取视口高度，他在移动端是不准确的，现在更推荐使用dvh，例如

```css
.happy{
    height:100dvh;
}
```



#### 小结

```sh
所谓tailwindcss就是一个别人写好的大的css库，开发时一个个套样式，通过文档来使用即可
```







