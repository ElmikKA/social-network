import React from 'react'
import { useRespondNotification } from '../../api'
import { useOutletContext } from 'react-router-dom'

const RespondNotificationButton = ({ refreshSidebar, setRefreshTrigger, idRef, type, response }) => {


    const respondNotification = useRespondNotification(refreshSidebar)

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
