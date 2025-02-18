# Go ServeMux Demo API

A simple HTTP REST API implementation using Go's standard `net/http` package and `ServeMux` for routing. This project demonstrates the basic usage of Go's built-in HTTP server capabilities without external routing libraries.

## Overview

This project showcases:
- Native Go HTTP server implementation
- Request routing with `ServeMux`
- RESTful API design principles
- In-memory data storage
- Basic CRUD operations for user management

## Project Structure

```
.
├── cmd
│   └── server        # Main application entry point
├── internal
│   ├── handlers      # HTTP request handlers
│   ├── models        # Data models
│   └── storage       # Data storage implementation
└── go.mod            # Go module definition
```

## Prerequisites

- Go 1.24 or higher

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/oragazzo/go-native-router-example.git
   ```

2. Navigate to the project directory:
   ```bash
   cd go-native-router-example
   ```

3. Run the server:
   ```bash
   go run cmd/server/main.go
   ```

The server will start on `localhost:8080` by default.

## API Endpoints

The following endpoints are available:

- `GET /users` - List all users
- `GET /users/{id}` - Get a specific user
- `POST /users` - Create a new user
- `PUT /users/{id}` - Update an existing user
- `DELETE /users/{id}` - Delete a user

## Why ServeMux?

`ServeMux` is Go's standard HTTP request multiplexer. While there are more feature-rich third-party routers available, `ServeMux` is:
- Simple and lightweight
- Part of the standard library
- Perfect for basic routing needs
- Great for learning HTTP server fundamentals in Go

## Contributing

Feel free to submit issues and enhancement requests! 