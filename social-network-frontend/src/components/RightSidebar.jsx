import React from 'react'
import { useGetContacts } from '../api'
import ContactElement from './ContactElement'
import GroupElement from './GroupElement'

const RightSidebar = ({isOpen}) => {

    const { contacts, loading } = useGetContacts()

    if (loading) {
        return <div>Loading...</div>
    }

    // add buttons to open message box
    // make the contacts and groups into elements

    return (
        <div className={`rightSidebar ${isOpen ? 'open' : ''}`}>
            <div className='right-sidebar-inner-div'>
                <div className='contactDiv'>
                    <ContactElement contacts={contacts.contacts} />
                </div>
                <div className='groupChatDiv'>
                    <GroupElement groupChat={contacts.groupChats} />
                </div>
            </div>
        </div >
    )
}

export default RightSidebar