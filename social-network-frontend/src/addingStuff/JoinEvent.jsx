import React, { useEffect } from 'react'

const JoinEvent = () => {

    const eventId = 2
    const response = 'completed' // "completed" or "rejected"

    useEffect(() => {
        const sendEventResponse = async () => {
            const requestOptions = {
                method: "POST",
                credentials: 'include',
                header: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ eventId: eventId, pending: response })
            }
            try {
                const response = await fetch('http://localhost:8080/api/respondEvent', requestOptions)
                const data = await response.json()
                console.log(data)
            } catch (err) {
                console.log(err)
            }
        }
        sendEventResponse()
    })


    return (
        <div>
            send event response

        </div>
    )
}

export default JoinEvent