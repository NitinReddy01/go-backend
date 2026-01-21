package json

import (
	"encoding/json"
	"log"
	"net/http"
)

type SuccessResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type ErrorDetails map[string]any

type ErrorObject struct {
	Message string       `json:"message"`
	Details ErrorDetails `json:"details"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   ErrorObject `json:"error"`
}

func JSON[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(SuccessResponse[T]{
		Success: true,
		Data:    data,
	}); err != nil {
		log.Printf("response encode failed: %v", err)
	}
}

func OK[T any](w http.ResponseWriter, data T) {
	JSON(w, http.StatusOK, data)
}

func ErrorJSON(w http.ResponseWriter, status int, message string, details ErrorDetails) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Error: ErrorObject{
			Message: message,
			Details: details,
		},
	}); err != nil {
		log.Printf("response encode failed: %v", err)
	}
}

func BadRequest(w http.ResponseWriter, message string, details ErrorDetails) {
	ErrorJSON(w, http.StatusBadRequest, message, details)
}

func Unauthorized(w http.ResponseWriter, message string) {
	ErrorJSON(w, http.StatusUnauthorized, message, nil)
}

func Internal(w http.ResponseWriter) {
	ErrorJSON(w, http.StatusInternalServerError, "Internal server error", nil)
}
