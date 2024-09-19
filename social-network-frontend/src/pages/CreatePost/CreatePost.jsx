import React, { useState } from 'react'

import { useCreatePost } from '../../api'

const CreatePost = ({ setRefreshTrigger, groupId }) => {

    const { postData, handleChange, handleFileChange, handleSubmit } = useCreatePost(groupId)

    const submit = (e) => {
        setRefreshTrigger(prev => !prev)
        handleSubmit(e)
    }


    return (
        <form onSubmit={submit} method='POST' style={{ display: 'flex', flexDirection: 'column', maxWidth: '400px', margin: '0 auto' }}>

            <label htmlFor="title">Title</label>
            <input type="text" id="title" placeholder='Title' required value={postData.title} onChange={handleChange} />

            <label htmlFor="content">Content</label>
            <input type="text" id="content" placeholder='Content' required value={postData.content} onChange={handleChange} />

            <label htmlFor="avatar">picture</label>
            <input type="file" id="avatar" accept='image/*' onChange={handleFileChange} />

            <label htmlFor="privacy">Choose post privacy</label>
            <select id="privacy" value={postData.privacy} onChange={handleChange}>
                <option value="public">public</option>
                <option value="private">priate</option>
                {/* <option value="almostPrivate">almost private</option> */}
            </select>

            <button type='submit'>Create post</button>
        </form>
    )
}

export default CreatePost
