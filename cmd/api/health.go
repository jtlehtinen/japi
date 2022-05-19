package main

import "net/http"

// swagger:route GET /health healthcheck
//
// Get information about the application status, operating environment and version.
//
// Responses:
//   200:
//   500:
func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// @TODO: JSON response in errors cases
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body := `{"status": "up", "environment": "development", "version": "0.0.1"}`

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(body))
}
