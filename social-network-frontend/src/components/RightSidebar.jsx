import React from 'react'
import { useGetContacts } from '../api'
import ContactElement from './ContactElement'
import GroupElement from './GroupElement'

const RightSidebar = () => {

    const { contacts, loading } = useGetContacts()

    if (loading) {
        return <div>Loading...</div>
    }

    // add buttons to open message box
    // make the contacts and groups into elements

    return (
        <div className='rightSidebar'>
            <div className='contactDiv'>
                <ContactElement contacts={contacts.contacts} />
            </div>
            <div className='groupChatDiv'>
                <GroupElement groupChat={contacts.groupChats} />
            </div>
        </div >
    )
}

export default RightSidebar
