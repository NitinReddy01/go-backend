package handler

import (
	"errors"
	"net/http"
	"unified_platform/internal/errs"
	"unified_platform/internal/server/json"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

func HandleError(w http.ResponseWriter, err error) {
	var httpErr *errs.HTTPError
	if errors.As(err, &httpErr) {
		json.Error(w, httpErr)
		return
	}

	// fallback
	json.Error(w, errs.NewInternalServerError())
}

func Handle(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			HandleError(w, err)
		}
	}
}
