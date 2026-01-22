package errs

import (
	"strings"
)

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type Action struct {
	Message string `json:"message"`
	Value   string `json:"value"`
}

type HTTPError struct {
	Code     string       `json:"code"`
	Message  string       `json:"message"`
	Status   int          `json:"status"`
	Errors   []FieldError `json:"errors"`
	Override bool         `json:"override"`
	Action   *Action      `json:"action"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

func MakeUpperCaseWithUnderscores(str string) string {
	return strings.ToUpper(strings.ReplaceAll(str, " ", "_"))
}
