package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type Message struct {
	SenderID   uint      `json:"sender_id"`
	SenderName string    `json:"sender_name"`
	GroupID    uint      `json:"group_id"`
	Content    string    `json:"content"`
	Disconnect bool      `json:"disconnect"`
	Timestamp  time.Time `json:"timestamp"`
}

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:7070")
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}

	fmt.Println("Client connected to server on: ", conn.LocalAddr())

	// TODO: Implement your client logic here

	for {
		// Read input from the user
		fmt.Print("Enter a message: ")
		input, err := readStdInput()
		if err != nil {
			fmt.Println("Failed to read input:", err)
			continue
		}

		msg := Message{
			SenderID:   1,
			SenderName: "Shree",
			GroupID:    2,
			Content:    input,
			Disconnect: false,
			Timestamp:  time.Now(),
		}

		msgBytes, _ := json.Marshal(msg)

		// Send the message to the server
		_, err = conn.Write(msgBytes)
		if err != nil {
			fmt.Println("Failed to send message: ", input, err)
			continue
		}

		// Check if the user wants to disconnect
		if msg.Disconnect {
			fmt.Println("Client disconnecting from server...")
			break
		}
	}

	if err := conn.Close(); err != nil {
		log.Println("error closing connection: ", err)
	} else {
		fmt.Println("conn disconnected by client")
	}
}

func readStdInput() (string, error) {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return "", err
	}

	return input[:len(input)-1], nil
}
