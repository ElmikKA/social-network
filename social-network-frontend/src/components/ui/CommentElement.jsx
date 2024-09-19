import React, { useState } from 'react'
import CommentInput from '../CommentInput'
import CommentBox from '../CommentBox'
import { useGetComments } from '../../api'

const CommentElement = ({ postId }) => {
    const [refreshCommentTrigger, setRefreshCommentTrigger] = useState(false)
    const { comments, loading } = useGetComments(postId, refreshCommentTrigger)

    const refreshComments = () => {
        console.log("refreshing comments")
        setRefreshCommentTrigger(prev => !prev)
    }
    if (loading) return <p>Loading comments...</p>

    return (
        <div>
            <CommentInput refreshComments={refreshComments} postId={postId} />
            <CommentBox comments={comments} />
        </div>
    )
}

export default CommentElement
