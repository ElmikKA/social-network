import React, { useEffect, useState } from 'react';
import { useGetMessages } from '../api';
import { useParams } from 'react-router-dom';
import { GetSocket } from '../WebSocket';

const MessageBox = () => {

    const { id } = useParams();
    const receiverId = Number(id);
    console.log("receiver:", receiverId)

    const { messages: initialMessages, loading } = useGetMessages(receiverId);
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
        console.log("socket message", newMessage)

        if (newMessage.type === "message") {
            console.log(receiverId, newMessage.partnerId)
            if (newMessage.partnerId === receiverId) {
                setMessages(prevMessages => prevMessages ? [...prevMessages, newMessage] : [newMessage]); // Append new message
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

export default MessageBox;
