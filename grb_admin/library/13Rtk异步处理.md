## 单元重点

```sh
1.React Router的高级功能
2.Redux的现实应用
```

## 技巧

### input的默认值

```jsx
function App({username}){
    return (
        <div>
        	<input type="text" defaultValue={username}/>
        </div>
    )
}
```

input标签可以用defaultValue属性设置默认值

## Redux

在现在的项目中我们将使用RTK，而不用在安装传统redux，安装步骤如下

```sh
npm i @reduxjs/toolkit react-redux
```

初步应用RTK

```js
import {createSlice } from '@reduxjs/toolkit';


const initialState = {
  username: '',
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    updateName(state, action) {
      state.username = action.payload;
    },
  },
});

export const { updateName } = userSlice.actions;

export default userSlice.reducer;
```

### selector函数

我们都知道在使用redux读取状态时会使用到useSelector钩子，传入一个回调函数，为了功能的通用性我们可以将这个回调函数抽离出来存放在对应的XxSlice中，这个被抽离的函数就被称为selector函数，该函数的命名规范是getXx，例如

**未抽离前**

`userSlice.js`

```js
import {createSlice } from '@reduxjs/toolkit';


const initialState = {
  username: '',
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    updateName(state, action) {
      state.username = action.payload;
    },
  },
});

export const { updateName } = userSlice.actions;

export default userSlice.reducer;
```

`store.js`

```js
import { configureStore } from '@reduxjs/toolkit';
import userReducer from './features/user/userSlice';

const store = configureStore({
  reducer: {
    user: userReducer,
  },
});

export default store;
```

`Login.jsx`

```jsx
import { useSelector } from 'react-redux';
function Login(){
    const usename = useSelector((store)=>store.user.username)
    return <p>username</p>
}
```

**抽离后**

`userSlice.js`

```js
import {createSlice } from '@reduxjs/toolkit';


const initialState = {
  username: '',
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    updateName(state, action) {
      state.username = action.payload;
    },
  },
});

export const { updateName } = userSlice.actions;

//暴露抽离的selector函数，命名规范是getXx
export const getUserName = (store)=>{
    return store.user.username
}

export default userSlice.reducer;
```

使用

```jsx
import { useSelector } from 'react-redux';
import {getUserName} from './userSlice.js'
function Login(){
    //使用selector函数
    const usename = useSelector(getUserName)
    return <p>username</p>
}
```

>优势：1.对于同样的逻辑可以复用 2.将selector函数统一管理，规范命名后可读性更高 3.对于复杂大的逻辑处理更友好

**注意**

这样使用selector函数可能会在大的应用程序中存在性能问题，这时候一般会借助第三方库，例如`reselect`

#### 带参数的selector函数

有时候selector函数可能会携带参数，我们先来看看不抽离时的写法

```jsx
function Login(){
    let id = 64
    const username = useSelector((store)=>{
        return store.user.userList.find(item=>item.userId==id)?.userName??""
    })
}
```

当我们抽离时

```js
function Login(){
    let id = 64
    const username = useSelector(getUserName(id))
}

export const getUserName = id => {
    return (store)=>{
        return store.user.userList.find(item=>item.userId==id)?.userName??""
    }
}
```

以上时箭头函数的完整写法，你可能看到简化版，省略了{}和return，例如

```js
export const getUserName = id => store => store.user.userList.find(item=>item.userId==id)?.userName??""
```

### caseReducers

当我们想在xxSlice内部使用器reducer时，可以用`xxSlice.caseReducers.方法名`，举例

```js
import {createSlice } from '@reduxjs/toolkit';


const initialState = {
    users:[
        {userId:66,username:"lisa",age:19}
    ]
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    //根据userId过滤，即删除该user
    deleteUser(state, action) {
      state.users = state.users.filter((item) => item.userId !== action.payload);
    },
    //减少岁数，当岁数为0时删掉该user
    decreaseItemAge(state,action){
        //找到对应的user项
        const item = state.users.find((item) => item.userId === action.payload);
        
        //该项age减1
        item.age --
        
        //当age为0时删除该角色，
        /*
        	我们可能首先想到的就是复制deleteUser的逻辑
        	实际上我们是可以直接调用其他reducer的方法的，
        */
        if(item.age===0){
            //此处的state和action按照变量传递原理，就是本次decreaseItemAge的参数
            userSlice.caseReducers.deleteUser(state,action)
        }
    }
  },
});

export const { updateName } = userSlice.actions;

export default userSlice.reducer;
```

