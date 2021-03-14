package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	//	style := ui.DefaultStyle()
	//	screen, err := ui.NewDefaultScreen(style)
	//	if err != nil {
	//		os.Exit(1)
	//	}
	//
	//	world, err := file.LoadWorld(file.Filename)
	//
	//	if err != nil {
	//		os.Exit(1)
	//	}
	//	ui.NewWorldView(screen, world).Run()
	Network()
}

//Network

func Network() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	urlServer := url.URL{
		Scheme: "ws",
		Host:   "localhost:8000",
		Path:   "/",
	}

	connection, _, err := websocket.DefaultDialer.Dial(urlServer.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()

	done := make(chan struct{})

	//Print message from the server
	go func() {
		defer close(done)
		for {
			_, message, err := connection.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Message from the server: %s", message)
		}
	}()

	ticket := time.NewTicker(time.Second)
	defer ticket.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticket.C:
			err := connection.WriteMessage(websocket.TextMessage, []byte("Hola desde el client"))
			if err != nil {
				log.Fatal(err)
			}
		case <-interrupt:
			err := connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Fatal(err)
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}
