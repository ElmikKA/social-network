import React, { useEffect, useState } from 'react'
import { useGetGroupData, useSendGroupJoinRequest } from '../../api'
import { useParams } from 'react-router-dom'
import EventBox from '../../components/EventBox'
import CreateEvent from '../../addingStuff/CreateEvent'
import CreateGroupPost from '../CreatePost/CreateGroupPost'
import GroupsBox from '../../components/GroupsBox'
import GroupPostBox from '../../components/GroupPostBox'
import InviteGroup from '../../components/InviteGroup'
import ToggleInviteGroup from '../../components/ui/ToggleInviteGroup'

const GroupPage = () => {
    const { id } = useParams()
    const [refreshTrigger, setRefreshTrigger] = useState(false)
    const { groupData, loading } = useGetGroupData(Number(id), refreshTrigger)

    // Only call the hook when `groupData` is available and contains group ID
    const sendJoinRequest = useSendGroupJoinRequest()

    if (loading) return <p>Loading...</p>

    if (groupData?.response === "failure") {
        return <p style={{ width: '75%' }}>Error 404, page not found</p>
    }

    const handleJoinGroup = () => {
        sendJoinRequest(groupData.groupData.id)
        setRefreshTrigger(prev => !prev)
    }

    return (
        <div className='group-page'>

            <div className='group-header-and-information'>
                <div className='group-picture-and-name'>
                    <div className='group-picture' style={{width: '100px', height: '100px', borderRadius: '50px', backgroundColor: 'white', display: 'flex', justifyContent: 'center', alignItems: 'center', color: 'black'}}>
                        <p>NEW</p>
                    </div>
                    <h2>{groupData.groupData.title}</h2>

                    <div className='group-posts-members-events-section'>
                        <div className='members-posts-events'>
                            <p>Members</p>
                            <h3>{groupData.groupMembers !== null ? groupData.groupMembers.length : '0'}</h3>
                        </div>
                        <div className='members-posts-events'>
                            <p>Posts</p>
                            <h3>{groupData.groupPosts !== null ? groupData.groupPosts.length : '0'}</h3>
                        </div>
                        <div className='members-posts-events'>
                            <p>Events</p>
                            <h3>{groupData.groupEvents !== null ? groupData.groupEvents.length : '0'}</h3>
                        </div>
                    </div>

                    <div className='join-or-invite-button'>

                    {groupData?.joinStatus === 'completed' && groupData.owner ? (
                        <div>
                            <ToggleInviteGroup groupId={groupData.groupData.id}  />
                        </div>
                    ) : groupData?.joinStatus === 'pending' ? (
                        <p>Join request pending...</p>
                    ) : groupData?.joinStatus === "" ? ( 
                        <button onClick={handleJoinGroup}>Join Group</button>
                    ) : null} 
                    </div>

                </div>

                <div className='about-the-group-div'>
                    <h3>About the Group</h3>
                    <p>{groupData.groupData.description}</p>
                </div>

                <div className='group-events'>
                    <h3>Upcoming Group Events</h3>
                    {groupData.joinStatus === 'completed' ? 
                        <EventBox setRefreshTrigger={setRefreshTrigger} events={groupData.groupEvents} /> : 
                        <p>Only Members can see the events!</p>
                    }
                </div>

            </div>

            <div className='group-posts'>
                {groupData?.joinStatus === 'completed' ? 
                    <div>
                        <CreateGroupPost setRefreshTrigger={setRefreshTrigger} groupId={groupData.groupData.id} />
                        <CreateEvent setRefreshTrigger={setRefreshTrigger} groupId={groupData.groupData.id} />
                        <GroupPostBox posts={groupData.groupPosts} />
                    </div> 
                    : 
                    <p>Need to join the group to see information</p>
                }
            </div>
        </div>
    )
}

export default GroupPage
