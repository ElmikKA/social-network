
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
            <p>Name: {userData.name}</p>
            <p>Email: {userData.email}</p>
            <p>First Name: {userData.firstName}</p>
            <p>Last Name: {userData.lastName}</p>
            {userData.avatar ? (
                <div>
                    <h3>Profile picture</h3>
                    <img
                        src={`http://localhost:8080/api/avatars/${userData.avatar}`}
                        alt="Profile Avatar"
                        style={{ width: '150px', height: '150px', borderRadius: '50%' }}
                    />
                </div>
            ) : (
                <p>No profile picture available</p>
            )}
            {userData.following !== 'not following' ? <p className='followButton'>{userData.following}</p> :
                <FollowButton userId={id} />}
        </div>
    );
};

export default Profile;
