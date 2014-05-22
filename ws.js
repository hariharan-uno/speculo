output = document.getElementById("output");
ws = new WebSocket("ws://localhost:6060/repl");

// Listen for the connection open event then call the sendMessage function
ws.onopen = function(e) {
    log("Connected");
}

// Listen for the close connection event
ws.onclose = function(e) {
    log("Disconnected: " + e.reason);
}
 
// Listen for connection errors
ws.onerror = function(e) {
    log("Error ");
}
 
// Listen for new messages arriving at the client
ws.onmessage = function(e) {
    log("Command received: " + e.data);
    try {
	var result = eval(e.data);
	ws.send(result);
	log("Result: "+result);
    } catch(err) {
	log(err);
    }
}

// Display logging information in the document.
function log(s) {
    var p = document.createElement("p");
    p.style.wordWrap = "break-word";
    p.textContent = s;
    output.appendChild(p);
    
    // Also log information on the javascript console
    console.log(s);
}
