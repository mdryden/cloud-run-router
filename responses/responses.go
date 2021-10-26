package responses

import (
	"encoding/json"
	"net/http"
)

func error(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func Ok(w http.ResponseWriter, payload interface{}) {
	respondWithJSON(w, http.StatusOK, payload)
}

func NotFound(w http.ResponseWriter, message string) {
	error(w, http.StatusNotFound, message)
}

func BadRequest(w http.ResponseWriter, message string) {
	error(w, http.StatusBadRequest, message)
}
