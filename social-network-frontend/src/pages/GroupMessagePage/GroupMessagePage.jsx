
import React from 'react';
import GroupMessageBox from '../../components/GroupMessageBox';
import GroupMessageInput from '../../components/GroupMessageInput';

const GroupMessagePage = () => {


    return (
        <div className='group-messages-page'>
            <h1>Group Message Page</h1>
            <GroupMessageBox />
            <GroupMessageInput />
        </div>
    );
};

export default GroupMessagePage;
