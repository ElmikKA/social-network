import React from 'react'
import { useNavigate } from 'react-router-dom'

const GroupElement = (groupChats) => {
    const navigate = useNavigate()
    return (
        <>
            <p>GroupChats:</p>
            {groupChats?.groupChat?.length > 0 ? (
                <div>
                    {groupChats.groupChat.map((groupChat) => (
                        <button key={groupChat.GroupId} className='groupChatList' onClick={() => navigate(`/groupMessage/${groupChat.GroupId}`)}>
                            <img
                                src={'http://localhost:8080/api/avatars/db/static/default.png'}
                                style={{ width: '50px', height: '50px', borderRadius: '50%' }}
                            />
                            {groupChat.Title}
                        </button>
                    ))}
                </div>
            ) : (<p>No groups available</p>)}
        </>
    )
}

export default GroupElement
