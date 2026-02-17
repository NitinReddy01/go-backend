package errs

import "net/http"

func NewUnauthorizedError(message string) *HTTPError {
	return &HTTPError{
		Code:    MakeUpperCaseWithUnderscores(http.StatusText(http.StatusUnauthorized)),
		Message: message,
		Status:  http.StatusUnauthorized,
	}
}

func NewForbiddenError(message string) *HTTPError {
	return &HTTPError{
		Code:    MakeUpperCaseWithUnderscores(http.StatusText(http.StatusForbidden)),
		Message: message,
		Status:  http.StatusForbidden,
	}
}

func NewBadRequestError(message string, code *string, errors []FieldError) *HTTPError {
	formattedCode := MakeUpperCaseWithUnderscores(http.StatusText(http.StatusBadRequest))

	if code != nil {
		formattedCode = *code
	}

	return &HTTPError{
		Code:    formattedCode,
		Message: message,
		Status:  http.StatusBadRequest,
		Errors:  errors,
	}
}

func NewNotFoundError(message string, code *string) *HTTPError {
	formattedCode := MakeUpperCaseWithUnderscores(http.StatusText(http.StatusNotFound))

	if code != nil {
		formattedCode = *code
	}

	return &HTTPError{
		Code:    formattedCode,
		Message: message,
		Status:  http.StatusNotFound,
	}
}

func NewInternalServerError() *HTTPError {
	return &HTTPError{
		Code:    MakeUpperCaseWithUnderscores(http.StatusText(http.StatusInternalServerError)),
		Message: http.StatusText(http.StatusInternalServerError),
		Status:  http.StatusInternalServerError,
	}
}
