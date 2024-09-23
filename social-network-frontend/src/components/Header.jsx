import { useNavigate } from "react-router-dom";
import LogoutButton from "./ui/LogoutButton";
import logo from '../assets/vibley-logo.png';
import searchIcon from '../assets/search.png';
import notificationBell from '../assets/notification-bell.png';
import chatLine from '../assets/chat-line.png';
import { useState } from "react";
import Notifications from '../components/Notifications'

function Header({ onToggleMessenger }) {
    const navigate = useNavigate()

    const [isNotificationOpen, setNotificationOpen] = useState(false);
    const [isLogoutOpen, setLogoutOpen] = useState(false);

    const toggleNotificationDropdown = () => {
        setNotificationOpen(!isNotificationOpen);
        if(!isNotificationOpen) {
            setLogoutOpen(false)
        }
    };

    const toggleLogoutDropdown = () => {
        setLogoutOpen(!isLogoutOpen);
        if(!isLogoutOpen) {
            setNotificationOpen(false);
        }
    };

    return (
        <header className="header">
            <div className="logo-and-search-bar">
                <img src={logo} alt="vibley-logo" style={{cursor: 'pointer'}}  onClick={() => navigate('/home')}/>
                <div class="search-container">
                    <img src={searchIcon} alt="search-icon" className="search-icon" />
                    <input type="text" className="search-input" placeholder="Start typing to search..." />
                </div>
            </div>

            <div className="notification-message-profile-buttons">
                <img src={notificationBell} alt="" style={{cursor: 'pointer'}} onClick={toggleNotificationDropdown}/>
                {isNotificationOpen && (
                    <div className="dropdown-menu">
                        <Notifications></Notifications>
                    </div>
                )}

                <img src={chatLine} alt="" style={{cursor: 'pointer'}} onClick={onToggleMessenger}/>

                <div className="header-profile-picture" onClick={toggleLogoutDropdown} style={{width: '30px', height: '30px', background: 'white', borderRadius: '50px', cursor: 'pointer'}}>
                    {isLogoutOpen && (
                        <div className="logout-dropdown-menu">
                            <LogoutButton></LogoutButton>
                        </div>
                    )}
                </div>
            </div>
        </header>
    );

}

export default Header