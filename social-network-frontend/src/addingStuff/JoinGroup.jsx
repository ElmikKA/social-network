import React, { useEffect } from 'react'
import { useSendGroupJoinRequest } from '../services/api'

const JoinGroup = () => {

    const groupId = 1
    useSendGroupJoinRequest(groupId)

    return (
        <div>
            sending group join request
        </div>
    )
}

export default JoinGroup
