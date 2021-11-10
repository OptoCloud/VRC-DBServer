package helpers

import (
	"encoding/json"
	"io"
	"net/http"
)

type GenericResponseMessage struct {
	Ok    bool   `json:"ok"`
	Field string `json:"field"`
	Error string `json:"error"`
}

func WriteGenericError(w http.ResponseWriter, fieldStr string, errorStr string, statusCode int) {
	response := GenericResponseMessage{false, fieldStr, errorStr}
	WriteJson(w, &response, statusCode)
}
func WriteJson(w http.ResponseWriter, v interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}
func WriteGenericOk(w http.ResponseWriter) {
	response := GenericResponseMessage{true, "", ""}
	WriteJson(w, &response, http.StatusOK)
}
func WriteTodo(w http.ResponseWriter, todoStr string) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, todoStr)
}
