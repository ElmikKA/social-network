import React from 'react'
import { useCreateGroup } from '../../api'
import { useOutletContext } from 'react-router-dom'

const CreateGroupPage = () => {
    const { onContactCreated } = useOutletContext()

    const { groupData, handleChange, handleSubmit } = useCreateGroup(onContactCreated)



    return (
        <div className='create-group-page-div'>
            <h2>Create a New Group</h2>
            <form onSubmit={handleSubmit} className='create-group-page'>

                <label htmlFor="title">Title</label>
                <input type="text" id="title" required value={groupData.title} onChange={handleChange} />

                <label htmlFor="description">Description</label>
                <input type="text" id="description" value={groupData.description} onChange={handleChange} required />

                <button type='submit'>Create Group</button>

            </form>
        </div>
    )
}

export default CreateGroupPage
