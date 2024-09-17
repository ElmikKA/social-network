
import { w3cwebsocket as Socket } from "websocket";

let globalSocket = null;

export const GetSocket = () => {
    return globalSocket;
};

export const InitSocket = () => {
    if (!globalSocket) {
        globalSocket = new Socket("ws://localhost:8080/api/websocket");

        globalSocket.onopen = () => {
            console.log("WebSocket connection established.");
        };

        globalSocket.onmessage = (event) => {
            console.log("WebSocket message received:", event.data);
        };

        globalSocket.onclose = () => {
            console.log("WebSocket connection closed.");
            globalSocket = null;
        };

        globalSocket.onerror = (error) => {
            console.error("WebSocket error", error);
        };
    }

    return globalSocket;
};

export const closeSocket = () => {
    if (globalSocket) {
        globalSocket.close()
        console.log("closed socket")
    } else {
        console.log("no active websockets")
    }
}