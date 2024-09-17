import React from 'react'
import { UnFollow } from '../../api'
import { useOutletContext } from 'react-router-dom'

const UnFollowButton = ({ userId, setRefreshTrigger }) => {
    const { onContactCreated } = useOutletContext()

    const unFollow = () => {
        UnFollow(userId, onContactCreated)
        console.log("unfollowing")
        setRefreshTrigger(prev => !prev)
    }

    return (
        <button className='followButton' onClick={unFollow}>unfollow</button>
    )
}

export default UnFollowButton
