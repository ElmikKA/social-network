import React from 'react'
import { useSendGroupInvite } from '../../api'

const GroupInviteUser = ({ userId, groupId }) => {
    const { sendInvite } = useSendGroupInvite()

    const handleClick = () => {
        console.log("sending invite", userId, groupId)
        sendInvite(groupId, userId)
    }


    return (
        <button onClick={handleClick}>
            {userId}
            {groupId}

        </button>
    )
}

export default GroupInviteUser
