package main

import (
	"fmt"
	"net/http"
	"strconv"
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

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err!= nil || id<1 {
		http.NotFound(w,r)
		return
	}

	// w.Write([]byte("Display a specific snippet..."))

	fmt.Fprintf(w, "Display a specific snippet...%d", id)
}

func snippetCrete(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost {
		w.Header().Set("Allow","POST")
		// w.WriteHeader(405)
		//  w.Write([]byte("Method Not Allowed"))

		// replacing above two commented lines with the below one

		http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}