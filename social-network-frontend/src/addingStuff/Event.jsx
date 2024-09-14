import React, { useState } from 'react'
import { useCreateEvent } from '../services/api'

const Event = () => {
    const groupId = 1
    const { eventData, handleChange, handleSubmit } = useCreateEvent(groupId)

    return (
        <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', maxWidth: '400px', margin: '0 auto' }}>
            create event
            < label htmlFor="title" > Title</label >
            <input type="text" id="title" value={eventData.title} onChange={handleChange} required />

            <label htmlFor="description">description</label>
            <input type="text" id="description" value={eventData.description} onChange={handleChange} required />

            <label htmlFor="time">Date of Birth</label>
            <input type="date" id="time" required value={eventData.time} onChange={handleChange} />

            <button type='submit'>create event</button>
        </form>
    )
}

export default Event
