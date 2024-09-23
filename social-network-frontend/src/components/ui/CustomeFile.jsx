import React, { useState } from 'react';

export const CustomFileInput = ({ id, accept, onChange }) => {
    const [fileName, setFileName] = useState('Select picture or GIF');

    const handleFileChange = (e) => {
        if (e.target.files.length > 0) {
            setFileName(e.target.files[0].name);
        } else {
            setFileName('No file selected.');
        }

        if (onChange) {
            onChange(e);
        }
    };

    return (
        <div className="file-input-container">
            <label htmlFor={id} className="custom-file-upload"></label>
            <input 
                id={id} 
                type="file" 
                className="file-input" 
                onChange={handleFileChange}
                accept={accept}
            />
            <span className="file-name">{fileName}</span>
        </div>
    );
};