import React from 'react'
import { useGetNotifications } from '../services/api'
import RespondNotificationButton from './ui/RespondNotificationButton'

const Notifications = () => {

    const { notificationData, loading } = useGetNotifications()
    if (loading) {
        return <div>Loading...</div>
    }




    return (
        <div className='notifications'>
            {notificationData?.notifications?.length > 0 ? (
                <ul>
                    {notificationData.notifications.map((notification) => (
                        <li key={notification.id}>
                            {notification.content}
                            <RespondNotificationButton idRef={notification.idRef} type={notification.type} response="completed" ></RespondNotificationButton >
                            <RespondNotificationButton idRef={notification.idRef} type={notification.type} response="rejected" ></RespondNotificationButton >

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
