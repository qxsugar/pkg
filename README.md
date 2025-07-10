# QxSugar Pkg

[![Tests](https://github.com/qxsugar/pkg/actions/workflows/test.yml/badge.svg)](https://github.com/qxsugar/pkg/actions/workflows/test.yml)
[![Auto Release](https://github.com/qxsugar/pkg/actions/workflows/auto-release.yml/badge.svg)](https://github.com/qxsugar/pkg/actions/workflows/auto-release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/qxsugar/pkg)](https://goreportcard.com/report/github.com/qxsugar/pkg)
[![codecov](https://codecov.io/gh/qxsugar/pkg/branch/main/graph/badge.svg)](https://codecov.io/gh/qxsugar/pkg)
[![GoDoc](https://godoc.org/github.com/qxsugar/pkg?status.svg)](https://godoc.org/github.com/qxsugar/pkg)
[![Release](https://img.shields.io/github/v/release/qxsugar/pkg)](https://github.com/qxsugar/pkg/releases)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.23-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

A comprehensive Go utility library for web applications built with the Gin framework, providing standardized error handling, database utilities, logging, and more.

## Features

- 🚀 **HTTP Handler Wrappers** - Standardized error handling for Gin framework
- 🔥 **Business Error Management** - Structured error system following Google's API Design Guide
- 📊 **JSON Database Fields** - Custom JSON type for database storage
- 📝 **Logging Utilities** - Production and development logger configurations with Zap
- 🔗 **Chain Pattern** - Alice-style chain-of-responsibility pattern for error handling
- 🔒 **Synchronization Utilities** - Thread-safe utilities and helpers
- 🛠 **SQL Helpers** - Database query utilities and helpers
- ✅ **100% Test Coverage** - Comprehensive test suite

## Installation

```bash
go get github.com/qxsugar/pkg
```

## Quick Start

### HTTP Handlers

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/qxsugar/pkg/kit"
)

func main() {
    r := gin.Default()
    
    // Use kit's RouterGroup wrapper
    api := kit.NewRouterGroup(r.Group("/api"))
    
    api.GET("/users/:id", func(ctx *gin.Context) (any, error) {
        id := ctx.Param("id")
        if id == "" {
            return nil, kit.NewInvalidArgumentError()
        }
        
        user := map[string]any{
            "id":   id,
            "name": "John Doe",
        }
        return user, nil
    })
    
    r.Run(":8080")
}
```

### Error Handling

```go
// Business errors with structured responses
func GetUser(id string) (any, error) {
    if id == "" {
        return nil, kit.NewInvalidArgumentError()
    }
    
    user, err := database.FindUser(id)
    if err != nil {
        return nil, kit.NewNotFoundError().WithErr(err)
    }
    
    return user, nil
}
```

### Chain Pattern

```go
// Error-safe operation chaining
result := kit.New(
    func() error { return connectDB() },
    func() error { return beginTransaction() },
    func() error { return insertData() },
    func() error { return commitTransaction() },
).Error()

if result != nil {
    log.Printf("Operation failed: %v", result)
}
```

### JSON Database Fields

```go
type User struct {
    ID       int64    `json:"id"`
    Name     string   `json:"name"`
    Metadata kit.JSON `json:"metadata" gorm:"type:jsonb"`
}

// The JSON field handles database scanning and marshaling automatically
```

### Logging

```go
// Production logger
logger := kit.MustProduction()
logger.Info("Service started", zap.Int("port", 8080))

// Development logger
devLogger := kit.MustDevelopment()
devLogger.Debug("Debug message", zap.String("component", "auth"))
```

## Error Codes

The library follows Google's API Design Guide for error codes:

| Code | HTTP | Description |
|------|------|-------------|
| `OK` | 200 | No error |
| `ErrInvalidArgument` | 400 | Invalid request parameters |
| `ErrUnauthenticated` | 401 | Authentication required |
| `ErrPermissionDenied` | 403 | Insufficient permissions |
| `ErrNotFound` | 404 | Resource not found |
| `ErrAlreadyExists` | 409 | Resource already exists |
| `ErrInternal` | 500 | Internal server error |

## Testing

```bash
# Run all tests
make test

# Run tests with coverage
go test -v -cover ./...

# Run specific package tests
go test -v ./kit
```

## Development

### Requirements

- Go 1.23 or later
- Make (optional)

### Building

```bash
go build ./...
```

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## CI/CD

This project uses GitHub Actions for continuous integration and deployment:

- **Tests** (`.github/workflows/test.yml`) - Runs on every push and PR
  - Tests against Go 1.23 and 1.24
  - Enforces minimum 90% test coverage
  - Runs golangci-lint for code quality
  - Uploads coverage reports to Codecov
- **Auto Release** (`.github/workflows/auto-release.yml`) - Automatic patch version releases on main branch pushes
- **Manual Release** (`.github/workflows/manual-release.yml`) - Support for manual version bumping (patch/minor/major)

All workflows must pass before any code can be merged or released.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](./LICENSE) file for details.

## Acknowledgments

- Inspired by Google's API Design Guide for error handling
- Built on top of the excellent [Gin](https://github.com/gin-gonic/gin) framework
- Uses [Zap](https://github.com/uber-go/zap) for high-performance logging
