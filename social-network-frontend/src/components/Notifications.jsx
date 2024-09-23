import React, { useEffect, useState } from 'react';

import RespondNotificationButton from './ui/RespondNotificationButton';
import { useGetNotifications, useGetUser } from '../api';

const Notifications = ({ refreshSidebar }) => {
    const [refreshTrigger, setRefreshTrigger] = useState(false);
    const [previousNotifications, setPreviousNotifications] = useState(null);

    const { notificationData, loading } = useGetNotifications(refreshTrigger);

    useEffect(() => {
        const intervalId = setInterval(() => {
            if (notificationData && JSON.stringify(notificationData) !== JSON.stringify(previousNotifications)) {
                setPreviousNotifications(notificationData);
                setRefreshTrigger(prev => !prev);
            }
        }, 5000);

        return () => clearInterval(intervalId);
    }, [notificationData, previousNotifications]);

    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <div className='notification-div'>
            {notificationData?.notifications?.length > 0 ? (
                <ul>
                    {notificationData.notifications.map((notification) => (
                        <NotificationItem
                            key={notification.id}
                            notification={notification}
                            refreshSidebar={refreshSidebar}
                            setRefreshTrigger={setRefreshTrigger}
                        />
                    ))}
                </ul>
            ) : (
                <p>No new notifications</p>
            )}
        </div>
    );
};

export default Notifications

const NotificationItem = ({ notification, refreshSidebar, setRefreshTrigger }) => {
    // const { userData, loading: userLoading, error } = useGetUser(notification.idRef, setRefreshTrigger);
    console.log('id', typeof notification.id)

    // if (userLoading) return <li>Loading user data...</li>;
    // if (error) return <li>Error loading user data...</li>;

    return (
        <li>
            {notification.content} 
            <div className='notification-button-div'>
                <RespondNotificationButton
                    refreshSidebar={refreshSidebar}
                    setRefreshTrigger={setRefreshTrigger}
                    idRef={notification.idRef}
                    type={notification.type}
                    response="completed"
                />
                <RespondNotificationButton
                    refreshSidebar={refreshSidebar}
                    setRefreshTrigger={setRefreshTrigger}
                    idRef={notification.idRef}
                    type={notification.type}
                    response="rejected"
                />
            </div>
        </li>
    );
};
