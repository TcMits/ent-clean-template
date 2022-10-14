package usecase

import "github.com/nicksnyder/go-i18n/v2/i18n"

type PublicMeUseCaseUpdateInput struct {
	Username  *string `json:"username,omitempty"   form:"username,omitempty"   validate:"alphanum,min=1,max=128"` // username has to be unique
	FirstName *string `json:"first_name,omitempty" form:"first_name,omitempty" validate:"max=128"`
	LastName  *string `json:"last_name,omitempty"  form:"last_name,omitempty"  validate:"max=128"`
	Email     *string `json:"email,omitempty"      form:"email,omitempty"      validate:"email"` // email has to be unique
}

func (i *PublicMeUseCaseUpdateInput) GetErrorMessageFromStructField(
	fieldName string,
) *i18n.Message {
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
