# fiberlog

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/jenggo/fiberlog)

HTTP request/response logger for [Fiber.v2](https://github.com/gofiber/fiber/v2) using [zerolog](https://github.com/rs/zerolog).

### Install

```sh
go get -u github.com/gofiber/fiber/v2
go get -u github.com/jenggo/fiberlog
```

### Usage

```go
package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/jenggo/fiberlog"
)

func main() {
  app := fiber.New()

  // Default
  app.Use(fiberlog.New())

  // Custom Config
  app.Use(fiberlog.New(fiberlog.Config{
    Logger: &zerolog.New(os.Stdout),
    Next: func(ctx *fiber.Ctx) bool {
      // skip if we hit /private
      return ctx.Path() == "/private"
    },
  }))

  app.Listen(":3000")
}
```

### Example

Run app server:

```sh
$ go run example/main.go
```

Test request:

```sh
$ curl http://localhost:3000/ok
$ curl http://localhost:3000/warn
$ curl http://localhost:3000/err
```

![screen](./example/screen.png)
