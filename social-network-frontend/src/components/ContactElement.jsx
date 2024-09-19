import React from 'react'
import { useNavigate } from 'react-router-dom'

const ContactElement = ({ contacts }) => {
    const navigate = useNavigate()
    return (
        <>
            <p style={{fontWeight: '600'}}>Contacts:</p>
            {contacts?.length > 0 ? (
                <div>
                    {contacts.map((contact) => (
                        <div key={contact.Id}
                            className='contactChatList'
                            onClick={() => navigate(`/message/${contact.Id}`)} >
                            <img
                                src={`http://localhost:8080/api/avatars/${contact.Avatar ? contact.Avatar : '/db/static/default.webp'}`}
                                alt={`${contact.Name}'s Avatar`}
                                style={{ width: '50px', height: '50px', borderRadius: '50%' }}
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
