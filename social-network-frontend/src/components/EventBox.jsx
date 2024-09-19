import React from 'react'
import RespondNotificationButton from './ui/RespondNotificationButton';

const EventBox = ({ setRefreshTrigger, events }) => {
    const formatDate = (dateString) => {
        const date = new Date(dateString);
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        return `${year}-${month}-${day}`;
    };

    return (
        <div className='groupEventDiv'>
            {events && events.length > 0 ? (
                events.map((event) => (
                    <div className='groupEventBox' key={event.Id}>
                        <p>title:{event.Title}</p>
                        <p>when:{formatDate(event.Time)}</p>
                        {event.Status === 'pending' ?
                            <>
                                <RespondNotificationButton setRefreshTrigger={setRefreshTrigger} idRef={event.IdRef} type={"e_ref"} response="completed" ></RespondNotificationButton >
                                <RespondNotificationButton setRefreshTrigger={setRefreshTrigger} idRef={event.IdRef} type={"e_ref"} response="rejected" ></RespondNotificationButton >
                            </>
                            : <div>{event.Status}</div>}
                    </div>
                ))
            ) : (<p>No evente</p>)}
        </div>
    )
}

export default EventBox
