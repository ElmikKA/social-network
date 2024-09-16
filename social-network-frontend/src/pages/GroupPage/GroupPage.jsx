import React from 'react'
import { useGetGroupData, useSendGroupJoinRequest } from '../../services/api'
import { useParams } from 'react-router-dom'
import EventBox from '../../components/EventBox'
import CreateEvent from '../../addingStuff/CreateEvent'
import CreateGroupPost from '../CreatePost/CreateGroupPost'
import GroupsBox from '../../components/GroupsBox'
import GroupPostBox from '../../components/GroupPostBox'

const GroupPage = () => {
    const { id } = useParams()
    const { groupData, loading } = useGetGroupData(Number(id))

    // Only call the hook when `groupData` is available and contains group ID
    const sendJoinRequest = useSendGroupJoinRequest()

    if (loading) return <p>Loading...</p>

    if (groupData?.response === "failure") {
        return <p style={{ width: '75%' }}>Error 404, page not found</p>
    }

    const handleJoinGroup = () => {
        sendJoinRequest(groupData.groupData.id)
    }

    return (
        <div className='groupPage'>
            <h2>Group page</h2>
            <p>{groupData.groupData.title}</p>
            <p>Description: {groupData.groupData.description}</p>

            {groupData?.joinStatus === 'completed' ? (
                <div>
                    <EventBox events={groupData.groupEvents} />
                    <CreateEvent groupId={groupData.groupData.id} />
                    <CreateGroupPost groupId={groupData.groupData.id} />
                    <GroupPostBox posts={groupData.groupPosts} />
                </div>
            ) : groupData?.joinStatus === 'pending' ? (
                <p>Join request pending...</p>
            ) : groupData?.joinStatus === "" ? (  // Render button only when joinStatus is empty string
                <button onClick={handleJoinGroup}>Join Group</button>
            ) : null} {/* Render nothing if joinStatus is undefined or any other value */}
        </div>
    )
}

export default GroupPage
