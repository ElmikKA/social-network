import React from 'react';
import MessageBox from '../../components/MessageBox';
import MessageInput from '../../components/MessageInput';

const MessagePage = () => {
    return (
        <div className='message-page'>
            <h1>Message Page</h1>
            <MessageBox />
            <MessageInput />
        </div>
    );
};

export default MessagePage;
