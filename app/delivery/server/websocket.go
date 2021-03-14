package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
}

type Hub struct {
	clients map[*Client]bool
}

func Hello(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	for {
		messageType, message, err := connection.ReadMessage()

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Mensaje: %s", message)

		err = connection.WriteMessage(messageType, []byte("Hola desde el servidor"))
		if err != nil {
			log.Fatal(err)
		}

	}
}
