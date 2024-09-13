import React, { useEffect } from 'react'
import { RespondFollow } from '../services/api'

const Follow = () => {

    RespondFollow(2, "completed")
    return (
        <div>
            accept/reject follow

        </div>
    )
}

export default Follow
