import React, { useState } from 'react';
import CommentToggle from './ui/CommentToggle';
import { useGetUser } from '../api';
import { getInitials, timeAgo } from '../utils/helpers.js';
import { useNavigate } from 'react-router-dom';
import menuDots from '../assets/Menu-dots.png';
import like from '../assets/like-button.png';
import comment from '../assets/comment-button.png';

const PostBox = ({ allPosts }) => {
    return (
        <div className='postMain'>
            {allPosts && allPosts.length > 0 ? (
                allPosts.map((post) => (
                    post.CanSee && (
                        <PostItem key={post.id} post={post} />
                    )
                ))
            ) : (
                <div className='no-posts-or-private-profile'>
                    <p>No posts</p>
                </div>
            )}
        </div>
    );
};

export const PostItem = ({ post, refreshTrigger }) => {
    const { userData, loading, error } = useGetUser(post.userId, refreshTrigger);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [showComments, setShowComments] = useState(false); // State to toggle comments

    const handleCommentToggle = () => {
        setShowComments(!showComments);
    };

    const openModal = () => {
        setIsModalOpen(true);
    };

    const closeModal = () => {
        setIsModalOpen(false);
    };

    if (loading) return <p>Loading user...</p>;
    if (error) return <p>Error loading user</p>;

    return (
        <div className='post-box'>
            <div className='post-header'>
                <div className='user-avatar-and-name'>
                    {userData.getUser.avatar ? (
                        <img src={`http://localhost:8080/api/avatars/${userData.getUser.avatar}`}alt="User Avatar" className="user-avatar" />
                    ) : (
                        <div className="avatar-placeholder">
                            {getInitials(`${userData.getUser.firstName} ${userData.getUser.lastName}`)}
                        </div>
                    )}
                    <div className='name-and-time-div'>
                        <p><strong>{userData ? `${userData.getUser.firstName} ${userData.getUser.lastName}` : 'Unknown'}</strong></p>
                        <p style={{ fontSize: '15px', color: '#7E7E7E' }}>{timeAgo(post.createdAt)}</p>
                    </div>
                </div>
                <div className='menu-dots'>
                    <img src={menuDots} alt="Menu" />
                </div>
            </div>
            
            <div className='post-content'>
                <div className='post-title-and-content'>
                    <h2>{post.title}</h2>
                    <p>{post.content}</p>
                </div>
                {post.avatar && (
                    <>
                        <div className="post-image-container" onClick={openModal}>
                            <img
                                src={`http://localhost:8080/api/avatars/${post.avatar}`}
                                alt="Post Image"
                                style={{ width: '100%', height: '100%', borderRadius: '10px' }}
                            />
                        </div>

                        {isModalOpen && (
                            <div className="modal" onClick={closeModal}>
                                <div className="modal-content">
                                    <img src={`http://localhost:8080/api/avatars/${post.avatar}`} alt="Full content" className="full-image" />
                                </div>
                            </div>
                        )}
                    </>
                )}
            </div>

            <div className='like-and-comment-section'>
                <div className='like-and-comment'>
                    <img src={like} alt="Like" style={{ marginTop: '4px' }} />
                    <p>Like</p>
                </div>
                <div className='like-and-comment' onClick={handleCommentToggle}>
                    <img src={comment} alt="Comment" style={{ height: '25px', width: '25px' }} />
                    <p>Comments</p>
                </div>
            </div>
            
            {showComments && <CommentToggle postId={post.id} />}
        </div>
    );
};

export default PostBox;
