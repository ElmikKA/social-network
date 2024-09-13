import React, { useState } from 'react'

const Event = () => {

    const groupId = 1
    const [eventData, setEventData] = useState({
        groupId: groupId,
        title: '',
        description: '',
        time: ''
    })

    const handleChange = (e) => {

        const { id, value } = e.target
        setEventData(prevState => ({
            ...prevState,
            [id]: value,
        }))
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        const requestOptions = {
            method: 'POST',
            credentials: 'include',
            header: { 'Content-Type': 'application/json' },
            body: JSON.stringify(eventData),
        }

        try {
            const response = await fetch('http://localhost:8080/api/createEvent', requestOptions)
            const data = await response.json()
            console.log(data)
        } catch (err) {
            console.log(err)
        }
    }

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
