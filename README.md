# Ecostars Mock Server

This repository contains a mock server built with [Gin](https://github.com/gin-gonic/gin), a fast and lightweight web framework for Go.

## Features

- RESTful API endpoints
- Easy to customize responses
- Fast development and testing

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher

### Installation

```bash
git clone https://github.com/yourusername/gin-quickstart.git
cd gin-quickstart
go mod tidy
```

### Running the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

## Example Endpoints

```http
GET /ping
POST /mock
```

## Customization

Modify `main.go` to add or change endpoints and responses.

## License

MIT