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

createRoot(document.getElementById('root')).render(
  // <StrictMode>
  <div>


    {/* <Register /> */}
    {/* <hr /> */}
    < Login />
    <hr />
    {/* <Profile /> */}
    {/* <hr /> */}
    {/* <Home /> */}
    {/* <hr /> */}
    {/* <Comment /> */}
    {/* <hr /> */}
    {/* <CreatePost /> */}
    {/* <Follow /> */}
    {/* < Groups /> */}
    {/* <JoinGroup /> */}
    {/* < Event /> */}
    {/* <RespondNotification /> */}

  </ div>

  // </StrictMode >
)
