import React from 'react'
import { useGetContacts } from '../api'
import ContactElement from './ContactElement'
import GroupElement from './GroupElement'

const RightSidebar = ({isOpen}) => {

    const { contacts, loading } = useGetContacts()

    if (loading) {
        return <div>Loading...</div>
    }

    console.log(contacts)

    // add buttons to open message box
    // make the contacts and groups into elements

    return (
        <div className={`rightSidebar ${isOpen ? 'open' : ''}`}>
            <div className='right-sidebar-inner-div'>
                <div className='contactDiv'>
                    {
                        contacts !== null ? <ContactElement contacts={contacts.contacts}/> : 'somethin wrong'

                    }
                </div>
                <div className='groupChatDiv'>
                    {
                        contacts !== null ? <GroupElement groupChat={contacts.groupChats} /> : 'somethin else'

                    }
                </div>
            </div>
        </div >
    )
}

export default RightSidebar