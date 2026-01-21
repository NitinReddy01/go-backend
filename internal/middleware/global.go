package middleware

import (
	"net/http"
)

func GlobalErrorHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if err := h(w, r); err != nil {
		// 	// handleError(w, r, err)
		// }
	}
}
