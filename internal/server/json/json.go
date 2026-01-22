package json

import (
	"encoding/json"
	"log"
	"net/http"
	"unified_platform/internal/errs"
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

func Error(w http.ResponseWriter, err *errs.HTTPError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)

	if e := json.NewEncoder(w).Encode(err); e != nil {
		log.Printf("response encode failed: %v", e)
	}
}
