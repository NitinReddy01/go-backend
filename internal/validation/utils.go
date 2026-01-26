package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/NitinReddy01/go-backend/internal/errs"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type Validatable interface {
	Validate() error
}

type CustomValidationError struct {
	Field   string
	Message string
}

type CustomValidationErrors []CustomValidationError

func (c CustomValidationErrors) Error() string {
	return "Validation failed"
}

func BindAndValidate(c *echo.Context, payload Validatable) error {

	if err := c.Bind(payload); err != nil {

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
				"invalid field type",
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
		customValidationErrors := err.(CustomValidationErrors)
		for _, err := range customValidationErrors {
			fieldErrors = append(fieldErrors, errs.FieldError{
				Field: err.Field,
				Error: err.Message,
			})
		}
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
