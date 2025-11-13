import logo from './logo.svg';
import './App.css';
import Header from './Header';
function Logo(){
  const userLogo=<img src={logo} alt={logo}/>
  return userLogo;
}
function App() {
  return (
    <>
    <Header title="Welcome" color="purple"/>
    <Logo/>
    </>
  );
}

export default App;
