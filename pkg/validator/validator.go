package validator

import (
	error "errors"
	"log"
	"reflect"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/smartbot/catalog/pkg/errors"
)

func GetMessageByTag(err validator.FieldError) string {
	tag := err.Tag()

	switch tag {
	case "required":
		return "required field"
	case "email":
		return "invalid email format"
	case "oneof":
		return err.Field() + " should be one of " + err.Param()
	case "mobileNo":
		return "invalid mobile number format. (ex: +91-9999999999)"
	case "gte":
		return err.Field() + " should be greater " + err.Param()
	default:
		return tag
	}
}

func isValidMobileNumberWithCC(fl validator.FieldLevel) bool {
	mobileNumber := fl.Field().String()
	if mobileNumber != "" {
		re := regexp.MustCompile(`^\+\d{1,3}-\d{6,15}$`)
		return re.MatchString(mobileNumber)
	}

	return true

}

func ValidateUUID(id string) *errors.ApiError {
	_, err := uuid.Parse(id)
	if err != nil {
		return errors.ValidationError("Invalid UUID, Please provide a valid UUID", []errors.FieldError{})
	}

	return nil
}
func ValidateBody(c *gin.Context, obj any) *errors.ApiError {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.ValidationError("Invalid request payload", []errors.FieldError{})
	}

	var validate = validator.New()
	validate.RegisterValidation("mobileNo", isValidMobileNumberWithCC)

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})

	if err := validate.Struct(obj); err != nil {
		var ve validator.ValidationErrors
		if error.As(err, &ve) {
			validationErrors := []errors.FieldError{}

			for _, e := range ve {
				validationErrors = append(validationErrors, errors.FieldError{Field: e.Field(), Message: GetMessageByTag(e)})
			}

			return errors.ValidationError("Invalid request payload", validationErrors)
		}
		return errors.ValidationError("Invalid request payload", []errors.FieldError{})
	}

	return nil

}

func ValidateQueryParams(c *gin.Context, obj any) *errors.ApiError {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.ValidationError("Invalid request payload", []errors.FieldError{})
	}

	log.Printf("Parsed QueryParams: %+v\n", obj)

	var validate = validator.New()
	validate.RegisterValidation("mobileNo", isValidMobileNumberWithCC)

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("form")
	})

	if err := validate.Struct(obj); err != nil {
		var ve validator.ValidationErrors
		if error.As(err, &ve) {
			validationErrors := []errors.FieldError{}

			for _, e := range ve {
				validationErrors = append(validationErrors, errors.FieldError{Field: e.Field(), Message: GetMessageByTag(e)})
			}

			return errors.ValidationError("Invalid request payload", validationErrors)
		}
		return errors.ValidationError("Invalid request payload", []errors.FieldError{})
	}

	return nil

}
