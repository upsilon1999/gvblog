import { useState } from "react";

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <div>数字为{count}</div>
      <button onClick={() => setCount((count) => count + 1)}>点我加1</button>
    </>
  );
}

export default App;
