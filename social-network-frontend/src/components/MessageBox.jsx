import React, { useEffect, useState, useRef } from 'react';
import { useGetMessages } from '../api';
import { useParams } from 'react-router-dom';
import { GetSocket } from '../WebSocket';

const MessageBox = () => {

    const { id } = useParams();
    const receiverId = Number(id);

    const { messages: initialMessages, loading } = useGetMessages(receiverId);
    const [messages, setMessages] = useState([])

    const messagesEndRef = useRef(null); // Create a ref for the end of the message box
    const scrollToBottom = () => {
        if (messagesEndRef.current) {
            messagesEndRef.current.scrollIntoView({ behavior: 'smooth' });
        }
    };

    useEffect(() => {
        if (initialMessages && initialMessages.messages) {

            setMessages(initialMessages.messages)
        } else {
            setMessages([])
        }
    }, [initialMessages])

    useEffect(() => {
        scrollToBottom(); 
    }, [messages]); 

    if (loading) {
        return <p>Loading messages...</p>;
    }

    const socket = GetSocket()

    socket.onmessage = (data) => {
        const newMessage = JSON.parse(data.data)
        console.log("socket message", newMessage)

        if (newMessage.type === "message") {
            if (newMessage.partnerId === receiverId) {
                setMessages(prevMessages => prevMessages ? [...prevMessages, newMessage] : [newMessage]); // Append new message
            }
        }
    }

    if (!messages || messages.length === 0) return <div className='messages-box'>No messages</div>
    console.log('messages', messages)
    return (
        <div className='messages-box'>
            {messages.map((message, index) => (
                <div key={index}>
                    <strong> {message.name}:</strong> {message.message} {index}
                </div>
            ))}
            <div ref={messagesEndRef} />
        </div>
    );
};

export default MessageBox;
