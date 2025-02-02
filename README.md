# httpcat

**httpcat** is a Go library that returns a "cat image" corresponding to a specific HTTP status code (e.g., 404, 500). It leverages Go's [embed](https://pkg.go.dev/embed) package to include images at build time, eliminating the need for distributing separate image files.

## Features

- **Easy to Use**: Simply call `SendError(w, code)` to return a cat image corresponding to the HTTP status code.
- **Framework Independent**: Uses only the standard Go `net/http` library, making it easy to integrate with Gin, Echo, Fiber, Chi, or any other web framework.
- **Embed-Based**: Uses Go 1.16+'s `//go:embed` feature to embed image files into the binary.

## Installation

```bash
go get github.com/Feralthedogg/httpcat
```

> Requires Go 1.16 or later to use the `embed` package.

## Directory Structure

```
httpcat/
├── go.mod
├── go.sum
├── pkg/
│   └── httpcat/
│       ├── httpcat.go
│       └── assets/
│           ├── 404.jpg
│           ├── 500.jpg
│           └── ... (other httpcat images)
└── cmd/
    └── main.go
```

## Using with Gin Framework
```go
package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Feralthedogg/httpcat/pkg/httpcat"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        codeParam := c.Query("code")
        if codeParam == "" {
            c.String(http.StatusOK, "Hello from Gin! Try ?code=404")
            return
        }

        code, err := strconv.Atoi(codeParam)
        if err != nil {
            c.String(http.StatusBadRequest, "Invalid code param")
            return
        }

        if code >= 400 && code <= 599 {
            httpcat.SendError(c.Writer, code)
            return
        }

        c.String(http.StatusOK, "Non-error code: %d", code)
    })

    r.Run(":8080")
}
```
### Testing

1. Run the server:
   ```bash
   go run ./cmd/main.go
   ```
2. Open your browser and visit:
   - `http://localhost:8080` → Default greeting message
   - `http://localhost:8080?code=404` → 404 cat image
   - `http://localhost:8080?code=500` → 500 cat image


## License

[MIT](LICENSE)
