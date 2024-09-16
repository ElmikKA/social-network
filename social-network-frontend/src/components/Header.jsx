import { useNavigate } from "react-router-dom";
import LogoutButton from "./ui/LogoutButton";

function Header() {
    const navigate = useNavigate()

    return (
        <header className="navbar">
            <nav>
                <button className='leftSidebarButton' onClick={() => navigate('/home')} >Home Page</button>
                home logo, search,notifications, logout
                <LogoutButton />
            </nav>
            <hr />
        </header>
    );

}

export default Header