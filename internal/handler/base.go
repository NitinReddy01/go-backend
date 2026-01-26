package handler

import (
	"context"
	"errors"
	"net/http"
	"unified_platform/internal/errs"
	"unified_platform/internal/server/json"
	"unified_platform/internal/validation"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

// handling the error
func HandleError(w http.ResponseWriter, err error) {
	var httpErr *errs.HTTPError
	if errors.As(err, &httpErr) {
		json.Error(w, httpErr)
		return
	}

	// fallback
	json.Error(w, errs.NewInternalServerError())
}

// catching error
func Handle(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			HandleError(w, err)
		}
	}
}

func HandleBody[Req validation.Validatable, Res any](
	handler func(ctx context.Context, req Req) (*Res, error),
	req Req,
	status int,
) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {

		if err := validation.BindAndValidateBody(r, req); err != nil {
			return err
		}

		res, err := handler(r.Context(), req)

		if err != nil {
			return err
		}

		json.JSON(w, status, res)

		return nil
	}
}