### RTK的异步处理

之前我们处理异步操作的时候在RTK的写法中混用传统redux，例如

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

RTK中也有自己的异步处理方案，步骤如下

* 先用createAsyncThunk生成要异步的操作函数，语法

  ```js
  const 实例名 = createAsyncThunk(reducer方法名,异步操作)
  //createAsyncThunk中异步操作的返回值就是该异步操作成功时的action.payload
  ```

  这个实例就是一个promise对象也是我们之前所说的action函数，我们待会儿会根据该实例的状态(pending、fulfilled、rejected)做出应对，举例

  ```js
  import {createSlice,createAsyncThunk } from '@reduxjs/toolkit';
  import getAdress from "./api/adress"
  
  const initialState = {
  	username: '',
      status: 'idle',
      address: '',
      error: '',
  };
  
  //使用createAsyncThunk创建异步promise实例
  //这个异步实例的返回值就是action.payload
  export const fetchAddress = createAsyncThunk(
    'user/fetchAddress',
    async function () {
       const addressObj = await getAddress(position);
       const address = `${addressObj?.locality}--${addressObj?.countryName}`;
       return adress;
    }
  );
  
  const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
     	updateName(state, action) {
        state.username = action.payload;
      },
    },
  });
  
  export const { updateName } = userSlice.actions;
  
  export default userSlice.reducer;
  ```

* 在reducers同级增加extraReducers，值是一个函数，该函数有一个builder参数，调用buildrer上的方法对异步实例的不同状态做处理，语法

  ```js
  extraReducers: (builder) =>
      builder
        .addCase(异步实例A.状态1, (state, action) => {
          对应操作
        })
        .addCase(异步实例A.状态2, (state, action) => {
          对应操作
        })
        .addCase(异步实例A.状态2, (state, action) => {
          对应操作
        })
        .addCase(异步实例B.状态1, (state, action) => {
          对应操作
        })
      ,
  ```

  addcase是链式语法，第一个参数是异步实例的对应状态，第二个参数是操作函数，举例

  ```js
  import {createSlice,createAsyncThunk } from '@reduxjs/toolkit';
  import getAdress from "./api/adress"
  
  const initialState = {
  	username: '',
      status: 'idle',
      address: '',
      error: '',
  };
  
  //使用createAsyncThunk创建异步promise实例
  export const fetchAddress = createAsyncThunk(
    'user/fetchAddress',
    async function () {
       const addressObj = await getAddress(position);
       const address = `${addressObj?.locality}--${addressObj?.countryName}`;
       //payload of the fulfilled state
       return adress;
    }
  );
  
  const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
     	updateName(state, action) {
        state.username = action.payload;
      },
    },
    extraReducers: (builder) =>
      builder
        .addCase(fetchAddress.pending, (state, action) => {
          state.status = 'loading';
        })
        .addCase(fetchAddress.fulfilled, (state, action) => {
          state.address = action.payload.address;
          state.status = 'idle';
        })
        .addCase(fetchAddress.rejected, (state, action) => {
          state.status = 'error';
          state.error =
            'There was a problem getting your address. Make sure to fill this field!';
        }),
  });
  
  export const { updateName } = userSlice.actions;
  
  export default userSlice.reducer;
  ```

* dispatch派发操作

  ```jsx
  function Login(){
      const dispatch = useDispatch()
      //这个fetchAddress()就是createAsyncThunk实例名
      //也是我们之前的action create函数
      dispatch(fetchAddress())
  }
  ```

## React Router高级

### Fetching Data Without Navgation useFetcher

在让导航不跳转到下一页的情况下获取或改变数据。使用举例

```js
import { useFetcher } from 'react-router-dom';
import { useEffect } from 'react';
function App(){
    const fetcher = useFetcher()

    useEffect(
        function () {
          if (!fetcher.data && fetcher.state === 'idle') fetcher.load('/menu');
        },
        [fetcher]
    );
}
```

#### 官网文档

在 HTML/HTTP 中，数据突变和加载是通过导航来模拟的： `<a href>` 和 `<form action>` 。两者都会在浏览器中引起导航。与 React Router 对应的是`Link`和`Form` 。

但有时您需要在导航之外调用 `loader` ，或调用 `action`（并获取页面上的数据以重新验证），而无需更改 URL。或者您需要同时进行多个突变。

