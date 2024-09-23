import React from 'react';
import MessageBox from '../../components/MessageBox';
import MessageInput from '../../components/MessageInput';
import { useGetContacts } from '../../api'
import { useParams } from 'react-router-dom';

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
