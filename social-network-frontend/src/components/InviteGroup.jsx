import React from 'react'
import GroupInviteUser from './ui/GroupInviteUser'

const InviteGroup = ({ users, groupId }) => {

    return (
        <div className='inviteBox'>
            {groupId}
            invite people
            {users.length > 0 ? (
                <div>
                    {users.map((user) => (
                        <div key={user.id}>
                            <p>{user.name}</p>
                            <GroupInviteUser userId={user.id} groupId={groupId} />
                        </div>
                    ))}
                </div>
            ) : (
                <p>no users</p>
            )
            }
        </div >
    )
}

export default InviteGroup
