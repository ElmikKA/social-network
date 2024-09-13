import React, { useEffect, useState } from 'react'
import { useGetUser } from '../../services/api';

const Profile = () => {

    // send the id of the user
    const userId = 1
    const { userData } = useGetUser(userId)

    // still needs to add:
    // all user made posts
    // all followers and following


    return (
        <div>
            <h2>Profile Page</h2>
            <p>Name: {userData.name}</p>
            <p>Email: {userData.email}</p>
            <p>First Name: {userData.firstName}</p>
            <p>Last Name: {userData.lastName}</p>
            {userData.avatar ? (
                <div> <h3>Profile picture</h3>
                    <img src={`http://localhost:8080/api/avatars/${userData.avatar}`} alt="Profile Avatar" style={{ width: '150px', height: '150px', borderRadius: '50%' }}
                    />
                </div>
            ) : (
                <p>No profile picture available</p>
            )}
        </div>
    );
}

export default Profile
