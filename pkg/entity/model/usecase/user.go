package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type PublicMeUseCaseUpdateInput struct {
	Username  *string `json:"username,omitempty"   form:"username,omitempty"   validate:"alphanum,min=1,max=128"` // username has to be unique
	FirstName *string `json:"first_name,omitempty" form:"first_name,omitempty" validate:"max=128"`
	LastName  *string `json:"last_name,omitempty"  form:"last_name,omitempty"  validate:"max=128"`
	Email     *string `json:"email,omitempty"      form:"email,omitempty"      validate:"email"` // email has to be unique
}

func (_ *PublicMeUseCaseUpdateInput) GetErrorMessageFromStructField(err error) *i18n.Message {
	fieldName := ""
	if validateErr, ok := err.(validator.FieldError); ok {
		fieldName = validateErr.StructField()
	}

	switch fieldName {
	case "Username":
		return _invalidUsernameMessage
	case "FirstName":
		return _invalidFirstNameMessage
	case "LastName":
		return _invalidLastNameMessage
	case "Email":
		return _invalidEmailMessage
	default:
		return _oneOrMoreFieldsFailedToBeValidatedMessage
	}
}
