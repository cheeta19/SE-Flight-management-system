import { useNavigate } from 'react-router-dom';
import './Register.css';
import Upload from './Uploads';
import { Form, Input, Select } from 'antd'
import { Base64 } from './Uploads';
import { useEffect } from 'react';

function App() {

    const navigate = useNavigate();
    const createAccount = () => {
        navigate('/')
    }

    useEffect(() => {
        console.log('B=', Base64)
    }, []);

    return (
        <div className="main">
            <div className='overlay'></div>
            <video src="SignUp_video.mp4" autoPlay loop muted />
            <div className='container'>
                <h2 style={{ marginTop: -570 }}>Sign up</h2>
                <div style={{ position: 'absolute', left: -220, top: 100 }}>
                    <Upload />
                </div>
                <div style={{ position: 'absolute', marginTop: 1200 }}>
                    <Form>
                        <Form.Item
                            name="Firstname"
                            rules={[
                                {
                                    required: true,
                                    message: "Please input your firstname",
                                },
                            ]}
                        >
                            <Input className='input' placeholder="Firstname" />
                        </Form.Item>
                        <Form.Item
                            name="Lastname"
                            rules={[
                                {
                                    required: true,
                                    message: "Please input your lastname",
                                },
                            ]}
                        >
                            <Input className='input' placeholder="Lastname" />
                        </Form.Item>
                        <Form.Item
                            name="Passport"
                            rules={[
                                {
                                    required: true,
                                    message: "Please input your passport",
                                },
                            ]}
                        >
                            <Input className='input' placeholder="Passport" />
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
                            <Input.Password className='input' placeholder="Password" />
                        </Form.Item>
                        <Form.Item
                            name="Email"
                            rules={[
                                {
                                    required: true,
                                    type: 'email',
                                    message: "Please input your email",
                                },
                            ]}
                        >
                            <Input className='input' placeholder="Email" />
                        </Form.Item>

                        <Form.Item
                            name="Phone"
                            rules={[
                                {
                                    required: true,
                                    message: "Please input your phone",
                                },
                            ]}
                        >
                            <Input className='input' placeholder="Phone" />
                        </Form.Item>
                        <Form.Item >
                            <button type='submit' className='button' style={{ position: 'absolute', top: 10, left: 750 }} onClick={createAccount}>
                                Sign Up
                            </button>
                        </Form.Item>
                    </Form>
                </div>
            </div>
        </div>
    )
}


export default App;