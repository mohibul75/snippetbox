package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServe := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", fileServe)

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snnipet/view", app.snippetView)
	mux.HandleFunc("/snnipet/create", app.snippetCreate)

	return mux
}
