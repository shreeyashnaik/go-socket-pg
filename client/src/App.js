import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Home from './components/Home';
import ChatPage from './components/ChatPage';

export default function App() {

    const router = createBrowserRouter([
        {
            path: "/",
            element: <Home />
        },
        {
            path: "/chats/:chatId",
            element: <ChatPage />
        },       
    ]);

    return (
        <>
         <RouterProvider router={router} />
        </>
    );
}

// function Home() {
//     const navigate = useNavigate();

//     function handleNewChat() {
//         console.log("New Chat");
//         navigate("/chats/new");
//     }

//     function handleJoinChat() {
//         console.log("Join Chat");
//         navigate("/chats/join");
//     }

//     return (
//         <>
//             <h1>Home</h1>
//             <button className="chat" onClick={handleNewChat}>New Chat</button>
//             <br /><br />
//             <button className="chat" onClick={handleJoinChat}>Join Chat</button>
//         </>
//     );
// }

// function NewChat() {
//     function NewChat() {
//         const [roomName, setRoomName] = useState('');
//         const [yourName, setYourName] = useState('');

//         // async function handleSubmit(event) {
//         //     event.preventDefault();
//         //     const response = await fetch('/api/addUserToChat', {
//         //         method: 'POST',
//         //         headers: {
//         //             'Content-Type': 'application/json',
//         //         },
//         //         body: JSON.stringify({ roomName, yourName }),
//         //     });

//         //     if (response.ok) {
//         //         console.log('User added to chat');
//         //         // Handle successful response
//         //     } else {
//         //         console.error('Failed to add user to chat');
//         //         // Handle error response
//         //     }
//         // }

//         return (
//             <>
//                 <h1>New Chat</h1>

//                 <form onSubmit={createNewChat}>
//                     <label>
//                         Room Name:
//                         <input
//                             type="text"
//                             name="roomName"
//                             value={roomName}
//                             onChange={(e) => setRoomName(e.target.value)}
//                         />
//                     </label>
//                     <br /><br />
//                     <label>
//                         Your Name:
//                         <input
//                             type="text"
//                             name="yourName"
//                             value={yourName}
//                             onChange={(e) => setYourName(e.target.value)}
//                         />
//                     </label>
//                     <br /><br />
//                     <button type="submit">Create Room</button>
//                 </form>
//             </>
//         );
//     }
// }

// function JoinChat() {
//     return (
//         <>
//             <h1>Join Chat</h1>

//             <form>
//                 <label>
//                     Enter Chat link:
//                     <input type="text" name="chatLink" />
//                 </label>
//                 <br /><br />
//                 <label>
//                     Your Name:
//                     <input type="text" name="yourName" />
//                 </label>
//                 <br /><br />
//                 <button type="submit">Join Room</button>
//             </form>
//         </>
//     );
// }