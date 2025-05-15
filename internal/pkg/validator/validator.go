package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin/binding"

	"github.com/FauzanAr/fleet-management/internal/pkg/wrapper"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Engine() interface{} {
	return cv.validator
}

func formatValidationErrors(errs validator.ValidationErrors) error {
	var messages []string

	for _, err := range errs {
		var message string
		param := err.Param()
		fieldName := err.Field()
		tag := err.Tag()

		message = fmt.Sprintf("[%s] must %s", fieldName, tag)
		if param != "" {
			message = message + fmt.Sprintf(" %s", param)
		}

		messages = append(messages, message)
	}

	return wrapper.ValidationError(strings.Join(messages, ", "))
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) ValidateStruct(obj interface{}) error {
	if err := cv.validator.Struct(obj); err != nil {
		return formatValidationErrors(err.(validator.ValidationErrors))
	}
	return nil
}

func InitCustomValidator() {
	binding.Validator = NewValidator()
}