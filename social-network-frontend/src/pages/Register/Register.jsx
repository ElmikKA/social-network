import React from 'react'
import { useRegister } from '../../api'
import { useNavigate } from 'react-router-dom';
import logo from '../../assets/vibley-logo.png';
import './Register.css';

const Register = () => {
    const { registerData, handleChange, handleFileChange, handleSubmit } = useRegister()

    return (
        <div className='register-section'>
            <div className='register-container'>
                <div className='register-left-panel'>
                    <img src={logo} alt="vibley logo" className="logo" /> 
                    <div className='welcome-message'>
                        <h1>Fill out the form</h1>
                        <p>And get access to Vibley!</p>
                    </div>
                </div>
                <div className='register-right-panel'>
                    <div className='registration-header'>
                        <h2>Registration</h2>
                    </div>
                    <form onSubmit={handleSubmit} method='POST' className='registration-form'>
                        <div className='register-form-group'>
                            <div className='register-form-field'>
                                <label htmlFor="email">Email</label>
                                <input type="email" id="email" placeholder="Email" required value={registerData.email} onChange={handleChange} />
                            </div>
                            <div className='register-form-field'>
                                <label htmlFor="firstName">First Name</label>
                                <input type="text" id="firstName" placeholder="First Name" required value={registerData.firstName} onChange={handleChange} />
                            </div>
                        </div>

                        <div className='register-form-group'>
                            <div className='register-form-field'>
                                <label htmlFor="lastName">Last Name</label>
                                <input type="text" id="lastName" placeholder="Last Name" required value={registerData.lastName} onChange={handleChange} />
                            </div>
                            <div className='register-form-field'>
                                <label htmlFor="dateOfBirth">Date of Birth</label>
                                <input type="date" id="dateOfBirth" required value={registerData.dateOfBirth} onChange={handleChange} />
                            </div>
                        </div>

                        <div className='register-form-group'>
                            <div className='register-form-field' style={{width: '100%'}}>
                                <label htmlFor="password">Password</label>
                                <input type="password" id="password" placeholder="Password" required value={registerData.password} onChange={handleChange} />
                            </div>
                        </div>

                        <div className='optional-header'>
                            <h1>Optional</h1>
                        </div>

                        <div className='register-form-group'>
                            <div className='register-form-field'>
                                <label htmlFor="avatar">Avatar</label>
                                <input type="file" id="avatar" accept="image/*" onChange={handleFileChange} />
                            </div>
                            <div className='register-form-field'>
                                <label htmlFor="nickname">Nickname</label>
                                <input type="text" id="nickname" placeholder="Nickname" value={registerData.nickname} onChange={handleChange} />
                            </div>
                        </div>

                        <div className='register-form-group'>
                            <div className='register-form-field' style={{width: '100%'}}>
                                <label htmlFor="aboutMe">About Me</label>
                                <textarea id="aboutMe" placeholder="Tell us about yourself" value={registerData.aboutMe} onChange={handleChange} />
                            </div>
                        </div>

                        <button type="submit" className="btn register-btn">
                            Register
                        </button>
                    </form>
                </div>
            </div>
        </div>
    )
}

export default Register


