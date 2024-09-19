import React, { useState } from 'react'
import InviteGroup from '../InviteGroup'
import { useGetGroupInviteUsers } from '../../api'

const ToggleInviteGroup = ({ groupId }) => {

    const { users, loading, fetchUsers } = useGetGroupInviteUsers(groupId)

    const [opened, setOpened] = useState(false)
    const handleClick = () => {
        setOpened(prev => !prev)
        if (!opened) {
            fetchUsers()
        }
    }
    if (loading) return <p>Loading...</p>
    return (
        <div className='toggleInviteGroup'>
            <button onClick={handleClick}>invite people</button>
            {
                opened && (


                    <>
                        <p>opened</p>
                        <InviteGroup users={users} groupId={groupId} />
                    </>
                )
            }
        </div>
    )
}

export default ToggleInviteGroup
