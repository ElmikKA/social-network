
import React, { useEffect, useState } from 'react';
import { useGetUser } from '../../api';
import FollowButton from '../../components/ui/FollowButton';
import { useParams } from 'react-router-dom';
import PostBox from '../../components/PostBox';
import UnFollowButton from '../../components/ui/UnFollowButton';
import ChangePrivacy from '../../components/ChangePrivacy';

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
                    <h2>Profile Page</h2>
                    <p>Name: {userData.getUser.name}</p>
                    <p>Email: {userData.getUser.email}</p>



                    <p>First Name: {userData.getUser.firstName}</p>
                    <p>Last Name: {userData.getUser.lastName}</p>
                </div>

                <div className='profile-posts'>
                    {userData.CanSee ?
                        <>
                            <PostBox allPosts={userData.posts} />
                        </>
                        : <p>private profile</p>
                    }   
                </div>

            </div>
        </div>
    );
};

export default Profile;
