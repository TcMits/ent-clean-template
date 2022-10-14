package usecase

import "github.com/nicksnyder/go-i18n/v2/i18n"

var (
	// i18n message
	_canNotUpdateNowMessage = &i18n.Message{
		ID:    "internal.usecase.can_not_update_now",
		Other: "Can't update now",
	}
	_incorrectLoginInputMessage = &i18n.Message{
		ID:    "internal.usecase.incorrect_login_input",
		Other: "Your username or password is incorrect",
	}
	_canNotLoginNowMessage = &i18n.Message{
		ID:    "internal.usecase.can_not_login_now",
		Other: "Can't login now",
	}
	_authenticationFailedMessage = &i18n.Message{
		ID:    "internal.usecase.authentication_failed",
		Other: "Authentication failed",
	}
	_canNotDeleteNowMessage = &i18n.Message{
		ID:    "internal.usecase.can_not_delete_now",
		Other: "Can't delete now",
	}
	_canNotCreateNowMessage = &i18n.Message{
		ID:    "internal.usecase.can_not_create_now",
		Other: "Can't create now",
	}
	_permissionDeniedMessage = &i18n.Message{
		ID:    "internal.usecase.permission_denied",
		Other: "Permission denied",
	}
	_emailIsRegisteredMessage = &i18n.Message{
		ID:    "internal.usecase.email_is_registered",
		Other: "Email is registered",
	}
	_usernameIsRegisteredMessage = &i18n.Message{
		ID:    "internal.usecase.username_is_registered",
		Other: "Username is registered",
	}
)
