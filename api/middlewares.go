package api

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ruhollahh/go-progressive-rendering/api/contextutil"
	"github.com/ruhollahh/go-progressive-rendering/api/httperrors"
)

func (a *API) secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'")

		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

func (a *API) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			proto  = r.Proto
			method = r.Method
			uri    = r.URL.RequestURI()
		)

		a.Logger.Info("received request", "ip", ip, "proto", proto, "method", method, "uri", uri)

		next.ServeHTTP(w, r)
	})
}

func (a *API) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				httperrors.ServerError(a.Logger, w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (a *API) setNonce(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var buf [16]byte

		_, err := io.ReadFull(rand.Reader, buf[:])
		if err != nil {
			panic("CSP Nonce rand.Reader failed" + err.Error())
		}

		nonce := base64.RawStdEncoding.EncodeToString(buf[:])
		ctx := contextutil.ContextSetNonce(r.Context(), nonce)
		r = r.WithContext(ctx)

		csp := []string{
			"default-src 'self'",
			fmt.Sprintf("script-src 'self' 'nonce-%s'", nonce),
			fmt.Sprintf("style-src 'self' 'nonce-%s'", nonce),
		}
		h := w.Header()
		h.Set("Content-Security-Policy", strings.Join(csp, "; "))

		next.ServeHTTP(w, r)
	})
}
