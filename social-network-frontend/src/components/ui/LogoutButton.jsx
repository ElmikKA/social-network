import React from 'react'
import { useLogout } from '../../services/api'

const LogoutButton = () => {
    const handleLogOut = () => {
        useLogout()
    }
    return (
        <button onClick={() => handleLogOut()}>Logout</button>
    )
}

export default LogoutButton
