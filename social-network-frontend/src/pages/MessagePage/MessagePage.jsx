import React, { useEffect } from 'react';
import { useParams } from 'react-router-dom';
import MessageBox from '../../components/MessageBox';
import MessageInput from '../../components/MessageInput';
import { WebSocketProviderForMessagePage } from '../../WebSocket';
import { useGetMessages } from '../../api';

const MessagePage = () => {
    const { id } = useParams();
    const receiverId = Number(id);
    const { messages: initialMessages, loading } = useGetMessages(receiverId);
    console.log(initialMessages.messages)

    if (loading) return <p>Messages are loading...</p>;



    return (
        <WebSocketProviderForMessagePage initialMessages={initialMessages.messages} partnerId={receiverId}>
            <div>
                <h1>Message Page</h1>
                <MessageBox />
                <MessageInput receiverId={receiverId} myId={initialMessages.myId} />
            </div>
        </WebSocketProviderForMessagePage>
    );
};

export default MessagePage;
