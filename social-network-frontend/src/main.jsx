import { createRoot } from 'react-dom/client'
import './index.css'
import Register from './pages/Register/Register'
import Login from './pages/Login/Login'
import Profile from './pages/profile/Profile'
import Home from './pages/HomePage/Home'
import LeftSidebar from './components/LeftSidebar'
import {
  createBrowserRouter,
  Outlet,
  RouterProvider,
} from "react-router-dom";
import Header from './components/Header'
import RightSidebar from './components/RightSidebar'
import Notifications from './components/Notifications'
import Footer from './components/Footer'
import Users from './pages/usersPage/Users'
import GroupsPage from './pages/groupsPage/GroupsPage'
import GroupPage from './pages/GroupPage/GroupPage'
import CreateGroupPage from './pages/CreateGroupPage/CreateGroupPage'
import MessagePage from './pages/MessagePage/MessagePage'
import { useCallback, useEffect, useState } from 'react'
import { InitSocket } from './WebSocket'
import GroupMessagePage from './pages/GroupMessagePage/GroupMessagePage'
import { useGetContacts } from './api'

const Layout = () => {

  const [refreshTrigger, setRefreshTrigger] = useState(false)
  const { contacts: fetchedContacts, loading } = useGetContacts(refreshTrigger);
  const [contacts, setContacts] = useState([]);
  const [isMessengerOpen, setMessengerOpen] = useState(true);

  useEffect(() => {
    InitSocket()
  }, [])

  useEffect(() => {
    if (!loading) {
      setContacts(fetchedContacts);
    }
  }, [fetchedContacts, loading]);

  // Callback function to refresh contacts or groups
  const handleContactCreated = useCallback(() => {
    console.log("group created, refreshing")
    setRefreshTrigger(prev => !prev)
  }, [])

  const toggleMessenger = () => {
      setMessengerOpen(!isMessengerOpen);
  };

  if (loading) return <div>Loading...</div>
  
  return (
    <div className='page'>
      <Header onToggleMessenger={toggleMessenger}/>
      <div className='main' style={{ display: "flex" }}>
        <LeftSidebar />
        <div className='outletBody'>
          <Outlet context={{ onContactCreated: handleContactCreated }} />
        </div>
        {/* <div className='notifications'>
          <Notifications refreshSidebar={handleContactCreated} />
        </div> */}
        <RightSidebar contacts={contacts} isOpen={isMessengerOpen} />
      </div>
      <Footer />
    </div>
  )
}
const router = createBrowserRouter([
  {
    path: '/',
    element: <Login />
  },
  {
    path: '/login',
    element: <Login />
  },
  {
    path: '/register',
    element: <Register />
  },
  {
    path: "/",
    element: <Layout />,
    children: [
      {
        path: '/home',
        element: <Home />
      },
      {
        path: '/user/:id',
        element: <Profile />
      },
      {
        path: '/users',
        element: <Users />
      },
      {
        path: '/groups',
        element: <GroupsPage />
      },
      {
        path: '/group/:id',
        element: <GroupPage />
      },
      {
        path: '/createGroup',
        element: <CreateGroupPage />
      },
      {
        path: '/message/:id',
        element: <MessagePage />
      },
      {
        path: '/groupMessage/:id',
        element: <GroupMessagePage />
      }
    ]
  },
]);
createRoot(document.getElementById('root')).render(
  <RouterProvider router={router} />
)