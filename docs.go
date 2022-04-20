package main

import (
	"bytes"
	"html/template"
	"net/http"
)

func (app *application) showDocs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// @TODO: JSON response in errors cases
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles("./templates/apidoc.tmpl")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	b := new(bytes.Buffer)
	err = t.Execute(b, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	b.WriteTo(w)
}

func (app *application) getDocsJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// @TODO: JSON response in errors cases
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./static/swagger.json")
}
