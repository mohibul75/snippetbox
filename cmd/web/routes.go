package main

import (
	"net/http"
)

func (app * application) routes() http.Handler{

	mux:= http.NewServeMux()

	fileServer:= http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static",http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static",fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCrete)

	return mux
	// return app.logRequest(secureHeaders(mux))
}
