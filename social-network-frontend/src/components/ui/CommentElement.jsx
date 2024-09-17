import React from 'react'
import CommentInput from '../CommentInput'
import { useGetComments } from '../../api'
import CommentBox from '../CommentBox'

const CommentElement = ({ postId }) => {
    console.log("fetch comments", postId)

    const { comments } = useGetComments(postId)
    if (comments) console.log(comments)

    return (
        <div>
            <CommentInput postId={postId} />
            <CommentBox comments={comments} />
        </div>
    )
}

export default CommentElement
