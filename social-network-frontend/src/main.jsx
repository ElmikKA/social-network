import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import Root from './Root'
import './index.css'

import Register from './pages/Register/Register'
import Login from './pages/Login/Login'
import Profile from './pages/profile/Profile'
import Home from './pages/HomePage/Home'
import CreatePost from './pages/CreatePost/CreatePost'
import Comment from './components/Comment'
import Follow from './addingStuff/Follow'
import Groups from './addingStuff/Groups'
import JoinGroup from './addingStuff/JoinGroup'
import Event from './addingStuff/Event'
import RespondNotification from './addingStuff/RespondNotification'
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

const Layout = () => {
  return (
    <div className='page'>
      <Header />
      <div className='main' style={{ display: "flex" }}>
        <LeftSidebar />
        <Outlet />
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
      }
    ]
  },
]);

createRoot(document.getElementById('root')).render(
  <RouterProvider router={router} />
)
