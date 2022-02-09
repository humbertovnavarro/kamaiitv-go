import { Button } from '@nextui-org/react';
import React from 'react';
import Login from '../components/login';
const Page = () => {
  function handleLogin(token:string, remember:boolean) {
    console.log(token, remember);
  }
  const [loginVisible, setLoginVisible] = React.useState(false);
  return (
    <div>
      <Button onClick={() => setLoginVisible(!loginVisible)}>Login</Button>
      <Login visible={loginVisible} handleLogin={handleLogin}/>
    </div>
  )
}
export default Page;
