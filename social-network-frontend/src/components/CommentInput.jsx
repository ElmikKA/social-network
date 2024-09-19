import React from 'react'
import { useSetComment } from '../api'

const CommentInput = ({ refreshComments, postId }) => {

    const { commentData, handleFileChange, handleChange, handleSubmit } = useSetComment(postId, refreshComments)

    const submit = (e) => {
        handleSubmit(e)

    }

    return (
        <form onSubmit={submit} method='POST' className='commentInput'>

            <input type="text" id='content' placeholder='comment...' required value={commentData.content} onChange={handleChange} />

            <input type="file" id="avatar" accept='image/*' onChange={handleFileChange} />

            <button type='submit'>Comment</button>
        </form>
    )
}

export default CommentInput
