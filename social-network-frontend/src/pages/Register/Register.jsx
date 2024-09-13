import React from 'react'
import { useRegister } from '../../services/api'

const Register = () => {
    const { registerData, handleChange, handleFileChange, handleSubmit } = useRegister()

    return (
        <form onSubmit={handleSubmit} method='POST' style={{ display: 'flex', flexDirection: 'column', maxWidth: '400px', margin: '0 auto' }}>
            <label htmlFor="name">Name</label>
            <input type="text" id="name" placeholder="Name" required value={registerData.name} onChange={handleChange} />

            <label htmlFor="email">Email</label>
            <input type="email" id="email" placeholder="Email" required value={registerData.email} onChange={handleChange} />

            <label htmlFor="password">Password</label>
            <input type="password" id="password" placeholder="Password" required value={registerData.password} onChange={handleChange} />

            <label htmlFor="firstName">First Name</label>
            <input type="text" id="firstName" placeholder="First Name" required value={registerData.firstName} onChange={handleChange} />

            <label htmlFor="lastName">Last Name</label>
            <input type="text" id="lastName" placeholder="Last Name" required value={registerData.lastName} onChange={handleChange} />

            <label htmlFor="dateOfBirth">Date of Birth</label>
            <input type="date" id="dateOfBirth" required value={registerData.dateOfBirth} onChange={handleChange} />

            <select id="privacy" value={registerData.privacy} onChange={handleChange}>
                <option value="public">public</option>
                <option value="private">private</option>
                <option value="almostPrivate">almost private</option>
            </select>

            <label htmlFor="avatar">Avatar</label>
            <input type="file" id="avatar" accept="image/*" onChange={handleFileChange} />

            <label htmlFor="nickname">Nickname</label>
            <input type="text" id="nickname" placeholder="Nickname" value={registerData.nickname} onChange={handleChange} />

            <label htmlFor="aboutMe">About Me</label>
            <textarea id="aboutMe" placeholder="Tell us about yourself" value={registerData.aboutMe} onChange={handleChange} />

            <button type="submit">Register</button>
        </form>
    );
}

export default Register
