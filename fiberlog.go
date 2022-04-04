package fiberlog

import (
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Config defines the config for logger middleware.
type Config struct {
	// Next defines a function to skip this middleware.
	Next func(ctx *fiber.Ctx) bool

	// Logger is a *zerolog.Logger that writes the logs.
	//
	// Default: log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	Logger *zerolog.Logger
}

// New is a zerolog middleware that allows you to pass a Config.
//
// 	app := fiber.New()
//
// 	// Without config
// 	app.Use(New())
//
// 	// With config
// 	app.Use(New(Config{Logger: &zerolog.New(os.Stdout)}))
func New(config ...Config) fiber.Handler {
	var conf Config
	if len(config) > 0 {
		conf = config[0]
	}

	var sublog zerolog.Logger
	if conf.Logger == nil {
		sublog = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		sublog = *conf.Logger
	}

	return func(c *fiber.Ctx) error {
		// Don't execute the middleware if Next returns true
		if conf.Next != nil && conf.Next(c) {
			return c.Next()
		}

		start := time.Now()

		// handle request
		msg := "Request"
		if err := c.Next(); err != nil {
			msg = err.Error()
		}

		code := c.Response().StatusCode()

		var ip string
		// ips := c.IPs()
		ips := c.Get("Cf-Connecting-Ip") // Behind Cloudflare
		if len(ips) == 0 {
			ip = c.IP()
		} else {
			ip = ips
		}

		dumplogger := sublog.With().
			Int("status", code).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("ip", ip).
			Str("latency", time.Since(start).String()).
			Str("user-agent", c.Get(fiber.HeaderUserAgent)).
			Logger()

		switch {
		case code >= fiber.StatusBadRequest && code < fiber.StatusInternalServerError:
			dumplogger.Warn().Msg(msg)
		case code >= http.StatusInternalServerError:
			dumplogger.Error().Msg(msg)
		default:
			dumplogger.Info().Msg(msg)
		}

		return nil
	}
}
