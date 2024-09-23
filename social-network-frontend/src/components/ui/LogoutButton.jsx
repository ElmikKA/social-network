import React from 'react'
import { useLogOut } from '../../api'

const LogoutButton = () => {

    const handleLogOut = useLogOut()

    return (
        <button onClick={handleLogOut} className='logout-button'>Logout</button>
    )
}

export default LogoutButton
