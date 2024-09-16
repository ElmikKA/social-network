import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
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

createRoot(document.getElementById('root')).render(
  <StrictMode>

    {/* <Register /> */}
    {/* <hr /> */}
    {/* <Root></Root> */}
    {/* <hr /> */}
    {/* <Profile /> */}
    {/* <hr /> */}
    {/* <Home /> */}
    {/* <hr /> */}
    {/* <Comment /> */}
    {/* <hr /> */}
    {/* <CreatePost /> */}
    {/* <Follow /> */}
    {/* <Groups /> */}

    <Router>
      <Routes>
          <Route path='/' element={<Login />} />
          <Route path='/register' element={<Register />} />
      </Routes>
    </Router>


  </StrictMode>
)
