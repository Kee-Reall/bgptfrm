package validation

import (
	"blog-api/src/shared"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
)

var validate = validator.New()

func BindAndValidate[T any](src io.ReadCloser, dst T) *shared.ValidationResponse {
	if err := json.NewDecoder(src).Decode(dst); err != nil {
		return shared.NewValidationResponse(shared.MakeErrorMessage("invalid JSON", ""))
	}

	if err := validate.Struct(dst); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errs := make([]shared.ErrorMessage, 0, len(ve))
			for _, fe := range ve {

				em := shared.MakeErrorMessage(
					fmt.Sprintf("faild on '%s' tag with '%v'", fe.Tag(), fe.Value()),
					fe.Field())

				errs = append(errs, em)
			}
			return shared.NewValidationResponse(errs...)
		}
	}
	return nil
}
