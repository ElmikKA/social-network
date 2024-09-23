import React from 'react'
import { useGetAllGroups, useGetGroupData } from '../../api'
import GroupsBox from '../../components/GroupsBox'

const GroupsPage = () => {

    const { groupData, loading } = useGetAllGroups()
    if (loading) return <p>Loading...</p>
    console.log(groupData)


    return (
        <div className='groups-main'>
            <GroupsBox groups={groupData.groupData} />
        </div>
    )
}

export default GroupsPage
