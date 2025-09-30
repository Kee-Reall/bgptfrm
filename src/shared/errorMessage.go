package shared

import (
	"fmt"
	"strings"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

func (em *ErrorMessage) Error() string {
	return fmt.Sprintf("error in field %s with message %s;", em.Field, em.Message)
}

func MakeErrorMessage(message, field string) ErrorMessage {
	return ErrorMessage{message, field}
}

type ValidationResponse struct {
	Errors []ErrorMessage `json:"errorsMessages"`
}

func (v *ValidationResponse) Error() string {
	var acc strings.Builder
	if len(v.Errors) == 0 {
		return acc.String()
	}
	for _, em := range v.Errors {
		acc.WriteString(em.Error())
	}
	return acc.String()
}

func NewValidationResponse(ems ...ErrorMessage) *ValidationResponse {
	return &ValidationResponse{ems}
}
