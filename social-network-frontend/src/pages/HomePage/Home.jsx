import React, { useEffect, useState } from 'react'
import { useGetAllPosts } from '../../api'
import CreatePost from '../CreatePost/CreatePost'
import PostBox from '../../components/PostBox'
import Notifications from '../../components/Notifications'
import EventBox from '../../components/EventBox'

const Home = () => {

    const [refreshTrigger, setRefreshTrigger] = useState(false)
    const { allPosts } = useGetAllPosts(refreshTrigger)

    return (
        <div className='home-page'>
            <div className='main-content'>
                <div className='create-post'>
                    <CreatePost setRefreshTrigger={setRefreshTrigger} groupId={0} />
                </div>
                <div className='home-posts'>
                    <PostBox allPosts={allPosts} />
                </div>
            </div>
            <div className='notifications-and-events'>
                <div className='notification-content'>
                    <Notifications></Notifications>
                </div>

                <div className='event-content'>
                <EventBox/>
                </div>
            </div>

            
        </div>
    )
}

export default Home
