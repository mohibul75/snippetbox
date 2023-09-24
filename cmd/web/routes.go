package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServe := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServe))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	middlewareChain := alice.New(app.panicRecover, app.logRequest, secureHeaders)

	return middlewareChain.Then(mux)
}
