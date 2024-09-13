import React, { useEffect } from 'react'

const JoinGroup = () => {

    const groupId = 1
    useEffect(() => {
        const sendGroupJoinRequest = async () => {
            const requestOptions = {
                method: 'POST',
                credentials: 'include',
            }
            try {
                const response = await fetch(`http://localhost:8080/api/requestGroupJoin/${groupId}`, requestOptions)
                const data = await response.json()
                console.log(data)
            } catch (err) {
                console.log(err)
            }
        }
        sendGroupJoinRequest()
    }, [])

    return (
        <div>

        </div>
    )
}

export default JoinGroup
