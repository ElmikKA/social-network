import React, { useState } from 'react'
import { useChangePrivacy } from '../api'

const ChangePrivacy = ({ status }) => {
    const [privacy, setPrivacy] = useState(status || 'public')

    const changePrivacy = useChangePrivacy()

    const handleSubmit = (e) => {
        e.preventDefault()
        changePrivacy(privacy)
    }
    const handleChange = (e) => {
        setPrivacy(e.target.value)
    }

    return (
        <form onSubmit={handleSubmit}>
            <select id="privacy" value={privacy} onChange={handleChange}>
                <option value="public">public</option>
                <option value="private">private</option>
            </select>
            <button type='submit'>change privacy</button>
        </form>

    )
}

export default ChangePrivacy
