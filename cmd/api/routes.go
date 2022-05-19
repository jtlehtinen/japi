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
	"net/http"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()
	router.Handle("/", app.cors(app.catchAll))
	router.Handle("/docs", app.cors(app.showDocs))
	router.Handle("/swagger.json", app.cors(app.getDocsJSON))
	router.Handle("/health", app.cors(app.healthHandler))

	return router
}
