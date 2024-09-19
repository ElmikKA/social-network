import React, { useEffect, useState } from 'react'
import CommentElement from './CommentElement'

const CommentToggle = ({ postId }) => {

    return (
        <div className='commentToggle' >
            <p>opened comments</p>
            <CommentElement postId={postId} />
        </div >
    )
}

export default CommentToggle