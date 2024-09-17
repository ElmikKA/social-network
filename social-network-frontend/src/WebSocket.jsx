// import React, { createContext, useContext, useRef, useState, useEffect } from 'react';

// const WebSocketContext = createContext(null);

// export const WebSocketProviderForMessagePage = ({ initialMessages = [], partnerId, children }) => {
//     const socket = useRef(null);
//     const [isConnected, setIsConnected] = useState(false);
//     const [messages, setMessages] = useState(initialMessages)
//     const [currentPartnerId, setCurrentPartnerId] = useState(partnerId)

//     useEffect(() => {
//         if (!socket.current) {
//             socket.current = new WebSocket('ws://localhost:8080/api/websocket');
//             socket.current.onopen = () => {
//                 console.log('WebSocket connection established');
//                 setIsConnected(true);
//             };
//             socket.current.onclose = () => {
//                 console.log('WebSocket connection closed');
//                 setIsConnected(false);
//             };
//             socket.current.onerror = (error) => {
//                 console.error('WebSocket error', error);
//             };

//             socket.current.onmessage = (event) => {
//                 const message = JSON.parse(event.data);
//                 console.log(message)
//                 if (message.type === "message") {
//                     console.log(currentPartnerId, message.partnerId)
//                     if (message.partnerId === currentPartnerId) {
//                         setMessages(prevMessages => prevMessages ? [...prevMessages, message] : [message]); // Append new message
//                     }
//                 } else {
//                     console.log("not current speaker")
//                 }
//             };
//         }


//         return () => {
//             if (socket.current) {
//                 socket.current.close();
//                 socket.current = null
//             }
//         };
//     }, []);

//     useEffect(() => {
//         setMessages(initialMessages);
//     }, [initialMessages]);
//     return (
//         <WebSocketContext.Provider value={{ socket: socket.current, messages, currentPartnerId }}>
//             {children}
//         </WebSocketContext.Provider>
//     );
// };

// export const useWebSocketForMessagePage = () => {
//     return useContext(WebSocketContext);
// };


const { WebSocketServer } = require("ws")