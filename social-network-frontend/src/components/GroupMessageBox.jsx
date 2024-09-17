import React, { useEffect, useState } from 'react';
import { useGetGroupMessages, useGetMessages } from '../api';
import { useParams } from 'react-router-dom';
import { GetSocket } from '../WebSocket';

const GroupMessageBox = () => {

    const { id } = useParams();
    const GroupId = Number(id);
    console.log("groupId:", GroupId)

    const { messages: initialMessages, loading } = useGetGroupMessages(GroupId)
    const [messages, setMessages] = useState([])

    useEffect(() => {
        if (initialMessages && initialMessages.messages)
            setMessages(initialMessages.messages)
    }, [initialMessages])


    if (!messages || messages.length === 0) return <div>No messages</div>

    if (loading) {
        return <p>Loading messages...</p>;
    }

    const socket = GetSocket()

    socket.onmessage = (data) => {
        const newMessage = JSON.parse(data.data)

        if (newMessage.type === "groupMessage") {
            if (newMessage.groupId === GroupId) {
                console.log("here")
                setMessages(prevMessages => prevMessages ? [...prevMessages, newMessage] : [newMessage]);
            }
        }
    }

    return (
        <div>
            {messages.map((message, index) => (
                <div key={index}>
                    <strong> {message.name}:</strong> {message.message}
                </div>
            ))}
        </div>
    );
};

export default GroupMessageBox;
