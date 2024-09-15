import React, { useEffect } from 'react'
import { useAddFollow } from '../services/api'

const Follow = () => {

    const userId = 2
    useAddFollow(userId)
    return (
        <div>
            send follow request
        </div>
    )
}

export default Follow
