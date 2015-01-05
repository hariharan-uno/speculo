##speculo

`speculo` is a simple REPL for javascript using websockets, written in Go.
This lets you execute javascript in a browser window from the terminal.

###How to run

Make sure you have Go installed and $GOPATH set.

```sh

$ go get github.com/hariharan-uno/speculo
$ speculo

```

Open localhost:8080 in your browser and control the web page through JavaScript from the terminal.

####Some Interesting commands to get you started

```javascript
document.body.style.backgroundColor="#333";

document.body.style.color="#f8f8f8";

document.body.style.fontFamily="monospace";
```

####Trivia
`speculo` means mirror in latin.