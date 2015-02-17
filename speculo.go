// Copyright 2014 Hari haran. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh/terminal"
)

func speculoHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Fatal(err)
		return
	}

	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, oldState)
	term := terminal.NewTerminal(os.Stdin, "> ")

	term.AutoCompleteCallback = func(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
		if key == '\x03' { // Ctrl-C keycode is \x03
			fmt.Println()
			if oldState != nil {
				terminal.Restore(0, oldState)
			}
			os.Exit(0)
		}
		return "", 0, false
	}

	go func() {
		newline := []byte("\n")
		for {
			_, r, err := ws.ReadMessage() // messageType is ignored
			if err != nil {
				log.Fatal(err)
			}
			term.Write(r)
			term.Write(newline)
		}
	}()

	for {
		text, err := term.ReadLine()
		if err != nil {
			// Ctrl-D is EOF. So, quit without printing stacktrace.
			if err == io.EOF {
				fmt.Println()
				if oldState != nil {
					terminal.Restore(0, oldState)
				}
				os.Exit(0)
			}
			panic(err)
		}
		inputcommand := []byte(text)
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
