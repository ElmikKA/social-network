import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import MessageBox from '../../components/MessageBox';
import MessageInput from '../../components/MessageInput';

const MessagePage = () => {


    return (
        <div>
            <h1>Message Page</h1>
            <MessageBox />
            <MessageInput />
        </div>
    );
};

export default MessagePage;
