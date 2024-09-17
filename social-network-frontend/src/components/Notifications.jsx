import React, { useEffect, useState } from 'react'

import RespondNotificationButton from './ui/RespondNotificationButton'
import { useGetNotifications } from '../api'
import { GetSocket } from '../WebSocket'

const Notifications = () => {
    const [refreshTrigger, setRefreshTrigger] = useState(false)

    const { notificationData, loading } = useGetNotifications(refreshTrigger)

    useEffect(() => {
        const intervalId = setInterval(() => {
            setRefreshTrigger(prev => !prev)
        }, 5000);
        return () => clearInterval(intervalId)
    }, [])

    if (loading) {
        return <div>Loading...</div>
    }
    return (
        <div>
            {notificationData?.notifications?.length > 0 ? (
                <ul>
                    {notificationData.notifications.map((notification) => (
                        <li key={notification.id}>
                            {notification.content}
                            <RespondNotificationButton setRefreshTrigger={setRefreshTrigger} idRef={notification.idRef} type={notification.type} response="completed" ></RespondNotificationButton >
                            <RespondNotificationButton setRefreshTrigger={setRefreshTrigger} idRef={notification.idRef} type={notification.type} response="rejected" ></RespondNotificationButton >

                        </li>
                    ))}
                </ul>
            ) : (
                <p>No new notifications</p>
            )}
        </div>
    );
}

export default Notifications
