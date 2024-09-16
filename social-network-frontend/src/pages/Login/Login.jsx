import React from 'react';
import { useLogin } from '../../services/api';
import './Login.css';
import logo from '../../assets/vibley-logo.png';

function Login() {
    const { loginData, handleChange, handleSubmit } = useLogin()

    return (
        <div className='singin-section'>
            <div className='signin-container'>
                <div className='left-panel'>
                    <img src={logo} alt="Vibley Logo" className="logo" /> 
                    <div className='welcome-message'>
                        <h1>Welcome to Vibley</h1>
                        <p>Sign in to continue access</p>
                    </div>
                </div>
                <div className='right-panel'>
                    <h2>Sign In</h2>
                    <form onSubmit={handleSubmit} method='POST'>
                        <div className='form-group'>
                            <label htmlFor="name">Email Address or Username</label>
                            <input 
                                type="text" 
                                id="name" 
                                value={loginData.name} 
                                onChange={handleChange} 
                                required 
                            />
                        </div>

                        <div className='form-group'>
                            <label htmlFor="password">Password</label>
                            <input 
                                type="text" 
                                id="password"
                                value={loginData.password} 
                                onChange={handleChange} 
                                required 
                            />
                        </div>
                        <button type="submit" className="btn continue-btn">Continue</button>
                        <div className="or-separator">Or</div>
                        <button type="button" className="btn register-btn">Register</button>
                    </form>
                </div>
            </div>
        </div>
    )
}

export default Login
