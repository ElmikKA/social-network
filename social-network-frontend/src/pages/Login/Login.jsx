import React from 'react'

import { useLogin } from '../../services/api'
import { Link } from 'react-router-dom'

function Login() {
    const { loginData, handleChange, handleSubmit } = useLogin()

    return (
        <form onSubmit={handleSubmit} method='POST' style={{ display: 'flex', flexDirection: 'column', maxWidth: '400px', margin: '0 auto' }}>
            <label htmlFor="name">Name/email</label>
            <input type="text" id="name" placeholder='name/email' required value={loginData.name} onChange={handleChange} />

            <label htmlFor="password">Password</label>
            <input type="text" id="password" required placeholder='password' value={loginData.password} onChange={handleChange} />
            <button type='submit'>LogIn</button>
        </form>
    )
}

export default Login
