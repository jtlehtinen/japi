// Japi API:
//   version: 0.0.1
//   title: Japi API
// Schemes: http
// Host: localhost:3001
// BasePath: /
// Produces:
//   - application/json
// swagger:meta
package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (app *application) writeJSON(w http.ResponseWriter, data any) {
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (app *application) catchAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	if r.URL.Path == "/" {
		w.Write([]byte("index page"))
		return
	}

	segments := strings.Split(r.URL.Path, "/")
	if len(segments) == 0 {
		app.writeJSON(w, app.data)
	}

	data := app.data.(map[string]any)
	for _, segment := range segments {
		for k, v := range data {
			if segment == k {
				app.writeJSON(w, v)
				return
			}
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *application) routes() http.Handler {
	router := http.NewServeMux()
	router.Handle("/", app.cors(app.catchAll))
	router.Handle("/v1/docs", app.cors(app.showDocs))
	router.Handle("/v1/swagger.json", app.cors(app.getDocsJSON))
	router.Handle("/v1/health", app.cors(app.healthHandler))

	return router
}
