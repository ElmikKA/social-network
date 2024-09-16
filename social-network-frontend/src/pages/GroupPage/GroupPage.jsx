import React from 'react'
import { useGetGroupData } from '../../services/api'
import { useParams } from 'react-router-dom'

const GroupPage = () => {
    const { id } = useParams("id")

    const { groupData, loading } = useGetGroupData(Number(id))
    if (loading) return <p>Loading...</p>
    if (groupData) console.log("grouppage", groupData)
    if (groupData.response === "failure") {
        return <p style={{ width: '75%' }}>error 404, page not found</p>
        //add error page
    }


    return (
        <div className='groupPage'>
            <h2>Group page</h2>
            <p>{groupData.groupData.title}</p>
            <p>description:{groupData.groupData.description}</p>

        </div>
    )
}

export default GroupPage



