import React from 'react'
import { useRespondNotification } from '../../api'

const RespondNotificationButton = ({ refreshSidebar, setRefreshTrigger, idRef, type, response }) => {


    const respondNotification = useRespondNotification(refreshSidebar)

    const handleClick = () => {
        respondNotification(idRef, type, response)
        setRefreshTrigger(prev => !prev)
        return <div>response sent</div>
    }

    return (
        <button onClick={handleClick} className='notification-btn'>
            {response === "completed" && <p>Accept</p>}
            {response === "rejected" && <p>Reject</p>}
        </button>
    )
}

export default RespondNotificationButton
