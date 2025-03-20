import React from 'react';
import { Link } from "react-router-dom";
import NewChat from './NewChat';
import JoinChat from './JoinChat';

const Home = () => {
  return (
    <>
    <h2>My Chat App</h2> <br />
    <NewChat /> <br /><br />
    <JoinChat />
    </>
  )
}

export default Home