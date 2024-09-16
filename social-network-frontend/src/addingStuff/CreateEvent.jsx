import React from 'react'
import { useCreateEvent } from '../services/api'

const CreateEvent = () => {
    const groupId = 1
    const { eventData, handleChange, handleSubmit } = useCreateEvent(groupId)

    return (
        <div className='createEventDiv'>
            <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', maxWidth: '400px', margin: '0 auto' }}>
                create event
                < label htmlFor="title" > Title</label >
                <input type="text" id="title" value={eventData.title} onChange={handleChange} required />

                <label htmlFor="description">description</label>
                <input type="text" id="description" value={eventData.description} onChange={handleChange} required />

                <label htmlFor="time">Date of Birth</label>
                <input type="date" id="time" required value={eventData.time} onChange={handleChange} />

                <button type='submit'>create event</button>
            </form></div>
    )
}

export default CreateEvent