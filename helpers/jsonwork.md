// ===== JSON HELPERS ==================================================================================================

func writeJSON(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func writeJSONError(w http.ResponseWriter, statusCode int, message string) {
	if message == "" {
		message = http.StatusText(statusCode)
	}

	writeJSON(
		w,
		statusCode,
		&ErrorResponse{
			StatusCode: statusCode,
			Message:    message,
		},
	)
}

func writeJSONNotFound(w http.ResponseWriter) {
	writeJSONError(w, http.StatusNotFound, "")
}

func writeUnexpectedError(w http.ResponseWriter, err error) {
	writeJSONError(w, http.StatusInternalServerError, err.Error())
}
