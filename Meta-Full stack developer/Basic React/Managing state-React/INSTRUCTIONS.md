# Task

In the starter code of this code lab, you are given a Fruits component that has its own state. Based on this state, it outputs three fruits on the screen. Additionally, you have a FruitsCounter component which shows a message that reads: "Total fruits: 2".

Y​our task is to lift state up from the Fruits component to the App component, so that you can then pass the state information to both the Fruits component and the FruitsCounter component. This change to the app should fix the previously incorrect message of "Total fruits: 2".  The new message should be "Total fruits: 3".  However, the new message will not be just a hard-coded string. Instead, it should reflect the number of fruits that exist in the state variable, so based on how many fruits there are in the state array, this information should affect the output of the total number of fruits - as returned from the FruitsCounter component.

## Steps

**Step 1.** This task’s starting point is the stateless App component’s code:
```jsx
import Fruits from "./Fruits";
import FruitsCounter from "./FruitsCounter";

function App() {
  return (
    <div className="App">
      <h1>Where should the state go?</h1>
      <Fruits />
      <FruitsCounter />
    </div>
  );
}

export default App;
```
The first step of this task is to move the state from the `Fruits` component to `App` component.
```jsx
import React from "react";

function Fruits() {
    const [fruits] = React.useState([
        {fruitName: 'apple', id: 1},
        {fruitName: 'apple', id: 2},
        {fruitName: 'plum', id: 3},
    ]);

    return (
        <div>
            {fruits.map(f => <p key={f.id}>{f.fruitName}</p>)}
        </div>
    )
}

export default Fruits
```
Update the App component to pass this fruits state as props to both the Fruits and FruitsCounter components. 

**Step 2.** Once you've moved the state up from the `Fruits` component to the App component,Remove the state from the  Fruits component and update it to receive the  `fruits`  as prop and display the list of fruits dynamically based on the passed prop.  
```jsx
import React from "react";

function Fruits({ fruits }) {
  return (
    <div>
      {fruits.map((f) => (
        <p key={f.id}>{f.fruitName}</p>
      ))}
    </div>
  );
}

export default Fruits;
```
**Step 3.** Now,Modify the `FruitsCounter` component to accept the fruits prop and calculate the total fruits dynamically using the length property of the array. 

**Step 4.** 
- At the top of the lab environment, locate the Terminal menu. Click on it to open a dropdown, then select New Terminal.  Use the  `npm start` command to start the development server.   
- You can now view the App in your browser by navigating to localhost:3000. To view the output, click on the Browser Preview icon located on the left panel. It is the last icon in the panel.  
- In your browser, enter: http://localhost:3000 to see the output. The app displays the correct list of fruits.The total fruits count dynamically reflects the length of the fruits array . 
