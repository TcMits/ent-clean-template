package usecase

const (
	_defaultInvalidErrorTranslateKey                          = "pkg.entity.model.usecase.user.PublicMeUseCaseUpdateInput.GetErrorMessageFromStructField.InvalidError"
	_defaultInvalidErrorMessage                               = "One or more fields failed to be validated"
	_publicMeUseCaseUpdateInputInvalidUsernameTranslateKey    = "pkg.entity.model.usecase.user.PublicMeUseCaseUpdateInput.GetErrorMessageFromStructField.InvalidUsernameError"
	_publicMeUseCaseUpdateInputInvalidUsernameDefaultMessage  = "Invalid username"
	_publicMeUseCaseUpdateInputInvalidFirstNameTranslateKey   = "pkg.entity.model.usecase.user.PublicMeUseCaseUpdateInput.GetErrorMessageFromStructField.InvalidFirstNameError"
	_publicMeUseCaseUpdateInputInvalidFirstNameDefaultMessage = "Invalid first name"
	_publicMeUseCaseUpdateInputInvalidLastNameTranslateKey    = "pkg.entity.model.usecase.user.PublicMeUseCaseUpdateInput.GetErrorMessageFromStructField.InvalidLastNameError"
	_publicMeUseCaseUpdateInputInvalidLastNameDefaultMessage  = "Invalid last name"
	_publicMeUseCaseUpdateInputInvalidEmailTranslateKey       = "pkg.entity.model.usecase.user.PublicMeUseCaseUpdateInput.GetErrorMessageFromStructField.InvalidEmailError"
	_publicMeUseCaseUpdateInputInvalidEmailDefaultMessage     = "Invalid email"
)

type PublicMeUseCaseUpdateInput struct {
	Username  *string `json:"username,omitempty"   form:"username,omitempty"   validate:"alphanum,min=1,max=128"` // username has to be unique
	FirstName *string `json:"first_name,omitempty" form:"first_name,omitempty" validate:"max=128"`
	LastName  *string `json:"last_name,omitempty"  form:"last_name,omitempty"  validate:"max=128"`
	Email     *string `json:"email,omitempty"      form:"email,omitempty"      validate:"email"` // email has to be unique
}

func (i *PublicMeUseCaseUpdateInput) GetErrorMessageFromStructField(
	fieldName string,
) (string, string) {
	switch fieldName {
	case "Username":
		return _publicMeUseCaseUpdateInputInvalidUsernameTranslateKey, _publicMeUseCaseUpdateInputInvalidUsernameDefaultMessage
	case "FirstName":
		return _publicMeUseCaseUpdateInputInvalidFirstNameTranslateKey, _publicMeUseCaseUpdateInputInvalidFirstNameDefaultMessage
	case "LastName":
		return _publicMeUseCaseUpdateInputInvalidLastNameTranslateKey, _publicMeUseCaseUpdateInputInvalidLastNameDefaultMessage
	case "Email":
		return _publicMeUseCaseUpdateInputInvalidEmailTranslateKey, _publicMeUseCaseUpdateInputInvalidEmailDefaultMessage
	default:
		return _defaultInvalidErrorTranslateKey, _defaultInvalidErrorMessage
	}
}
