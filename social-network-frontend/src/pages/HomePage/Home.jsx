import React, { useEffect, useState } from 'react'
import { useGetAllPosts } from '../../api'
import CreatePost from '../CreatePost/CreatePost'
import PostBox from '../../components/PostBox'

const Home = () => {

    const [refreshTrigger, setRefreshTrigger] = useState(false)
    const { allPosts } = useGetAllPosts(refreshTrigger)

    return (
        <div className='homePage'>
            <div className='createPost'>
                <CreatePost setRefreshTrigger={setRefreshTrigger} groupId={0} />
            </div>
            <div className='homePosts'>
                <PostBox allPosts={allPosts} />
            </div>
        </div>
    )
}

export default Home
