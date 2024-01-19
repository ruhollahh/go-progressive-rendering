package api

import (
	"log/slog"

	"github.com/ruhollahh/go-progressive-rendering/internal/service"
)

type Config struct {
	Port int
	Env  string
}

type API struct {
	Services service.Services
	Logger   *slog.Logger
	Config   Config
}
