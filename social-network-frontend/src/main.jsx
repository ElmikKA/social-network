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
import { useEffect } from 'react'
import { InitSocket } from './WebSocket'

const Layout = () => {
  useEffect(() => {
    InitSocket()
  }, [])
  return (
    <div className='page'>
      <Header />
      <div className='main' style={{ display: "flex" }}>
        <LeftSidebar />
        <div className='outletBody'>
          <Outlet />
        </div>
        <Notifications />
        <RightSidebar />
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
        element: <p>group message</p>
      }
    ]
  },
]);

createRoot(document.getElementById('root')).render(
  <RouterProvider router={router} />
)
