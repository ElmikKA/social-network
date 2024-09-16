import React from 'react'

const GroupElement = (groupChats) => {
    return (
        <>
            <p>GroupChats:</p>
            {groupChats?.groupChat?.length > 0 ? (
                <div>
                    {groupChats.groupChat.map((groupChat) => (
                        <div key={groupChat.GroupId} className='groupChatList'>
                            <img
                                src={'http://localhost:8080/api/avatars/db/static/default.png'}
                                style={{ width: '50px', height: '50px', borderRadius: '50%' }}
                            />
                            {groupChat.Title}
                        </div>
                    ))}
                </div>
            ) : (<p>No groups available</p>)}
        </>
    )
}

export default GroupElement
