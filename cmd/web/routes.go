package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app * application) routes() http.Handler{

	//mux:= http.NewServeMux()

	router:= httprouter.New()

	router.NotFound= http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		app.notFound(w)
	})

	fileServer:=http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet,"/static/*filepath",http.StripPrefix("/static",fileServer))

	router.HandlerFunc(http.MethodGet,"/", app.home)
	router.HandlerFunc(http.MethodGet,"/snippet/view/:id",app.snippetView)
	router.HandlerFunc(http.MethodGet,"/snippet/create",app.snippetCrete)
	router.HandlerFunc(http.MethodPost,"/snippet/create",app.snippetCretePost)


	// fileServer:= http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	// mux.Handle("/static",http.NotFoundHandler())
	// mux.Handle("/static/", http.StripPrefix("/static",fileServer))

	// mux.HandleFunc("/", app.home)
	// mux.HandleFunc("/snippet/view", app.snippetView)
	// mux.HandleFunc("/snippet/create", app.snippetCrete)


	return app.recoverPanic(app.logRequest(secureHeaders(router)))
}
