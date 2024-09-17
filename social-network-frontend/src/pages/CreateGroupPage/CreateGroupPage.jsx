import React from 'react'
import { useCreateGroup } from '../../api'
import { useOutletContext } from 'react-router-dom'

const CreateGroupPage = () => {
    const { onContactCreated } = useOutletContext()

    const { groupData, handleChange, handleSubmit } = useCreateGroup(onContactCreated)



    return (
        <form onSubmit={handleSubmit} className='createGroupPage'>

            <label htmlFor="title">Titl</label>
            <input type="text" id="title" required value={groupData.title} onChange={handleChange} />

            <label htmlFor="description">description</label>
            <input type="text" id="description" value={groupData.description} onChange={handleChange} required />

            <button type='submit'>create group</button>

        </form>
    )
}

export default CreateGroupPage
