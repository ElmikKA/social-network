import React, { useState } from 'react'
import { useCreatePost } from '../../api'
import { CustomFileInput } from '../../components/ui/CustomeFile'

const CreateGroupPost = ({ setRefreshTrigger, groupId }) => {

    const { postData, handleChange, handleFileChange, handleSubmit } = useCreatePost(groupId)
    const submit = (e) => {
        handleSubmit(e)
        setRefreshTrigger(post => !post)
    }

    return (
        <div className='create-group-post-div'>
            <form onSubmit={submit} method='POST' style={{width: '100%'}}>
                <h3>Create a New Post</h3>
                <div className='form-group'>
                    <div className='form-field'>
                        <input type="text" id="title" placeholder='Title' required value={postData.title} onChange={handleChange} />
                    </div>
                    
                </div>
                
                <div className='form-field' style={{width: '97%'}}>
                    <textarea type="text" id="content" placeholder="What's on your mind?" required value={postData.content} onChange={handleChange} />
                </div>
                
                <div className='form-group' style={{marginTop: '20px'}}>
                    <div className='form-field'>
                        <CustomFileInput id="avatar" accept="image/*" onChange={handleFileChange}/>
                    </div>
                    <div>
                        <button type='submit' className='create-post-submit-button'>Create post</button>
                    </div>
                </div>
            </form>
        </div>
    )
}

export default CreateGroupPost
