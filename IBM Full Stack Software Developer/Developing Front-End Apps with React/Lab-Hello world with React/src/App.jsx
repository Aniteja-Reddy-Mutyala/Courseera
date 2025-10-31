function App(props) {
    const currDate = new Date();
  return (
    <>
      <h1>Aniteja Reddy Mutyala</h1>
      <h2>The time is {currDate.toLocaleTimeString()}</h2>
      </>
  );
}

export default App;