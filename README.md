# Own HTTP Protocol

A simple HTTP/1.1-like protocol implementation written from scratch in Go.

The goal of this project is to understand how HTTP works at a lower level by implementing request parsing, response serialization, headers, bodies, and TCP communication without using Go's `net/http` package.

## Features

- TCP-based communication
- HTTP/1.1-style Request-Line parsing
- HTTP method parsing
- Path parsing
- Header parsing
- `Content-Length` support
- Request body parsing
- Response status line generation
- Response headers
- Response body handling
- Response serialization
- Basic `curl` compatibility

## Architecture

```text
Client (curl)
      │
      │ TCP
      ▼
┌───────────────┐
│   TCP Server  │
└───────┬───────┘
        │
        ▼
┌──────────────────┐
│  BufferedReader  │
│                  │
│  ReadLine()      │
│  ReadN()         │
└────────┬─────────┘
         │
         ▼
┌──────────────────┐
│      Parser      │
│                  │
│ Request-Line     │
│ Headers          │
│ Body             │
└────────┬─────────┘
         │
         ▼
      Request
         │
         │
         ▼
      Response
         │
         ▼
     Serialize()
         │
         ▼
      []byte
         │
         ▼
     TCP Write()
         │
         ▼
       curl
```

## Project Structure

```text
.
├── README.md
├── go.mod
├── main.go
├── request
│   ├── body
│   │   └── body.go
│   ├── header
│   │   └── header.go
│   ├── identifier
│   │   └── readline.go
│   ├── request-line
│   │   └── request-line.go
│   └── request.go
└── response
    ├── response.go
    ├── setGenerateHeader
    │   └── setGenerateHeader.go
    └── status-line
        └── status-line.go
```

## Request

A request contains three main components:

```text
Request
├── Request-Line
├── Headers
└── Body
```

Example:

```http
POST /users HTTP/1.1
Host: example.com
Content-Length: 5

hello
```

The request parser reads the TCP byte stream and builds a structured `Request`.

## Response

A response contains:

```text
Response
├── Status-Line
├── Headers
└── Body
```

Example:

```http
HTTP/1.1 200 OK
Content-Length: 5

hello
```

The response is then serialized into bytes before being written to the TCP connection.

## Example

Start the server:

```bash
go run .
```

Then send a request using `curl`:

```bash
curl -v -X POST http://localhost:8000/users \
  -H "Host: example.com" \
  -H "Content-Length: 5" \
  -d "hello"
```

The server responds with:

```text
HTTP/1.1 200 OK
Content-Length: 5

hello
```

## Core Concepts

This project focuses on understanding the fundamentals behind HTTP and TCP communication.

### TCP is a byte stream

TCP does not preserve HTTP message boundaries. Data can arrive in fragments or multiple messages can arrive in a single read.

The project therefore uses a buffered reader to handle:

- Reading complete lines
- Reading an exact number of bytes
- Preserving unread bytes

### Content-Length

The request body is read according to the `Content-Length` header.

For example:

```text
Content-Length: 5

hello
```

The parser reads exactly five bytes for the body.

### Serialization

A `Response` is converted into bytes before being sent over TCP:

```text
Response
   ↓
Serialize()
   ↓
[]byte
   ↓
conn.Write()
```

## Design Goals

The main goal is educational: implement the basic mechanics of HTTP from the ground up and understand the relationship between HTTP and TCP.

The project intentionally avoids Go's `net/http` package.

## Current Limitations

This is a learning project and is not intended to be a production HTTP server.

Current limitations include:

- Limited HTTP methods
- HTTP/1.1 only
- Basic header validation
- `Content-Length` based request bodies
- Single TCP connection handling
- No HTTP keep-alive implementation
- No chunked transfer encoding
- No TLS
- No routing system
- No production-level error handling

## Status

🚧 Work in progress.

The project is being developed incrementally to explore HTTP, TCP, parsing, buffering, and serialization in Go.

## License

MIT License
