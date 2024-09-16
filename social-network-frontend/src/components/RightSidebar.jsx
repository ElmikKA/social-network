import React from 'react'
import { useGetContacts } from '../services/api'

const RightSidebar = () => {

    const { contacts, loading } = useGetContacts()

    if (loading) {
        return <div>Loading...</div>
    }

    // add buttons to open message box

    return (
        <div className='rightSidebar'>
            <div className='contactDiv'>
                <p>Contacts:</p>
                {contacts?.contacts?.length > 0 ? (
                    <ul>
                        {contacts.contacts.map((contact) => (
                            <li key={contact.Id} className='contactChatList' >
                                {contact.Name}
                                {contact.Avatar ? (
                                    <img
                                        src={`http://localhost:8080/api/avatars/${contact.Avatar}`}
                                        alt={`${contact.Name}'s Avatar`}
                                        style={{ width: '50px', height: '50px', borderRadius: '50%' }}
                                    />
                                ) : (
                                    <p>No avatar available</p>
                                )}
                            </li>
                        ))}
                    </ul>
                ) : (
                    <p>No contacts available</p>
                )}
            </div>
            <div className='groupChatDiv'>
                <p>GroupChats:</p>
                {contacts?.groupChats?.length > 0 ? (
                    <ul>
                        {contacts.groupChats.map((groupChat) => (
                            <li key={groupChat.GroupId} className='groupChatList'>
                                {groupChat.Title}
                            </li>
                        ))}
                    </ul>

                ) : (
                    <p>no groupchat available</p>
                )}
            </div>
        </div>
    )
}

export default RightSidebar
