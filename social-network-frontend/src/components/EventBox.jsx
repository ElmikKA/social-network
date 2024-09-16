import React from 'react'

const EventBox = ({ events }) => {
    console.log(events)
    return (
        <div className='groupEventDiv'>
            {events && events.length > 0 ? (

                events.map((event) => (
                    <div className='groupEventBox' key={event.Id}>
                        <p>title:{event.Title}</p>
                        <p>when:{event.Time}</p>
                    </div>
                ))
            ) : (<p>No evente</p>)}
        </div>
    )
}

export default EventBox
