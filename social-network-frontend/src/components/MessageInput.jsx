import React, { useState, useEffect } from 'react';
import { GetSocket } from '../WebSocket';
import { useParams } from 'react-router-dom';

const MessageInput = () => {
    const [message, setMessage] = useState('');
    const { id } = useParams();
    const receiverId = Number(id);

    const socket = GetSocket()

    const handleSendMessage = () => {
        if (socket && socket.readyState === WebSocket.OPEN && message.trim()) {
            const messageObj = {
                message: message.trim(),
                receiverId: receiverId,
                type: 'message',
            };
            socket.send(JSON.stringify(messageObj));
            setMessage('');
        } else {
            console.log('WebSocket is not connected or message is empty');
        }
    };

    const handleChange = (e) => {
        setMessage(e.target.value);
    };

    const handleKeyPress = (e) => {
        if (e.key === 'Enter') {
            handleSendMessage();
        }
    };

    return (
        <div className='messages-input-div'>
            <input
                type="text"
                placeholder='Type a message...'
                value={message}
                onChange={handleChange}
                onKeyPress={handleKeyPress}
            />
            <button onClick={handleSendMessage}>Send</button>
        </div>
    );
};

export default MessageInput;
