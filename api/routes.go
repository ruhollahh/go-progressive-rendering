package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/ruhollahh/go-progressive-rendering/api/handler"
)

func (a *API) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./web/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	handler := handler.New(a.Logger, a.Services)

	router.Handler(http.MethodGet, "/product/:id", a.setNonce(handler.ShowProduct))

	standard := alice.New(a.recoverPanic, a.logRequest, a.secureHeaders)
	return standard.Then(router)
}
