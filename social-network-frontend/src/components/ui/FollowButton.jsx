import React from 'react'
import { useAddFollow } from '../../api'

const FollowButton = ({ userId }) => {
    console.log(userId)
    const { addFollow, isFollowing } = useAddFollow(userId)

    return (
        <button className='followButton' onClick={addFollow}>
            {isFollowing ? 'sent follow request' : 'follow'}
        </button >
    )
}

export default FollowButton
