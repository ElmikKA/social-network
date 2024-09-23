import React, { useEffect, useState } from 'react'
import CommentElement from './CommentElement'

const CommentToggle = ({ postId }) => {

    return (
        <div className='comment-toggle' >
            <CommentElement postId={postId} />
        </div >
    )
}

export default CommentToggle