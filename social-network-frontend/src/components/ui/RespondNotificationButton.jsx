import React from 'react'
import { useRespondNotification } from '../../api'

const RespondNotificationButton = ({ setRefreshTrigger, idRef, type, response }) => {

    const respondNotification = useRespondNotification()

    const handleClick = () => {
        respondNotification(idRef, type, response)
        setRefreshTrigger(prev => !prev)
        return <div>response sent</div>
    }

    return (
        <button onClick={handleClick}>
            {response === "completed" && <p>accept</p>}
            {response === "rejected" && <p>reject</p>}
        </button>
    )
}

export default RespondNotificationButton
