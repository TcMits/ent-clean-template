package usecase

import "github.com/nicksnyder/go-i18n/v2/i18n"

var (
	_defaultInvalidErrorMsg = &i18n.Message{
		ID:    "pkg.entity.model.usecase.user.PublicMeUseCaseUpdateInput.GetErrorMessageFromStructField.InvalidError",
		Other: "One or more fields failed to be validated",
	}
	_publicMeUseCaseUpdateInputInvalidUsernameMsg = &i18n.Message{
		ID:    "pkg.entity.model.usecase.user.PublicMeUseCaseUpdateInput.GetErrorMessageFromStructField.InvalidUsernameError",
		Other: "Invalid username",
	}
	_publicMeUseCaseUpdateInputInvalidFirstNameMsg = &i18n.Message{
		ID:    "pkg.entity.model.usecase.user.PublicMeUseCaseUpdateInput.GetErrorMessageFromStructField.InvalidFirstNameError",
		Other: "Invalid first name",
	}
	_publicMeUseCaseUpdateInputInvalidLastNameMsg = &i18n.Message{
		ID:    "pkg.entity.model.usecase.user.PublicMeUseCaseUpdateInput.GetErrorMessageFromStructField.InvalidLastNameError",
		Other: "Invalid last name",
	}
	_publicMeUseCaseUpdateInputInvalidEmailMsg = &i18n.Message{
		ID:    "pkg.entity.model.usecase.user.PublicMeUseCaseUpdateInput.GetErrorMessageFromStructField.InvalidEmailError",
		Other: "Invalid email",
	}
)

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
		return _publicMeUseCaseUpdateInputInvalidUsernameMsg
	case "FirstName":
		return _publicMeUseCaseUpdateInputInvalidFirstNameMsg
	case "LastName":
		return _publicMeUseCaseUpdateInputInvalidLastNameMsg
	case "Email":
		return _publicMeUseCaseUpdateInputInvalidEmailMsg
	default:
		return _defaultInvalidErrorMsg
	}
}
