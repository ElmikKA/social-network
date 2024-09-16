
import React from 'react';
import { useGetUser } from '../../services/api';
import FollowButton from '../../components/ui/FollowButton';
import { useParams } from 'react-router-dom';

const Profile = () => {
    const { id } = useParams()
    const { userData, loading, error } = useGetUser(id);

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error}</p>;
    console.log(userData)

    return (
        <div className='profile'>
            <h2>Profile Page</h2>
            <p>Name: {userData.getUser.name}</p>
            <p>Email: {userData.getUser.email}</p>
            <p>First Name: {userData.getUser.firstName}</p>
            <p>Last Name: {userData.getUser.lastName}</p>

            <div>
                <h3>Profile picture</h3>
                <img
                    src={`http://localhost:8080/api/avatars/${userData.getUser.avatar ? userData.getUser.avatar : '/db/assets /default.webp'}`}
                    alt="Profile Avatar"
                    style={{ width: '100px', height: '100px', borderRadius: '50%' }}
                />
            </div>

            {!userData.ownPage && <div>

                {
                    userData.following !== 'not following' ? <p className='followButton'>{userData.following === "completed" ? 'following' : userData.following}</p> :
                        <FollowButton userId={id} />
                }
            </div>
            }
        </div>
    );
};

export default Profile;
