import React, { useState, useEffect } from 'react';
import { useWebSocketForMessagePage } from '../WebSocket';

const MessageInput = ({ receiverId, myId }) => {
    const [message, setMessage] = useState('');
    const { socket } = useWebSocketForMessagePage();

    useEffect(() => {
        if (socket && socket.readyState !== WebSocket.OPEN) {
            console.error('WebSocket is not open. Current state:', socket.readyState);
        }
    }, [socket]);

    const handleSendMessage = () => {
        if (socket && socket.readyState === WebSocket.OPEN && message.trim()) {
            const messageObj = {
                message: message.trim(),
                receiverId: receiverId,
                senderId: myId,
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
        <div>
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
