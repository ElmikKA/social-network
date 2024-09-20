import React from 'react'
import { useGetAllUsers } from '../../api'
import UserBox from '../../components/UserBox'

const Users = () => {

    const { allUsers, loading } = useGetAllUsers()
    if (loading) return <p>Loading...</p>


    return (
        <div className='users-main'>
            <UserBox users={allUsers} />
        </div>
    )
}

export default Users
