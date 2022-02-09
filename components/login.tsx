import Button from '@nextui-org/react/Button';
import Modal from '@nextui-org/react/Modal';
import Text from '@nextui-org/react/Text';
import Input from '@nextui-org/react/Input';
import Checkbox from '@nextui-org/react/Checkbox';
import Link from '@nextui-org/react/Link';
import Row from '@nextui-org/react/Row';
import React from 'react';
import axios from "axios";
import type { AxiosResponse } from "axios";
interface LoginProps {
  handleLogin: (token: string, remember: boolean) => void;
  visible: boolean;
}
const Login = (props: LoginProps) => {
type FormStatus = "default" | "primary" | "secondary" | "success" | "warning" | "error"
const [status, setStatus] = React.useState<FormStatus>('default');
const [email, setEmail] = React.useState('');
const [password, setPassword] = React.useState('');
const [remember, setRemember] = React.useState(false);
const signInHandler = async () => {
  try {
    interface LoginResponse extends AxiosResponse {
      data: {
        token?: string;
      };
    }
    const resp  = await axios.post("/api/v1/user/login", {
      email: email,
      username: email,
      password: password,
    }) as LoginResponse;
    if(!resp.data.token) {
      throw new Error("No token returned");
    }
    props.handleLogin(resp.data.token, remember);
  } catch (err) {
    setStatus('error');
  }
}
return (
    <Modal
        blur
        closeButton
        aria-labelledby="Login form"
        open={props.visible}
    >
        <Modal.Header>
            <Text size={18} className="mx-1">
              Login
            </Text>
        </Modal.Header>
        <Modal.Body>
            <Input
                status={status}
                value={email}
                onClick={() => {
                  if(status === 'error') {
                    setStatus('default')
                    setEmail('')
                    setPassword('')
                  }
                }}
                onChange={(e) => setEmail(e.target.value)}
                required
                clearable
                bordered
                fullWidth
                color="primary"
                size="lg"
                placeholder="Email or Username"
            />
            <Input.Password
                status={status}
                onClick={() => {
                  if(status === 'error') {
                    setStatus('default')
                    setPassword('')
                    setEmail('')
                  }
                }}
                onChange={(e) => setPassword(e.target.value)}
                value={password}
                required
                clearable
                bordered
                fullWidth
                color="primary"
                size="lg"
                placeholder="Password"
            />
            <Row justify="space-between">
            <Checkbox
              onChange={(e) => setRemember(e.target.checked)}
            >
                <Text size={14}>
                Remember me
                </Text>
            </Checkbox>
            <Text size={14}>
              <Link href="iforgot">
                Forgot password?
              </Link>
            </Text>
            </Row>
        </Modal.Body>
        <Modal.Footer>
            <Button auto flat color="error" onClick={()=> {
              setStatus('default');
              setPassword('');
              setEmail('');
            }}>
            Close
            </Button>
            <Button auto onClick={signInHandler}>
            Sign in
            </Button>
        </Modal.Footer>
    </Modal>
);
};
export default Login;
