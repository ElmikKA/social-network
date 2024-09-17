import React from 'react'
import { useNavigate } from 'react-router-dom'

const UserBox = (users) => {
    console.log(users.users)
    const navigate = useNavigate()
    return (
        <div>
            {users.users.map((user) => (
                <div key={user.id} className='userBox' onClick={() => navigate(`/user/${user.id}`)}>
                    <img
                        src={`http://localhost:8080/api/avatars/${user.avatar ? user.avatar : '/db/assets/default.webp'}`}
                        alt={`${user.name}'s Avatar`}
                        style={{ width: '50px', height: '50px', borderRadius: '50%' }} />
                    {user.name}
                </div>
            ))
            }
        </div >
    )
}

export default UserBox
