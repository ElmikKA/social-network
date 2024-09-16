import React from 'react'
import { useGetAllUsers } from '../../services/api'
import UserBox from '../../components/UserBox'

const Users = () => {

    const { allUsers, loading } = useGetAllUsers()
    if (loading) return <p>Loading...</p>
    console.log(allUsers)



    return (
        <div className='usersMain'>
            <UserBox users={allUsers} />
        </div>
    )
}

export default Users
