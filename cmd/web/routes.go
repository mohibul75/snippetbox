package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.NoFound(w)
	})

	fileServe := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServe))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view", app.snippetView)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreate)

	middlewareChain := alice.New(app.panicRecover, app.logRequest, secureHeaders)

	return middlewareChain.Then(router)
}
