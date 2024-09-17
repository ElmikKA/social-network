import React from 'react'

const PostBox = ({ allPosts }) => {
    return (
        <div className='postMain'>
            {allPosts && allPosts.length > 0 ? (
                allPosts.map((post) => (
                    post.CanSee && (

                        <div key={post.id} className='postBox'>
                            <h2>{post.title}</h2>
                            <p>{post.content}</p>
                            <p><strong>Created by:</strong> {post.creator}</p>
                            <p><strong>Created at:</strong> {new Date(post.createdAt).toLocaleString()}</p>
                            {post.avatar ? (
                                <img
                                    src={`http://localhost:8080/api/avatars/${post.avatar}`}
                                    alt="Post Avatar"
                                    style={{ width: '150px', height: '150px', borderRadius: '50%' }}
                                />
                            ) : (
                                <p>No avatar available</p>
                            )}
                        </div>
                    ))
                )
            ) : (
                <p>no posts</p>
            )}
        </div>
    )
}

export default PostBox
