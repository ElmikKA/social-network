import React, { useEffect } from 'react'
import { useRespondFollow, useAddFollow } from '../services/api'

const Follow = () => {

    useRespondFollow(1, "completed")
    // useAddFollow(2)
    return (
        <div>
            accept/reject follow
        </div>
    )
}

export default Follow
