import React from 'react'
import { useNavigate, useLocation } from 'react-router-dom'
import { useCheckLoggedIn } from '../api'
import homePageButton from '../assets/home-page-button.png'
import groups from '../assets/groups.png';
import userProfile from '../assets/user-profile.png'
import users from '../assets/users.png'


const LeftSidebar = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const { userData, loading } = useCheckLoggedIn();

    if (loading) {
        return <div>Loading...</div>
    }
    return (
        <div className='left-sidebar'>
            <div className='right-sidebar-main-pages-div'>

                <h5 style={{color: '#7E7E7E', marginBottom: '10px'}}>Main Pages</h5>
                <div className={`left-sidebar-button-div ${location.pathname === '/home' ? 'active' : ''}`} onClick={() => navigate('/home')}>
                    <img src={homePageButton} alt="" />
                    <h4>Home Page</h4>
                </div>

                <div className={`left-sidebar-button-div ${location.pathname === `/user/${userData.userId}` ? 'active' : ''}`} onClick={() => navigate(`/user/${userData.userId}`)}>
                    <img src={userProfile} alt="" />
                    <h4>User Profile</h4>
                </div>

                <div className={`left-sidebar-button-div ${location.pathname === '/users' ? 'active' : ''}`} onClick={() => navigate('/users')}>
                    <img src={users} alt="" />
                    <h4>Users</h4>
                </div>

                <div className={`left-sidebar-button-div ${location.pathname === '/groups' ? 'active' : ''}`} onClick={() => navigate('/groups')}>
                    <img src={groups} alt="" />
                    <h4>Groups</h4>
                </div>
            </div>

            <div className='right-sidebar-main-pages-div'>
                <h5 style={{color: '#7E7E7E', marginBottom: '10px'}}>Create Pages</h5>

                <div className={`left-sidebar-button-div ${location.pathname === '/createGroup' ? 'active' : ''}`} onClick={() => navigate('/createGroup')}>
                    <img src={homePageButton} alt="" />
                    <h4>Create a group</h4>              
                </div>

            </div>
        </div>
    )
}

export default LeftSidebar