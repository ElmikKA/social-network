import React from 'react'
import { useGetAllGroups, useGetGroupData } from '../../services/api'
import GroupsBox from '../../components/GroupsBox'

const GroupsPage = () => {

    const { groupData, loading } = useGetAllGroups()
    if (loading) return <p>Loading...</p>
    console.log(groupData)


    return (
        <div className='groupsMain'>
            <GroupsBox groups={groupData.groupData} />
        </div>
    )
}

export default GroupsPage
