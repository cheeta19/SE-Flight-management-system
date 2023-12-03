import './App.css';
import { Form, Input } from "antd";
import { useNavigate } from 'react-router-dom';

function App() {
  
  const navigate = useNavigate();
  const createAccount = () => {
    navigate('/Register')
  }
  
  return (
    <div className='main'>
      <div className='overlay'></div>
      <video src="Background_videos.mp4" autoPlay loop muted  className='vdo'/>
      <div className='container'>
        <div className='box'>
          <h1 style={{ textAlign: 'center' }}>Sign in</h1>
          <div className='content'>
            <div>
              <h1 style={{ marginTop: 10 }}>Email</h1>
              <h1 style={{ marginTop: 95 }}>Password</h1>
            </div>
            <div style={{ marginLeft: 25 }}>
              <Form>
                <Form.Item
                  name="Email"
                  rules={[
                    {
                      required: true,
                      type: "email",
                      message: "Please input your email",
                    },
                  ]}
                >
                  <Input style={{ width: 516, height: 65, background: 'white', borderRadius: 20, border: '1px black solid', fontSize: 32 }}
                    placeholder="example123@email.com" />
                </Form.Item>
                <Form.Item
                  name="Password"
                  rules={[
                    {
                      required: true,
                      message: "Please input your password",
                    },
                  ]}
                >
                  <Input.Password style={{ marginTop: 50, width: 516, height: 65, background: 'white', borderRadius: 20, border: '1px black solid', fontSize: 32 }} />
                </Form.Item>
                <Form.Item >
                  <button type='submit' className='button' style={{ position: 'absolute', top: 50, left: 450 }}>
                    Login
                  </button>
                </Form.Item>
              </Form>
              <button type='submit' className='button' style={{ position: 'absolute', top: 512, left: 860 }} onClick={createAccount}>
                Sign up
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
