import React, { useEffect } from 'react'
import { useSendGroupJoinRequest } from '../services/api'

const JoinGroup = () => {

    useSendGroupJoinRequest(1)

    return (
        <div>
            sending group join request
        </div>
    )
}

export default JoinGroup
