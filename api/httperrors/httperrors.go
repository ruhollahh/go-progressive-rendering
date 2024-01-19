package httperrors

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

func ServerError(logger *slog.Logger, w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	logger.Error(err.Error(), "uri", uri, "method", method, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func NotFoundError(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}
