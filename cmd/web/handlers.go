package main

import (
	"bytes"
	"fmt"
	"megatroncodrr/snippetbox/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, data *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.errorLog.Printf("no such Template %s exist", name)
		return
	}
	buff := &bytes.Buffer{}
	err := ts.Execute(buff, data)
	if err != nil {
		app.serverError(w, err)
		return
	}
	buff.WriteTo(w)
}
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{Snippets: s}
	app.render(w, r, "home.page.tmpl", data)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{Snippet: s}
	app.render(w, r, "show.page.tmpl", data)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Method", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "oz the wizzard"
	content := "he fell in a pond and then shark ate him. Well, i dont know why was there a shark in the pond."
	expired := "7"
	id, err := app.snippets.Insert(title, content, expired)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.infoLog.Print("redirecting....")
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
