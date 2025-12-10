# Quick Start Guide

## Prerequisites

- Go 1.21 or higher installed
- A web browser

## Getting Started

### 1. Install Dependencies

```bash
go mod download
```

### 2. Run the Server

```bash
go run main.go
```

Or build and run:

```bash
go build -o octo-chat
./octo-chat
```

### 3. Test the Application

1. Open your web browser and navigate to: **<http://localhost:8080>**
2. You should see a chat interface with a message input box
3. Open another browser tab/window to **<http://localhost:8080>**
4. Type a message in one tab and press "Send"
5. You should see the message appear in all open tabs!

## What's Happening?

- The server creates a WebSocket connection for each browser tab
- When you send a message, it goes through the WebSocket to the server
- The server's **Hub** broadcasts the message to all connected clients
- Each browser receives and displays the broadcast message

## Next Steps

- Check out the **README.md** for full documentation
- Read **.github/copilot-instructions.md** for architecture details
- Explore the code in `internal/` to understand the implementation

## Common Commands

```bash
# Run the server
go run main.go

# Build binary
go build -o octo-chat

# Run tests
go test ./...

# Clean build artifacts
rm -f octo-chat

# View dependencies
go mod graph
```

## Troubleshooting

**Port 8080 already in use?**

- Kill the process: `lsof -ti:8080 | xargs kill -9`
- Or modify the port in `main.go`

**Connection refused?**

- Ensure the server is running
- Check that port 8080 is not blocked by a firewall

Enjoy your WebSocket chat! 🚀
