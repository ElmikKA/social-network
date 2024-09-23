import React from 'react'
import { useCreateEvent } from '../api'

const CreateEvent = ({ setRefreshTrigger, groupId }) => {
    const { eventData, handleChange, handleSubmit } = useCreateEvent(groupId)

    const submit = (e) => {
        handleSubmit(e)
        setRefreshTrigger(post => !post)
    }

    return (
        <div className='create-event-div'>
            <form onSubmit={submit} style={{width: '100%'}}>
                <h3>Create event</h3>

                <div className='form-group'>
                    <div className='form-field'>
                        < label htmlFor="title" > Title</label >
                        <input type="text" id="title" value={eventData.title} onChange={handleChange} required />
                    </div>

                    <div className='form-field'>
                        <label htmlFor="description">description</label>
                        <input type="text" id="description" value={eventData.description} onChange={handleChange} required />
                    </div>
                </div>

                <div className='form-group'>
                    <div className='form-field'>
                        <label htmlFor="time">Date of Birth</label>
                        <input type="date" id="time" required value={eventData.time} onChange={handleChange} />
                    </div>

                    <div className='form-field'>
                        <button type='submit'>Create event</button>
                    </div>
                </div>
                                
            </form></div>
    )
}

export default CreateEvent