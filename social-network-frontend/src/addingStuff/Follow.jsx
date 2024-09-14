import React, { useEffect } from 'react'
import { useAddFollow } from '../services/api'

const Follow = () => {

    useAddFollow(2)
    return (
        <div>
            send follow request
        </div>
    )
}

export default Follow
