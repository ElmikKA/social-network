import React, { useState } from 'react'

import { useCreatePost } from '../../api'

const CreateGroupPost = ({ setRefreshTrigger, groupId }) => {

    const { postData, handleChange, handleFileChange, handleSubmit } = useCreatePost(groupId)
    const submit = (e) => {
        handleSubmit(e)
        setRefreshTrigger(post => !post)
    }

    return (
        <div className='createGroupPostDiv'>

            <form onSubmit={submit} method='POST' style={{ display: 'flex', flexDirection: 'column', maxWidth: '400px', margin: '0 auto' }}>
                Create post

                <label htmlFor="title">Title</label>
                <input type="text" id="title" placeholder='Title' required value={postData.title} onChange={handleChange} />

                <label htmlFor="content">Content</label>
                <input type="text" id="content" placeholder='Content' required value={postData.content} onChange={handleChange} />

                <label htmlFor="avatar">picture</label>
                <input type="file" id="avatar" accept='image/*' onChange={handleFileChange} />

                <button type='submit'>Create post</button>
            </form>
        </div>
    )
}

export default CreateGroupPost
