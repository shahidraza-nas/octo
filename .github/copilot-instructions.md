# Copilot Instructions for Octo WebSocket Chat

## Project Overview

This is a real-time WebSocket chat application built with Go using the Gorilla WebSocket library. The architecture follows a hub-and-spoke pattern where a central hub manages all client connections and broadcasts messages.

## Architecture

### Hub-Based Broadcasting Pattern
- **Hub** (`internal/hub/hub.go`): Central coordinator managing client lifecycle and message broadcasting
  - Uses Go channels for concurrent message handling
  - Three main channels: `register`, `unregister`, `broadcast`
  - Single event loop in `Run()` method handles all hub operations
  
- **Client** (`internal/client/client.go`): Represents a WebSocket connection
  - Each client runs two goroutines: `ReadPump()` and `WritePump()`
  - `ReadPump()`: Reads messages from WebSocket and sends to hub
  - `WritePump()`: Receives messages from hub and writes to WebSocket
  - Uses ping/pong for connection health monitoring

### Message Flow
1. Client sends message → `ReadPump()` → Hub's broadcast channel
2. Hub receives from broadcast → sends to all clients' `Send` channels
3. Each client's `WritePump()` receives from `Send` → writes to WebSocket

## Key Patterns & Conventions

### Goroutine Management
- Each WebSocket connection spawns exactly 2 goroutines (read and write pumps)
- Always defer connection cleanup and channel closure
- Hub runs in a separate goroutine started from `main()`

### Channel Usage
- Client's `Send` channel is buffered (256 messages) to prevent blocking
- Hub channels are unbuffered for immediate processing
- Close channels only in defer statements to ensure cleanup

### Error Handling
- Check for unexpected close errors using `websocket.IsUnexpectedCloseError()`
- Log errors but don't panic - clean up connections gracefully
- Use timeouts (`writeWait`, `pongWait`) to prevent hanging connections

### WebSocket Configuration
- `CheckOrigin` allows all origins in development (change for production)
- Message size limited to 512 bytes (configurable in `client.go`)
- Ping interval is 54s (90% of 60s pong wait time)

## Development Workflow

### Building & Running
```bash
go run main.go              # Run directly
go build -o octo-chat      # Build binary
```

### Testing
- Test WebSocket by opening multiple browser tabs to `http://localhost:8080`
- Messages should broadcast to all connected clients

### Adding Features
- New message types: Modify `client.ReadPump()` to handle different message formats
- Authentication: Add middleware in `handlers.ServeWs()` before upgrade
- Rooms/Channels: Extend hub with multiple broadcast maps keyed by room ID

## Project Structure
- `main.go`: Entry point, sets up HTTP routes and starts hub
- `internal/hub/`: Hub implementation for client management
- `internal/client/`: WebSocket client with read/write pumps
- `internal/handlers/`: HTTP handlers for WebSocket upgrade and home page

## Dependencies
- `github.com/gorilla/websocket`: WebSocket implementation
- Standard library only for HTTP server

## Common Gotchas
- Don't write to WebSocket from multiple goroutines - use the `WritePump()` pattern
- Always set deadlines before WebSocket operations to prevent hangs
- Client cleanup must close the `Send` channel to terminate `WritePump()`
- Hub must not send to closed `Send` channels - check map membership first
