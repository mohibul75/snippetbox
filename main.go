package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){


	// Check the url match withes exactly "/"  or not
	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}

	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCrete(w http.ResponseWriter, r *http.Request){

	if r.Method != "POST"{
		w.Header().Set("Allow","POST")
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main(){
	mux:= http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/snippet/view",snippetView)
	mux.HandleFunc("/snippet/create",snippetCrete)

	log.Println("Starting Server on : 4000 ")
	err:= http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}