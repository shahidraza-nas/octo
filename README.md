# Octo WebSocket Chat

A real-time WebSocket chat application built with Go and the Gorilla WebSocket library.

## Features

- Real-time bidirectional communication using WebSockets
- Hub-based architecture for broadcasting messages to multiple clients
- Automatic reconnection on connection loss
- Simple web interface for testing

## Project Structure

```
.
├── main.go                      # Application entry point
├── internal/
│   ├── hub/
│   │   └── hub.go              # Hub for managing client connections and broadcasting
│   ├── client/
│   │   └── client.go           # WebSocket client implementation
│   └── handlers/
│       ├── websocket.go        # WebSocket upgrade handler
│       └── home.go             # Home page handler
├── go.mod                       # Go module definition
└── go.sum                       # Go module checksums
```

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

```bash
# Install dependencies
go mod download
```

### Running the Application

```bash
# Run the server
go run main.go
```

The server will start on `http://localhost:8080`

### Development

```bash
# Build the application
go build -o octo-chat

# Run tests
go test ./...
```

## Usage

1. Open your browser and navigate to `http://localhost:8080`
2. Open multiple browser tabs to simulate multiple clients
3. Type messages in any tab and see them broadcast to all connected clients

## Architecture

### Hub Pattern

The application uses a hub pattern where:
- The **Hub** manages all active client connections
- Clients send messages to the hub
- The hub broadcasts messages to all connected clients
- Each client runs two goroutines: one for reading and one for writing

### Message Flow

1. Client sends message → ReadPump → Hub.broadcast
2. Hub receives message → broadcasts to all clients
3. WritePump sends message → Client receives message

## API Endpoints

- `GET /` - Serves the web interface
- `GET /ws` - WebSocket upgrade endpoint

## Configuration

Key constants in `internal/client/client.go`:
- `writeWait`: Time allowed to write a message (10s)
- `pongWait`: Time allowed to read pong message (60s)
- `pingPeriod`: Send pings with this period (54s)
- `maxMessageSize`: Maximum message size (512 bytes)

## License

MIT
