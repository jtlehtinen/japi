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
	"errors"
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

var notFoundError = errors.New("not found")

// findFromArray finds object from the data slice
// that has an 'id' field with the value wantId
func findFromArray(data []any, wantId string) (any, error) {
	// @TODO: Handle float64 id?
	// search for an object having 'id' field
	// matching the given segment
	for _, d := range data {
		if m, ok := d.(map[string]any); ok {
			if id, ok := m["id"]; ok {
				if id == wantId {
					return d, nil
				}
			}
		}
	}

	return nil, notFoundError
}

func findFromMap(data map[string]any, segment string) (any, error) {
	result, ok := data[segment]
	if !ok {
		return nil, notFoundError
	}
	return result, nil
}

func findData(data any, segments []string) (any, error) {
	var err error
	for _, segment := range segments {
		switch t := data.(type) {
		case map[string]any:
			data, err = findFromMap(t, segment)
		case []any:
			data, err = findFromArray(t, segment)
		default:
			err = notFoundError
		}

		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (app *application) catchAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	if r.URL.Path == "/" {
		app.writeJSON(w, app.data)
		return
	}

	segments := strings.Split(r.URL.Path, "/")
	if len(segments) != 0 && segments[0] == "" {
		segments = segments[1:]
	}

	result, err := findData(app.data, segments)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	app.writeJSON(w, result)
}

func (app *application) routes() http.Handler {
	router := http.NewServeMux()
	router.Handle("/", app.cors(app.catchAll))
	router.Handle("/docs", app.cors(app.showDocs))
	router.Handle("/swagger.json", app.cors(app.getDocsJSON))
	router.Handle("/health", app.cors(app.healthHandler))

	return router
}
