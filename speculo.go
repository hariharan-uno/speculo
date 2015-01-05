// Copyright 2014 Hari haran. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func speculoHandler(w http.ResponseWriter, r *http.Request) {
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
			_, r, err := ws.ReadMessage() // messageType is ignored
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
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.Dir(path)))
	http.HandleFunc("/repl", speculoHandler)

	fmt.Println("Open http://localhost:8080 and type js commands in the terminal!")
	http.ListenAndServe(":8080", nil)
}
