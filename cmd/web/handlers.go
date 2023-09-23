package main

import (
	"fmt"
	"github.com/mohibul75/snippetbox/internal/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	snippets *models.SnippetModel
}

func NewApplication(ilog *log.Logger, elog *log.Logger, content *models.SnippetModel) *application {
	return &application{
		infoLog:  ilog,
		errorLog: elog,
		snippets: content,
	}
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		app.NoFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.ServerError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.ServerError(w, err)
		return
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NoFound(w)
		return
	}

	fmt.Fprintf(w, "Snnipet showing for id : %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("allow", "POST")
		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.ServerError(w, err)
	}
	w.Write([]byte("Create a new snippet %d"))
}
