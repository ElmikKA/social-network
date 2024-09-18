
import React from 'react'

const CommentBox = ({ comments = [] }) => { // Default to empty array

    return (
        <div className='commentMain'>
            {comments.length > 0 ? (
                comments.map((comment) => (
                    <div key={comment.id} className='commentBox'>
                        <p>{comment.content}</p>
                        <p>Created by: {comment.creator}</p>
                        <p>Created at: {new Date(comment.createdAt).toLocaleString()}</p>

                        {comment.avatar && (
                            <img
                                src={`http://localhost:8080/api/avatars/${comment.avatar}`}
                                alt="Post Avatar"
                                style={{ width: '150px', height: '150px', borderRadius: '50%' }}
                            />
                        )}
                    </div>
                ))
            ) : <p>No comments</p>}
        </div>
    )
}

export default CommentBox
