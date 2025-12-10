package handlers

import (
"net/http"
)

var homeHTML = []byte(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>WebSocket Chat</title>
<style>
body {
    font-family: Arial, sans-serif;
    max-width: 800px;
    margin: 50px auto;
    padding: 20px;
}
#chat-log {
    border: 1px solid #ccc;
    height: 400px;
    overflow-y: scroll;
    padding: 10px;
    margin-bottom: 20px;
    background-color: #f9f9f9;
}
#message-form {
    display: flex;
    gap: 10px;
}
#message-input {
    flex: 1;
    padding: 10px;
    font-size: 16px;
}
#send-button {
    padding: 10px 20px;
    font-size: 16px;
    background-color: #007bff;
    color: white;
    border: none;
    cursor: pointer;
}
#send-button:hover {
    background-color: #0056b3;
}
.message {
    margin: 5px 0;
    padding: 5px;
}
</style>
</head>
<body>
<h1>WebSocket Chat</h1>
<div id="chat-log"></div>
<form id="message-form">
    <input id="message-input" type="text" placeholder="Type a message..." autocomplete="off">
    <button id="send-button" type="submit">Send</button>
</form>

<script>
(function() {
    var conn;
    var log = document.getElementById("chat-log");
    var form = document.getElementById("message-form");
    var input = document.getElementById("message-input");

    function appendLog(msg) {
        var d = document.createElement("div");
        d.className = "message";
        d.textContent = msg;
        log.appendChild(d);
        log.scrollTop = log.scrollHeight;
    }

    function connect() {
        var wsProtocol = window.location.protocol === "https:" ? "wss:" : "ws:";
        conn = new WebSocket(wsProtocol + "//" + window.location.host + "/ws");
        
        conn.onopen = function(evt) {
            appendLog("Connected to server");
        };
        
        conn.onclose = function(evt) {
            appendLog("Connection closed");
            setTimeout(connect, 1000);
        };
        
        conn.onmessage = function(evt) {
            appendLog("Received: " + evt.data);
        };
        
        conn.onerror = function(evt) {
            appendLog("Error: " + evt.data);
        };
    }

    connect();

    form.onsubmit = function(evt) {
        evt.preventDefault();
        if (!conn || conn.readyState !== WebSocket.OPEN) {
            appendLog("Not connected");
            return false;
        }
        if (!input.value) {
            return false;
        }
        conn.send(input.value);
        appendLog("Sent: " + input.value);
        input.value = "";
        return false;
    };
})();
</script>
</body>
</html>
`)

func ServeHome(w http.ResponseWriter, r *http.Request) {
if r.URL.Path != "/" {
http.Error(w, "Not found", http.StatusNotFound)
return
}
if r.Method != http.MethodGet {
http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
return
}
w.Header().Set("Content-Type", "text/html; charset=utf-8")
w.Write(homeHTML)
}
