import React, { useEffect, useState } from 'react'
import CommentElement from './CommentElement'

const CommentToggle = ({ postId }) => {
    const [opened, setOpened] = useState(false)
    const handleClick = () => {
        setOpened(prev => !prev)
    }
    return (
        <div className='commentToggle' >
            <button onClick={handleClick}>click me</button>
            {
                opened ? (
                    <>
                        <p>opened comments</p>
                        <CommentElement postId={postId} />

                    </>
                ) : <p>comment</p>
            }
        </div >
    )
}

export default CommentToggle
