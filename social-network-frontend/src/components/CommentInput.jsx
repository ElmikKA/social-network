import React from 'react'
import { useSetComment } from '../api'
import { CustomFileInput } from './ui/CustomeFile'

const CommentInput = ({ refreshComments, postId }) => {

    const { commentData, handleFileChange, handleChange, handleSubmit } = useSetComment(postId, refreshComments)

    const submit = (e) => {
        handleSubmit(e)

    }

    return (
        <form onSubmit={submit} method='POST' className='comment-input'>
            <div style={{width: '100%', marginTop: '10px'}}>
                <CustomFileInput id="avatar" accept="image/*" onChange={handleFileChange}/>
                <input type="text" id='content' placeholder='comment...' required value={commentData.content} onChange={handleChange} />
            </div>


            <button type='submit'>Comment</button>
        </form>
    )
}

export default CommentInput
