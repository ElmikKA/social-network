import React, { useEffect, useState } from 'react'
import { useCreateGroup } from '../services/api'

const Groups = () => {

    const { groupData, handleChange, handleSubmit } = useCreateGroup()

    return (
        <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', maxWidth: '400px', margin: '0 auto' }}>
            <label htmlFor="title">Titl</label>
            <input type="text" id="title" required value={groupData.title} onChange={handleChange} />

            <label htmlFor="description">description</label>
            <input type="text" id="description" value={groupData.description} onChange={handleChange} required />

            <button type='submit'>create group</button>

        </form>
    )
}

export default Groups
