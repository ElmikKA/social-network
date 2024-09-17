import React from 'react'
import { useAddFollow } from '../../api'
import { useOutletContext } from 'react-router-dom'

const FollowButton = ({ userId, setRefreshTrigger }) => {
    const { onContactCreated } = useOutletContext()
    console.log(userId)
    const { addFollow, isFollowing } = useAddFollow(userId, onContactCreated)

    const follow = () => {
        addFollow()
        setRefreshTrigger(prev => !prev)
    }

    return (
        <button className='followButton' onClick={follow}>
            {isFollowing ? 'sent follow request' : 'follow'}
        </button >
    )
}

export default FollowButton