与服务器的许多交互都不是导航事件。此钩子可让你在不导航的情况下在用户界面插入`action`和`loader`。

这在需要时非常有用：

- 获取与用户界面路由（弹出窗口、动态表单等）无关的数据
- 无需导航即可将数据提交至操作（共享组件，如即时通讯注册）
- 在一个列表中处理多个并发提交（典型的 "待办事项应用程序 "列表，您可以点击多个按钮，所有按钮都应同时待处理）
- 无限滚动容器
- 以及更多！

如果您正在构建一个高度交互、"类似应用程序 "的用户界面，那么您将经常 `useFetcher` 。

```jsx
import { useFetcher } from "react-router-dom";

function SomeComponent() {
  const fetcher = useFetcher();

  // call submit or load in a useEffect
  React.useEffect(() => {
    fetcher.submit(data, options);
    fetcher.load(href);
  }, [fetcher]);

  // build your UI with these properties
  fetcher.state;
  fetcher.formData;
  fetcher.json;
  fetcher.text;
  fetcher.formMethod;
  fetcher.formAction;
  fetcher.data;

  // render a form that doesn't cause navigation
  return <fetcher.Form />;
}
```

`Fetchers`有很多内置行为：

- 自动处理获取中断时的取消操作

- 使用 POST、PUT、PATCH、DELETE 提交时，首先调用操作

  - 操作完成后，页面上的数据会重新验证，以捕捉可能发生的任何变化，从而自动保持用户界面与服务器状态同步

- 当多个

  ```
  fetchers
  ```

  同时运行时，它会...

  - 在每次登陆时提交最新的可用数据
  - 确保无论响应返回的顺序如何，都不会有陈旧的负载覆盖较新的数据

- 通过渲染最近的 `errorElement` 来处理未捕获的错误（就像从 `<Link>` 或 `<Form>` 进行正常导航一样）。

- 如果调用的操作/加载器返回重定向，应用程序将重定向（就像从 `<Link>` 或 `<Form>` 进行普通导航一样）。

##### `key`

默认情况下， `useFetcher` 会生成一个唯一的`fetcher`，该`fetcher`的作用域为该组件（不过，在运行过程中，该`fetcher`可能会在 `useFetchers()`中查找）。如果你想用自己的 `key` 来识别一个 fetcher，以便从应用程序的其他地方访问它，可以使用 `key` 选项来实现：

```jsx
function AddToBagButton() {
  const fetcher = useFetcher({ key: "add-to-bag" });
  return <fetcher.Form method="post">...</fetcher.Form>;
}

// Then, up in the header...
function CartCount({ count }) {
  const fetcher = useFetcher({ key: "add-to-bag" });
  const inFlightCount = Number(
    fetcher.formData?.get("quantity") || 0
  );
  const optimisticCount = count + inFlightCount;
  return (
    <>
      <BagIcon />
      <span>{optimisticCount}</span>
    </>
  );
}
```

##### `fetcher.Form`

