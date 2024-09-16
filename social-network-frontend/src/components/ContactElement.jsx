import React from 'react'

const ContactElement = (contacts) => {
    return (
        <>
            <p>Contacts:</p>
            {contacts?.contacts?.length > 0 ? (
                <div>
                    {contacts.contacts.map((contact) => (
                        <div key={contact.Id} className='contactChatList' >
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
