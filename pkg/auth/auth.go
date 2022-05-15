package auth

import (
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

type jsonData struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
}

type Message struct {
	jsonData
	Params params `json:"params"`
}

type params struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func Auth(conn *websocket.Conn) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Please, check if you have it")
	}

	req := Message{
		jsonData{
			Jsonrpc: "2.0",
			Method:  "public/auth",
		},
		params{
			GrantType:    "client_credentials",
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
		},
	}
	var resp Message
	conn.WriteJSON(req)
	error := conn.ReadJSON(&resp)
	if error != nil {
		panic(err)
	}
}
