import React from 'react'
import { useGetContacts } from '../services/api'

const RightSidebar = () => {

    const { contacts, loading } = useGetContacts()
    if (loading) {
        return <div>Loading...</div>
    }

    if (contacts) {

    }

    return (
        <div className='rightSidebar'>

            contacts, groupchats
        </div>
    )
}

export default RightSidebar
