
import React, { useEffect, useState } from 'react';
import { useGetUser } from '../../api';
import FollowButton from '../../components/ui/FollowButton';
import { useParams } from 'react-router-dom';
import PostBox from '../../components/PostBox';
import UnFollowButton from '../../components/ui/UnFollowButton';
import ChangePrivacy from '../../components/ChangePrivacy';
import privateStatus from '../../assets/private-status.png'
import publicStatus from '../../assets/public-status.png'
import { formatDate } from '../../utils/helpers';

const Profile = () => {
    const { id } = useParams()
    const [refreshTrigger, setRefreshTrigger] = useState(false)
    const { userData, loading, error } = useGetUser(id, refreshTrigger);
    const [privacy, setPrivacy] = useState("")


    console.log(userData)
    useEffect(() => {
        if (userData) {
            setPrivacy(userData.status)
        }
    }, [userData])
    console.log(privacy)
    console.log(userData)

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error}</p>;
    return (
        <div className='profile'>

            <div className='user-profile-header'>
                <div className='user-picture-name-and-email'>
                    <div className='profile-picture'>
                        <img
                            src={`http://localhost:8080/api/avatars/${userData.getUser.avatar ? userData.getUser.avatar : '/db/assets /default.webp'}`}
                            alt="Profile Avatar"
                            style={{ width: '100px', height: '100px', borderRadius: '50%' }}
                        />
                    </div>
                    <div className='user-name-and-email'>
                        <p className='user-name'>{`${userData.getUser.firstName} ${userData.getUser.lastName}`}</p>
                        <p className='user-email'>{userData.getUser.email}</p>
                    </div>
                </div>

                <div className='follow-or-privacy-button'>

                    {userData.ownPage && privacy && <ChangePrivacy status={privacy} />}
                    
                    {!userData.ownPage && <div>
                    {
                        userData.following !== '' ? <p className='followButton'>{userData.following === "completed" ? <UnFollowButton userId={userData.getUser.id} setRefreshTrigger={setRefreshTrigger} /> : userData.following}</p> :
                            <FollowButton userId={id} setRefreshTrigger={setRefreshTrigger} />
                    }
                    </div>}
                </div>
            </div>

            <div className='main-content-for-profile'>

                <div className='profile-information-div'> 
                    <div className='follower-and-post'>
                        <div className='followers-and-posts-conatiner'>
                            <p>Follower</p>
                            <h3>{userData.followers !== null ? userData.followers.length : 0}</h3>
                        </div>

                        <div className='followers-and-posts-conatiner'>
                            <p>Posts</p>
                            <h3>{userData.posts !== null ? userData.posts.length : 0}</h3>
                        </div>
                    </div>

                    <div className='profile-status-and-about-me'>
                        <h3>Profile Status</h3>

                        <div className='profile-status'>
                            {userData.getUser.privacy === 'private' ? 
                                <>
                                    <img src={privateStatus}></img>
                                    <h4>Private</h4>
                                </>
                                :
                                <>
                                    <img src={publicStatus}></img>
                                    <h4>Public</h4>
                                </>
                            }
                        </div>

                        <div className='about-me'>
                            <h3>About Me</h3>
                            <p>{userData.getUser.aboutMe !== '' ? userData.getUser.aboutMe : 'No information'}</p>
                        </div>
                        


                    </div>

                    <div className='user-information'>
                        <h3>User Information</h3>
                        <div className='user-information-label'>
                            <p>Name</p>
                            <h4>{userData.getUser.name}</h4>
                        </div>

                        <div className='user-information-label'>
                            <p>First Name</p>
                            <h4>{userData.getUser.firstName}</h4>
                        </div>

                        <div className='user-information-label'>
                            <p>Last Name</p>
                            <h4>{userData.getUser.lastName}</h4>
                        </div>

                        <div className='user-information-label'>
                            <p>Email</p>
                            <h4>{userData.getUser.email}</h4>
                        </div>

                        <div className='user-information-label'>
                            <p>Date of Birth</p>
                            <h4>{formatDate(userData.getUser.dateOfBirth)}</h4>
                        </div>

                        <div className='user-information-label'>
                            <p>Nickname</p>
                            <h4>{userData.getUser.nickname}</h4>
                        </div>
                    </div>
                </div>

                <div className='profile-posts'>
                    {userData.CanSee ?
                        <>
                            <PostBox allPosts={userData.posts} />
                        </>
                        : <div className='no-posts-or-private-profile'>
                            <p>private profile</p>
                        </div> 
                    }   
                </div>

            </div>
        </div>
    );
};

export default Profile;
