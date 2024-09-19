import React from 'react'
import { useCreateEvent } from '../api'

const CreateEvent = ({ setRefreshTrigger, groupId }) => {
    const { eventData, handleChange, handleSubmit } = useCreateEvent(groupId)

    const submit = (e) => {
        handleSubmit(e)
        setRefreshTrigger(post => !post)
    }

    return (
        <div className='createEventDiv'>
            <form onSubmit={submit} style={{ display: 'flex', flexDirection: 'column', maxWidth: '400px', margin: '0 auto' }}>
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