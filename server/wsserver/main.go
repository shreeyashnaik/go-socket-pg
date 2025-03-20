package main

import (
	"chat-app/common/constants/enums"
	"chat-app/common/models"
	"chat-app/common/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

type ClientMap map[*Client]bool

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// Registered clients.
	clients map[string]ClientMap
	// Inbound messages from the clients.
	broadcast chan *models.Message
}

// Client is a user with its websocket connection to the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	// send chan []byte

	// The username of the client
	username string

	// The chatID of the client
	chatID string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {

	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Parse chat_id from query parameters
	chatID := r.URL.Query().Get("chat_id")
	if chatID == "" {
		http.Error(w, "chat_id is required", http.StatusBadRequest)
		return
	}

	client := &Client{hub: hub, conn: ws, chatID: chatID}

	if hub.clients[chatID] == nil {
		hub.clients[chatID] = make(ClientMap)
	}
	hub.clients[chatID][client] = true

	receiver(client)

	delete(hub.clients[chatID], client)
}

func receiver(client *Client) {
	for {
		// read in a message
		// readMessage returns messageType, message, err
		// messageType: 1-> Text Message, 2 -> Binary Message
		_, p, err := client.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		m := &models.Message{}

		err = json.Unmarshal(p, m)
		if err != nil {
			log.Println("error while unmarshaling chat", err)
			continue
		}

		fmt.Println("host", client.conn.RemoteAddr())
		if m.Type == enums.BOOTUP {
			client.username = m.Username
			client.chatID = m.ChatID
			fmt.Println("client successfully mapped", &client, client, client.username)
		} else if m.Type == enums.DISCONNECT {
			fmt.Println("client disconnected", client.username)
			client.conn.Close()
			return
		} else {
			fmt.Println("received message", m)

			m.Timestamp = time.Now()
		}
		client.hub.broadcast <- m
	}
}

func broadcaster(hub *Hub) {
	for {
		message := <-hub.broadcast
		for client := range hub.clients[message.ChatID] {
			log.Println("Message: ", message, "Client: ", client)
			if message.Username != client.username {
				if err := client.conn.WriteJSON(message); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}

func main() {

	// Load environment variables
	utils.LoadEnv()

	// Create a new hub
	hub := &Hub{
		broadcast: make(chan *models.Message),
		clients:   make(map[string]ClientMap),
	}

	// Start the broadcaster
	go broadcaster(hub)

	// Listen for incoming messages
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	// Listen for websocket server
	if err := http.ListenAndServe(viper.GetString("SOCK_PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
