import React from 'react'
import { useNavigate } from 'react-router-dom'

const ContactElement = ({ contacts }) => {
    const navigate = useNavigate()
    return (
        <>
            <p style={{fontWeight: '600'}}>Contacts:</p>
            {contacts?.length > 0  && contacts !== null ? (
                <div className='message-window-group'>
                    {contacts.map((contact) => (
                        <div key={contact.Id}
                            className='group-chat-list'
                            onClick={() => navigate(`/message/${contact.Id}`)} >
                            <img
                                src={`http://localhost:8080/api/avatars/${contact.Avatar ? contact.Avatar : '/db/static/default.webp'}`}
                                alt={`${contact.Name}'s Avatar`}
                                style={{ width: '30px', height: '30px', borderRadius: '50%', marginRight: '10px' }}
                            />
                            {contact.Name}
                        </div>
                    ))}
                </div>
            ) : (<p>No contacts available</p>)}
        </>
    )
}

export default ContactElement
