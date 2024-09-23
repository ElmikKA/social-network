import React from 'react'
import { useGetUser } from '../api'
import { getInitials, timeAgo } from '../utils/helpers'

const CommentBox = ({ comments = [] }) => { // Default to empty array

    console.log('comments', comments)

    return (
        <div className='comment-main'>
            {comments.length > 0 ? (
                comments.map((comment) => (
                    <CommentItem key={comment.id} comment={comment}></CommentItem>
                ))
            ) : <p className='comment-box'>No comments</p>}
        </div>
    )
}

const CommentItem = ({comment}) => {
    const { userData, loading, error } = useGetUser(comment.userId, false);

    if (loading) return <p>Loading user...</p>;
    if (error) return <p>Error loading user</p>;

    return (
        <div key={comment.id} className='comment-box'>
            <div className='user-avatar-and-name-comment'>
                    {userData.getUser.avatar ? (
                        <img src={`http://localhost:8080/api/avatars/${userData.getUser.avatar}`}alt="User Avatar" className="user-avatar-comment" />
                    ) : (
                        <div className="avatar-placeholder-comment">
                            {getInitials(`${userData.getUser.firstName} ${userData.getUser.lastName}`)}
                        </div>
                    )}
                    <div className='name-and-time-div-comment'>
                        <p><strong>{userData ? `${userData.getUser.firstName} ${userData.getUser.lastName}` : 'Unknown'}</strong></p>
                        <p style={{ fontSize: '12px', color: '#7E7E7E' }}>{timeAgo(comment.createdAt)}</p>
                    </div>
                </div>
            <p style={{marginLeft: '10px', marginTop: '5px'}}>{comment.content}</p>

            {comment.avatar && (
                <img
                    src={`http://localhost:8080/api/avatars/${comment.avatar}`}
                    alt="Post Avatar"
                    style={{ width: '150px', height: '150px', borderRadius: '50%' }}
                />
            )}
        </div>
    )
}

export default CommentBox
