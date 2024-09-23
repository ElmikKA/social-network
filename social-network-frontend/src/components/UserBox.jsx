import React from 'react'
import { useNavigate } from 'react-router-dom'

const UserBox = (users) => {
    console.log(users.users)
    const navigate = useNavigate()
    return (
        <div className='users-profile-grid'>
            {users.users.map((user) => (
                <div key={user.id} className='user-profile-items' onClick={() => navigate(`/user/${user.id}`)}>
                    <img
                        src={`http://localhost:8080/api/avatars/${user.avatar ? user.avatar : '/db/assets/default.webp'}`}
                        alt={`${user.name}'s Avatar`}
                         />
                    <h2>{user.name}</h2>
                </div>
            ))
            }
        </div >
    )
}

export default UserBox
