import React from 'react';
import { useWebSocketForMessagePage } from '../WebSocket';

const MessageBox = () => {
    const { messages } = useWebSocketForMessagePage();

    if (!messages || messages.length === 0) return <div>No messages</div>;

    return (
        <div>
            {messages.map((message, index) => (
                <div key={index}>
                    {message.name}: {message.message}
                </div>
            ))}
        </div>
    );
};

export default MessageBox;
