const connectWebSocket = (chatId, username, onMessageCallback) => {
    return new Promise((resolve, reject) => {
        const ws = new WebSocket(`ws://localhost:8071/ws?chat_id=${chatId}`);
    
        ws.onopen = () => {
            console.log("WebSocket connection established.");
            const bootupMessage = JSON.stringify({
                type: "BOOTUP",
                chat_id: chatId,
                username: username,
                timestamp: new Date().toISOString(),
            });
            ws.send(bootupMessage);
            console.log("BOOTUP message sent.");
            resolve(ws);
        };
    
        ws.onerror = (error) => {
            console.error("WebSocket error:", error);
            reject(error);
        };
    
        ws.onclose = () => {
            const disconnectMessage = JSON.stringify({
                type: "DISCONNECT",
                chat_id: chatId,
                username: username,
                timestamp: new Date().toISOString(),
            });
            ws.send(disconnectMessage);
            console.log("WebSocket connection closed.");
        };
    
        ws.onmessage = (event) => {
            if (onMessageCallback) {
                onMessageCallback(event.data);
            }
            console.log("Message from server:", event.data);
        };
    });
};

export { connectWebSocket };