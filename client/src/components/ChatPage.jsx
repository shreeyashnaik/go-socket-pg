import React, { useEffect, useState } from 'react';
import { Offcanvas, Nav, Button } from 'react-bootstrap';
import { useParams, useLocation } from "react-router-dom";
import { connectWebSocket } from "../utils/websocket";

const ChatPage = () => {

  const { chatId } = useParams();
  const {state} = useLocation();
  const { chat_id, chat_name, username } = state;

  useEffect( () => {
    let ws = connectWebSocket();

    return () => {
      if (ws) {
        ws.close();
        console.log("WebSocket connection cleaned up.");
      }
    };
  }, [chat_id]);

  return (
    <div>
      <h1>Welcome to Chat Room: {chat_name}</h1>
      <p>Room ID: {chat_id}</p>
      <p>{username} has entered the chat!</p>
      {/* Add your chat UI here */}

      
    </div>
  );
};

export default ChatPage;
