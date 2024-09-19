import React from 'react'
import { useLogOut } from '../../api'

const LogoutButton = () => {

    const handleLogOut = useLogOut()

    return (
        <button onClick={handleLogOut}>Logout</button>
    )
}

export default LogoutButton
