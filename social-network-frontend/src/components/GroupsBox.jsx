import React from 'react'
import { useNavigate } from 'react-router-dom'

const GroupsBox = ({ groups }) => {
    const navigate = useNavigate()
    console.log(groups)

    return (
        <div>
            {groups ? (
                groups.map((group) => (
                    <div key={group.id} className='groupBox' onClick={() => navigate(`/group/${group.id}`)}>
                        {group.title}
                        <img src={`http://localhost:8080/api/avatars/db/assets/default.png`}
                            alt="group picture"
                            style={{ width: '50px', height: '50px', borderRadius: '50%' }} />
                    </div>
                ))
            ) : <p>No posts</p>}
        </div>
    )
}

export default GroupsBox
