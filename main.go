package main

import (
	"deribee/pkg/auth"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Please, check if you have it")
	}

	conn, _, err := websocket.DefaultDialer.Dial(os.Getenv("URL"), nil)
	if err != nil {
		log.Fatal(err)
	}

	auth.Auth(conn)
}
