import React from 'react'
import { useNavigate } from 'react-router-dom'

const GroupElement = (groupChats) => {
    const navigate = useNavigate()
    return (
        <>
            <p style={{fontWeight: '600'}}>GroupChats:</p>
            {groupChats?.groupChat?.length > 0 ? (
                <div className='message-window-group'>
                    {groupChats.groupChat.map((groupChat) => (
                        <div key={groupChat.GroupId} className='group-chat-list' onClick={() => navigate(`/groupMessage/${groupChat.GroupId}`)}>
                            <img
                                src={'http://localhost:8080/api/avatars/db/static/default.png'}
                                style={{ width: '30px', height: '30px', borderRadius: '50%', marginRight: '10px'}}
                            />
                            <p>{groupChat.Title}</p>
                        </div>
                    ))}
                </div>
            ) : (<p>No groups available</p>)}
        </>
    )
}

export default GroupElement
