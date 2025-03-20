package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var chat_ids = map[string]*os.File{}

type Message struct {
	SenderName string `json:"sender_name"`
	DM         *struct {
		ReceiverName string `json:"receiver_name"`
	} `json:"dm,omitempty"`
	Group *struct {
		GroupID uint
	} `json:"group,omitempty"`
	Content    string    `json:"content"`
	Disconnect bool      `json:"disconnect"`
	Timestamp  time.Time `json:"timestamp"`
}

type NewChatRequest struct {
	RoomName    string `json:"room_name"`
	CreatorName uint   `json:"creator_name"`
}

type NewChatResponse struct {
	ChatID string `json:"chat_id"`
}

func main() {

	server, err := net.Listen("tcp", ":7070")
	if err != nil {
		panic(err)
	}

	log.Println("Server started on: ", server.Addr())

	// Using an infinite loop to accept connections from multiple clients
	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		log.Println("Accepted connection from client: ", conn.RemoteAddr())

		go handleSingleClientConnection(conn)
	}

}

func handleSingleClientConnection(conn net.Conn) {

	// Using an infinite loop to check messages sent by Client_X
	for {
		// Get message
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			log.Println("error reading message from client: ", err)
			continue
		}

		// Parse message into struct msg
		msg := Message{}
		if err := json.Unmarshal(buff[:n], &msg); err != nil {
			log.Println("error unmarshalling message from client: ", err)
			continue
		}
		log.Println("Received message from client: ", msg)

		// When the Client_X sends "!DISCONNECT", the server will close the connection
		if msg.Disconnect {
			log.Println("Client disconnected")
			if err := conn.Close(); err != nil {
				log.Println("error closing connection: ", err)
			} else {
				log.Println("conn disconnected by server")
			}
			return
		}

		if msg.DM != nil {
			if err := processDM(&msg); err != nil {
				log.Println("failed to process dm:", err)
				continue
			}
		} else if msg.Group != nil {

		} else {
			log.Println("chat type cannot be determined")
			continue
		}

	}
}

func processDM(msg *Message) error {
	key := ""
	if msg.DM.ReceiverName < msg.SenderName {
		key = fmt.Sprintf("%v_%v", msg.DM.ReceiverName, msg.SenderName)
	} else {
		key = fmt.Sprintf("%v_%v", msg.DM.ReceiverName, msg.DM.ReceiverName)
	}

	fileObj, ok := chat_ids[key]
	if !ok {
		// create a file with name 'key'
		var err error
		fileObj, err = os.OpenFile("chats/"+key, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		chat_ids[key] = fileObj
	}

	fileObj.WriteString(fmt.Sprintf("[%v]%v >>>> %v\n", time.Now().Format("15:04"), msg.SenderName, msg.Content))

	return nil
}
