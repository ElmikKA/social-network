import React, { useEffect, useState } from 'react'

const Groups = () => {

    const [groupData, setGroupData] = useState({
        title: '',
        description: ''
    })


    const handleChange = (e) => {
        const { id, value } = e.target
        setGroupData(prevState => ({
            ...prevState,
            [id]: value
        }))
    }
    const handleSubmit = (e) => {
        e.preventDefault()
        const AddGroup = async () => {
            try {
                data = await fetch('http://localhost:8080/api/createGroup', {
                    method: "POST",
                    credentials: 'include',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(groupData)
                })
                console.log(data)
            } catch (err) {
                console.log(err)
                // status 409 conflict means group already exists
            }
        }
        AddGroup()
    }


    return (
        <div style={{ display: 'flex', flexDirection: 'column', maxWidth: '400px', margin: '0 auto' }}>
            <label htmlFor="title">Titl</label>
            <input type="text" id="title" required value={groupData.title} onChange={handleChange} />

            <label htmlFor="description">description</label>
            <input type="text" id="description" value={groupData.description} onChange={handleChange} />

            <button onClick={handleSubmit}>create group</button>

        </div>
    )
}

export default Groups
