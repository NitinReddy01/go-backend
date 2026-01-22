package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"unified_platform/internal/errs"

	"github.com/go-playground/validator/v10"
)

type Validatable interface {
	Validate() error
}

func BindAndValidateBody(r *http.Request, payload Validatable) error {
	if r.Body == nil {
		return errs.NewBadRequestError("request body is required", true, nil, nil, nil)
	}

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(payload); err != nil {

		var syntaxErr *json.SyntaxError
		var typeErr *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxErr):
			return errs.NewBadRequestError(
				"invalid JSON syntax",
				true,
				nil,
				nil,
				nil,
			)

		case errors.As(err, &typeErr):
			return errs.NewBadRequestError(
				fmt.Sprintf("%s must be a %s", typeErr.Field, typeErr.Type),
				true,
				nil,
				[]errs.FieldError{
					{
						Field: strings.ToLower(typeErr.Field),
						Error: fmt.Sprintf("must be a %s", typeErr.Type),
					},
				},
				nil,
			)

		case strings.HasPrefix(err.Error(), "json: unknown field"):
			field := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return errs.NewBadRequestError(
				"unknown field in request body",
				true,
				nil,
				[]errs.FieldError{
					{
						Field: strings.Trim(field, `"`),
						Error: "is not allowed",
					},
				},
				nil,
			)

		default:
			return errs.NewBadRequestError(
				"invalid request body",
				true,
				nil,
				nil,
				nil,
			)
		}
	}

	if err := payload.Validate(); err != nil {
		msg, fields := extractValidationErrors(err)
		return errs.NewBadRequestError(msg, true, nil, fields, nil)
	}

	return nil
}

func extractValidationErrors(err error) (string, []errs.FieldError) {
	var fieldErrors []errs.FieldError

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {

	}

	for _, err := range validationErrors {
		field := strings.ToLower(err.Field())
		var msg string
		switch err.Tag() {
		case "required":
			msg = "is required"
		case "email":
			msg = "must be a valid email address"
		case "oneof":
			msg = fmt.Sprintf("must be one of: %s", err.Param())
		case "uuid":
			msg = "must be a valid UUID"
		case "max":
			if err.Kind() == reflect.String {
				msg = fmt.Sprintf("must not exceed %s characters", err.Param())
			} else {
				msg = fmt.Sprintf("must not exceed %s", err.Param())
			}
		case "min":
			if err.Kind() == reflect.String {
				msg = fmt.Sprintf("must be at least %s characters", err.Param())
			} else {
				msg = fmt.Sprintf("must be at least %s", err.Param())
			}
		case "dive":
			msg = "some items are invalid"
		}

		fieldErrors = append(fieldErrors, errs.FieldError{
			Field: field,
			Error: msg,
		})

	}

	return "Validation Failed", fieldErrors
}
