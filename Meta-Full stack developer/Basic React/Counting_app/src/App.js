import { useState } from 'react';
import './App.css';

function App() {
  const [num, setNum] = useState(0);
  return (
    <>
      <h1>Current number is :{num}</h1>
      <button onClick={()=>setNum(num+1)}>Increment</button>
      <button onClick={()=>setNum(num-1)}>Decrement</button>
    </>
  );
}

export default App;
