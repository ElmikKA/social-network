import React, { useEffect } from 'react'
import { useRespondNotification } from '../services/api'

const RespondNotification = () => {
    const idRef = 2
    const type = 'e_ref' // f_ref, g_ref, e_ref
    const response = 'completed' // "completed" / "rejected"

    // g_ref to respond to group join request
    // f_ref to respond to follow request
    // e_ref to respond to event

    useRespondNotification(idRef, type, response)

    return (
        <div>


        </div>
    )
}

export default RespondNotification
