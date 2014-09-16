##websocket-js-repl

A simple REPL for javascript using websockets, written in Go.
This lets you execute javascript in a browser window from the terminal.

###How to run

Make sure you have Go installed and $GOPATH set.

Run `go get github.com/hariharan-uno/websocket-js-repl`

Now, `cd` into the project folder and type `go run serve.go`

Open localhost:8080 in your browser and control the page through js from the terminal.

####Some Interesting commands to get you started

`document.body.style.backgroundColor="#333";`

`document.body.style.color="#f8f8f8";`

`document.body.style.fontFamily="monospace";`
