import React from 'react'
import { useNavigate } from 'react-router-dom'
import { useCheckLoggedIn } from '../services/api'

const LeftSidebar = () => {
    const navigate = useNavigate()
    const userId = 1
    const { userData, loading } = useCheckLoggedIn()

    if (loading) {
        return <div>Loading...</div>
    }
    if (userData) {
        console.log(userData)
    }
    return (
        <div className='leftSidebar'>
            <button className='leftSidebarButton' onClick={() => navigate('/home')} >Home Page</button>
            <button className='leftSidebarButton' onClick={() => navigate(`/user/${userData.userId}`)}>User Profile</button>
            <button className='leftSidebarButton' onClick={() => navigate('/users')}>Users</button>
            <button className='leftSidebarButton' onClick={() => navigate('/groups')}>Groups</button>
        </div>
    )
}

export default LeftSidebar
