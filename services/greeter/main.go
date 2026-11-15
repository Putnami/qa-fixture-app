package main

import (
	"errors"
	"os"

	"go.putnami.dev/app"
	pconfig "go.putnami.dev/config"
	"go.putnami.dev/http"
	"go.putnami.dev/logger"
	"go.putnami.dev/platform"

	"github.com/putnami/qa-fixture-app/libs/greeting"
)

type ServerConfig struct {
	Port int `json:"port" default:"3000" env:"PORT"`
}

var RuntimeConfig = pconfig.Config[ServerConfig]("server")

func main() {
	cfg, err := pconfig.Load(RuntimeConfig)
	if err != nil {
		logger.Default().Error("config failed", err)
		os.Exit(1)
	}

	a, _ := newApp(cfg)
	if err := a.ListenAndServe(); err != nil {
		logger.Default().Error("application failed", err)
		os.Exit(1)
	}
}

// newApp composes the application and returns it with its HTTP server. The
// platform plugin answers /livez, /healthz, /readyz and /version on that
// server, and the health plugin answers GET /_/health.
func newApp(cfg ServerConfig) (*app.Application, *http.ServerPlugin) {
	server := http.NewServerPlugin(http.ServerConfig{Port: cfg.Port})
	server.Use(http.Recovery())
	server.Use(http.RequestID())
	server.Use(http.Logging(http.LoggerOptions{
		Exclude: []string{"/_/health", "/livez", "/healthz", "/readyz"}, // keep probes out of the access log
	}))

	server.GET("/", func(ctx *http.Context) *http.Response {
		return http.JSON(map[string]string{"Hello": "World"})
	})
	server.GET("/greet/{name}", greetHandler)

	a := app.New("server")
	a.Use(server)
	a.Use(platform.NewPlugin(platform.Config{}))
	a.Use(http.NewHealthPlugin())
	return a, server
}

// greetHandler answers GET /greet/{name} with {"message": "Hello, <name>!"},
// or 400 when the name is empty or too long.
func greetHandler(ctx *http.Context) *http.Response {
	message, err := greeting.Greet(ctx.Param("name"))
	if errors.Is(err, greeting.ErrEmptyName) || errors.Is(err, greeting.ErrNameTooLong) {
		return http.JSONStatus(400, map[string]string{"error": err.Error()})
	}
	if err != nil {
		return http.ErrorResponse(err)
	}
	return http.JSON(map[string]string{"message": message})
}
