import React, { useState } from 'react'
import { useSetComment } from '../services/api'

const Comment = () => {

    // hardcoded to add comments to postId = 1
    const postId = 1

    const { commentData, handleFileChange, handleChange, handleSubmit } = useSetComment(postId)

    return (
        <form onSubmit={handleSubmit} method='POST' style={{ display: 'flex', flexDirection: 'column', maxWidth: '400px', margin: '0 auto' }}>

            <label htmlFor="content">Content</label>
            <input type="text" id="content" placeholder='Content' required value={commentData.content} onChange={handleChange} />

            <label htmlFor="avatar">picture</label>
            <input type="file" id="avatar" accept='image/*' onChange={handleFileChange} />

            <button type='submit'>Create post</button>
        </form>
    )
}

export default Comment