就像 `<Form>` 一样，只是它不会导致导航。(你会克服 JSX 中的圆点......我们希望！）。

```jsx
function SomeComponent() {
  const fetcher = useFetcher();
  return (
    <fetcher.Form method="post" action="/some/route">
      <input type="text" />
    </fetcher.Form>
  );
}
```

##### `fetcher.load(href, options)`

从路由`loader`中加载数据。

```jsx
import { useFetcher } from "react-router-dom";

function SomeComponent() {
  const fetcher = useFetcher();

  useEffect(() => {
    if (fetcher.state === "idle" && !fetcher.data) {
      fetcher.load("/some/route");
    }
  }, [fetcher]);

  return <div>{fetcher.data || "Loading..."}</div>;
}
```

虽然一个 URL 可能匹配多个嵌套路由，但 `fetcher.load()` 调用只会调用叶匹配（或`索引路由`的父路由）上的`loader`。

如果您发现自己在点击处理程序中调用此函数，您可以使用 `<fetcher.Form>` 来简化代码。

> 为重新验证的一部分，页面上激活的任何 `fetcher.load` 调用都将重新执行（在导航提交、另一个`fetcher`提交或 `useRevalidator()` 调用之后）。

##### `options.unstable_flushSync`

`unstable_flushSync` 选项告诉 React Router DOM 将此 `fetcher.load` 的初始状态更新封装在 `ReactDOM.flushSync`调用中，而不是默认的 `React.startTransition`中。这样就可以在更新刷新到 DOM 后立即执行同步 DOM 操作。

> 请注意，该应用程序接口标记为不稳定状态，在未发布重大版本之前可能会出现破坏性更新。

##### `fetcher.submit()`

`<fetcher.Form>` 的命令式版本。如果是由用户交互启动获取，则应使用 `<fetcher.Form>` 。但如果是由程序员启动获取（而不是响应用户点击按钮等），则应使用该函数。

例如，您可能希望在闲置一定时间后注销用户：

```jsx
import { useFetcher } from "react-router-dom";
import { useFakeUserIsIdle } from "./fake/hooks";

export function useIdleLogout() {
  const fetcher = useFetcher();
  const userIsIdle = useFakeUserIsIdle();

  useEffect(() => {
    if (userIsIdle) {
      fetcher.submit(
        { idle: true },
        { method: "post", action: "/logout" }
      );
    }
  }, [userIsIdle]);
}
```

`fetcher.submit` 是 `useSubmit`调用 fetcher 实例的包装器，因此也接受与`useSubmit`相同的选项。

如果要提交索引路由，请使用`?index` 参数。

如果您发现自己在点击处理程序中调用此函数，您可以使用 `<fetcher.Form>` 来简化代码。

##### `fetcher.state`

您可以通过 `fetcher.state` 了解 fetcher 的状态。它将是:

- **idle** - 没有获取任何信息。
- **submitting** - 由于使用 POST、PUT、PATCH 或 DELETE 提交了取件，路由操作被调用
- **loading** - fetcher 正在调用`loader`（来自 `fetcher.load` ），或在单独提交或调用 `useRevalidator` 后正在重新验证

##### `fetcher.data`

加载器或操作返回的数据存储在这里。数据一旦设置完毕，即使重新加载和重新提交，也会在获取器上持续存在。

```jsx
function ProductDetails({ product }) {
  const fetcher = useFetcher();

  return (
    <details
      onToggle={(event) => {
        if (
          event.currentTarget.open &&
          fetcher.state === "idle" &&
          !fetcher.data
        ) {
          fetcher.load(`/product/${product.id}/details`);
        }
      }}
    >
      <summary>{product.name}</summary>
      {fetcher.data ? (
        <div>{fetcher.data}</div>
      ) : (
        <div>Loading product details...</div>
      )}
    </details>
  );
}
```

##### `fetcher.formData`

当使用 `<fetcher.Form>` 或 `fetcher.submit()` 时，表单数据可用于构建优化的用户界面。

```jsx
function TaskCheckbox({ task }) {
  let fetcher = useFetcher();

  // while data is in flight, use that to immediately render
  // the state you expect the task to be in when the form
  // submission completes, instead of waiting for the
  // network to respond. When the network responds, the
  // formData will no longer be available and the UI will
  // use the value in `task.status` from the revalidation
  let status =
    fetcher.formData?.get("status") || task.status;

  let isComplete = status === "complete";

  return (
    <fetcher.Form method="post">
      <button
        type="submit"
        name="status"
        value={isComplete ? "complete" : "incomplete"}
      >
        {isComplete ? "Mark Complete" : "Mark Incomplete"}
      </button>
    </fetcher.Form>
  );
}
```

##### `fetcher.json`

使用 `fetcher.submit(data, { formEncType: "application/json" })` 时，提交的 JSON 可通过 `fetcher.json` 获取。

##### `fetcher.text`

使用 `fetcher.submit(data, { formEncType: "text/plain" })` 时，提交的文本可通过 `fetcher.text` 获取。

##### `fetcher.formAction`

告诉您提交表单的操作`url`。

```jsx
<fetcher.Form action="/mark-as-read" />;

// when the form is submitting
fetcher.formAction; // "mark-as-read"
```

##### `fetcher.formMethod`

告诉您提交表单的方法：get、post、put、patch 或 delete。

```jsx
<fetcher.Form method="post" />;

// when the form is submitting
fetcher.formMethod; // "post"
```

> `fetcher.formMethod` 字段为小写，没有 `future.v7_normalizeFormMethod` 。为了与 `fetch()` 在 v7 中的行为保持一致，我们正在将其规范化为大写，因此请升级 React Router v6 应用程序以采用大写 HTTP 方法。

### Updating Data Without Navigation

P325

就是提交表单时不触发导航，使用到`fetcher.Form`