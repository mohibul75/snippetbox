package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application)home(w http.ResponseWriter, r *http.Request){


	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.infoLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application)snippetView(w http.ResponseWriter, r *http.Request){

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err!= nil || id<1 {
		http.NotFound(w,r)
		return
	}

	// w.Write([]byte("Display a specific snippet..."))

	fmt.Fprintf(w, "Display a specific snippet...%d", id)
}

func (app *application)snippetCrete(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost {
		w.Header().Set("Allow","POST")

		// w.WriteHeader(405)
		//  w.Write([]byte("Method Not Allowed"))
		// replacing above two commented lines with the below one

		http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}

	//create one new entry

	title:="0 snail"
	content:= "0 Snail\nClimb Mount,\nBut Slow"
	expires:=7

	id, err:= app.snippet.Insert(title,content,expires)
	if err!=nil{
		//app.serverError(w,err)

		return
	}

	// w.Write([]byte("Create a new snippet..."))

	http.Redirect(w,r,fmt.Sprintf("/snippet/view?id=%d",id),http.StatusSeeOther)
}