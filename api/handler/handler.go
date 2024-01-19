package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/ruhollahh/go-progressive-rendering/internal/service"
)

type Handler struct {
	Logger   *slog.Logger
	Services service.Services
}

func New(logger *slog.Logger, services service.Services) *Handler {
	return &Handler{
		logger,
		services,
	}
}

func (h *Handler) getIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}
