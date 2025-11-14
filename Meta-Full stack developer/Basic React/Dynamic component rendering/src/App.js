import LoginButton from './LoginButton';
import LogoutButton from './LogoutButton';
import HomePage  from './HomePage';
import { useState } from 'react';
import './App.css';

function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const login = () => setLoggedIn(true);
  const logout = () => setLoggedIn(false);
  return (
    <>
      {loggedIn ? (
        <>
          <HomePage />
          <LogoutButton logout={logout} />
        </>
        
      ) : (
          <>
            <LoginButton login={login}/>
            </>
      )}
    </>
  )
}

export default App;
