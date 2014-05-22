package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func handler(w http.ResponseWriter, r *http.Request) {
	scanner := bufio.NewScanner(os.Stdin)
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	go func() {
		for {
			_, r, err := ws.ReadMessage() //messageType is ignored
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(r))
		}
	}()
	for scanner.Scan() {
		inputcommand := []byte(scanner.Text())
		if err := ws.WriteMessage(1, inputcommand); err != nil {
			log.Println(err)
		}
	}
}

func main() {
	http.HandleFunc("/repl", handler)
	http.ListenAndServe(":6060", nil)

}
