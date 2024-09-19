import React, { useEffect } from 'react'
import ContactElement from './ContactElement'
import GroupElement from './GroupElement'

const RightSidebar = ({ contacts }) => {


    useEffect(() => {


    }, [contacts])

    // add buttons to open message box
    // make the contacts and groups into elements

    return (
        <div className='rightSidebar'>
            <div className='contactDiv'>
                <p>Contacts:</p>
                {contacts?.contacts?.length > 0 ? (
                    <ContactElement contacts={contacts.contacts} />
                ) : (
                    <p>No contacts available</p>
                )}
            </div>
            <div className='groupChatDiv'>
                <p>GroupChats:</p>
                {contacts?.groupChats?.length > 0 ? (
                    <GroupElement groupChat={contacts.groupChats} />
                ) : (
                    <p>no groupchat available</p>
                )}
            </div>
        </div >
    )
}

export default RightSidebar
