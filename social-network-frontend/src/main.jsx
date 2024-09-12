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

createRoot(document.getElementById('root')).render(
  <StrictMode>

    {/* <Register /> */}
    {/* <hr /> */}
    <Login />
    <hr />
    <Profile />
    <hr />
    {/* <Home /> */}
    {/* <hr /> */}
    <Comment />
    {/* <hr /> */}
    {/* <CreatePost /> */}

  </StrictMode>,
)
