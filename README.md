# Own HTTP Protocol

A simple HTTP/1.1-like server built from scratch in Go, without using `net/http`.

The goal of this project is to understand what happens underneath an HTTP server by building the main pieces directly on top of TCP.

## 🎯 Goal

Instead of hiding HTTP behind a high-level framework, this project builds the protocol step by step:

```text
TCP
 ↓
Buffered Reader
 ↓
HTTP Request Parser
 ↓
Router
 ↓
Handler
 ↓
HTTP Response
 ↓
TCP
```

The focus is learning how HTTP works internally rather than building a production-ready web framework.

## ✨ Features

### TCP Server

The server uses Go's `net` package directly to create a TCP listener and accept connections.

```text
Client
   │
   │ TCP
   ▼
Server
```

### Request Parsing

The server parses HTTP requests including:

- Request-Line
- Headers
- Request Body
- `Content-Length`

Example:

```http
POST /users HTTP/1.1
Host: localhost:8000
Content-Length: 5

hello
```

### Buffered Reading

TCP is a byte stream, so the project includes a buffered reader to correctly handle data boundaries.

It supports:

- Reading lines
- Reading an exact number of bytes
- Preserving extra bytes for the next request

This becomes especially important when multiple HTTP requests are sent over the same TCP connection.

### Router

The project includes a simple HTTP router with method and path matching.

Supported methods:

- GET
- POST
- PUT
- PATCH
- DELETE

Example:

```go
router.GET("/", handler)
router.POST("/users", handler)
```

### Dynamic Routes

Routes can contain parameters:

```text
/users/:id
```

A request such as:

```text
/users/42
```

produces:

```text
id = 42
```

Route parameters are available through the request:

```go
request.Params
```

### Handler

Handlers are represented as functions:

```go
type IHandler func(request request.Request) (response.Response, error)
```

This keeps request handling simple and allows the router to dispatch requests to user-defined handlers.

### 404 / 405

The router distinguishes between:

- `404 Not Found` — no matching path
- `405 Method Not Allowed` — path exists but the HTTP method is not supported

### Concurrent Connections

Each accepted TCP connection is handled independently using a goroutine.

```text
Client A ─────┐
Client B ─────┼──> Server
Client C ─────┘
```

Conceptually:

```go
go handleConnection(connection, router)
```

### Persistent Connections

A single TCP connection can handle multiple HTTP requests:

```text
TCP Connection
      │
      ├── Request 1 → Response 1
      ├── Request 2 → Response 2
      ├── Request 3 → Response 3
      └── Request 4 → Response 4
```

The server keeps parsing requests until the connection is closed.

The client can explicitly request connection closure with:

```http
Connection: close
```

The server also handles a clean `io.EOF` when the client closes the connection.

## 📁 Project Structure

```text
.
├── handler/
│   └── handler.go
├── header/
│   └── ...
├── params/
│   └── params.go
├── request/
│   └── ...
├── requestline/
│   └── ...
├── response/
│   └── ...
├── router/
│   └── ...
├── server/
│   └── ...
├── statusline/
│   └── ...
└── main.go
```

The project is intentionally split into small packages so each part of the HTTP protocol can be understood independently.

## 🚀 Running

Clone the repository:

```bash
git clone https://github.com/moeinht/http-protocol.git
cd http-protocol
```

Run the server:

```bash
go run .
```

The server currently listens on port `8000`.

You can test it with:

```bash
curl -v http://localhost:8000/
```

## 🧪 Testing Persistent Connections

For testing multiple requests over the same TCP connection, `nc` can be useful:

```bash
nc -v localhost 8000
```

Then send multiple HTTP requests without closing the connection:

```http
GET / HTTP/1.1
Host: localhost:8000

GET / HTTP/1.1
Host: localhost:8000

```

This allows you to observe multiple HTTP request/response cycles over a single TCP connection.

## 📌 Current Scope

The project currently focuses on the fundamentals of HTTP/1.1-style communication over TCP.

Implemented:

- TCP listener
- HTTP request parsing
- Request-Line parsing
- Header parsing
- `Content-Length`
- Request body reading
- Response serialization
- Router
- Static routes
- Dynamic routes
- Handler abstraction
- Concurrent connections
- Persistent connections
- `Connection: close`
- Basic `404` / `405` / `400` / `500` responses

## 🛣️ Roadmap

Possible future improvements:

- Header validation and case-insensitive header handling
- Better HTTP error handling
- Connection timeouts
- Middleware
- More complete HTTP/1.1 semantics
- Chunked Transfer Encoding
- Improved routing precedence
- More comprehensive protocol tests

## 📚 Why Build This?

The goal isn't to replace Go's `net/http`.

The goal is to understand what happens underneath it.

Building the server from raw TCP makes concepts like:

- TCP streams
- buffering
- message framing
- connection lifecycle
- request parsing
- routing
- persistent connections

much easier to understand.

---

Built with Go 🐹

Learning HTTP by building it from scratch.
