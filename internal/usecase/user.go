package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/TcMits/ent-clean-template/copygen"
	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/ent/user"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/lazy"
)

const (
	_currentUserKey          = "User"
	_HALIDPublicMeColumnName = "self"
)

var (
	_wrapGetPublicMeUseUserIsNotAuthenticatedError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("getPublicMeUseCase - Get - ctx.Value: %w", err),
			"internal.usecase.user.getPublicMeUseCase.Get.UserIsNotAuthenticatedError",
			"Permission denied",
			PermissionDeniedError,
		)
	}
	_wrapUpdatePublicMeUseCaseUpdateError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("usecase - NewPublicMeUseCase: %w", err),
			"internal.usecase.user.NewPublicMeUseCase.UpdateError",
			"Can't update now",
			DBError,
		)
	}
	_wrapValidateUpdateInputPublicMeUseCaseEmailIsAlreadyRegisteredError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf(
				"validateUpdateInputPublicMeUseCase - validateEmail - u.repository.Get: %w",
				err,
			),
			"internal.usecase.user.validateUpdateInputPublicMeUseCase.validateEmail.EmailIsAlreadyRegisteredError",
			"Email is registered",
			ValidationError,
		)
	}
	_wrapValidateUpdateInputPublicMeUseCaseUsernameIsAlreadyRegisteredError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf(
				"validateUpdateInputPublicMeUseCase - validateUsername - u.repository.Get: %w",
				err,
			),
			"internal.usecase.user.validateUpdateInputPublicMeUseCase.validateUsername.UsernameIsAlreadyRegisteredError",
			"Username is registered",
			ValidationError,
		)
	}
	_wrapIsAuthenticatedPermissionCheckerUserIsNotAuthenticatedError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("usecase - NewIsAuthenticatedPermissionChecker: %w", err),
			"internal.usecase.user.NewIsAuthenticatedPermissionChecker.UserIsNotAuthenticatedError",
			"Permission denied",
			PermissionDeniedError,
		)
	}
	_wrapIsSuperuserPermissionCheckerUserIsNotSuperuserError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("usecase - NewIsSuperuserPermissionChecker: %w", err),
			"internal.usecase.user.NewIsSuperuserPermissionChecker.UserIsNotSuperuserError",
			"Permission denied",
			PermissionDeniedError,
		)
	}
)

type (
	getPublicMeUseCase                 struct{}
	validateUpdateInputPublicMeUseCase struct {
		repository repository.GetModelRepository[*model.User, *model.UserWhereInput]
	}
)

func NewIsAuthenticatedPermissionChecker() UserPermissionCheckerUseCase[*model.User] {
	return NewDisallowZeroPermissionChecker[*model.User](
		_wrapIsAuthenticatedPermissionCheckerUserIsNotAuthenticatedError(
			errors.New("User is not authenticated"),
		),
	)
}

func NewIsSuperuserPermissionChecker() UserPermissionCheckerUseCase[*model.User] {
	return NewIsAuthenticatedPermissionChecker().And(
		&basePermissionCheckerUseCase[*model.User]{
			checkFunc: func(ctx context.Context, u *model.User) error {
				if u.IsSuperuser {
					return nil
				}
				return _wrapIsSuperuserPermissionCheckerUserIsNotSuperuserError(
					errors.New("User is not superuser"),
				)
			},
		},
	)
}

func (u *getPublicMeUseCase) Get(ctx context.Context, i *struct{}) (*model.User, error) {
	user, ok := ctx.Value(_currentUserKey).(lazy.LazyObject[*model.User])
	if !ok {
		return nil, _wrapGetPublicMeUseUserIsNotAuthenticatedError(
			errors.New("User is not authenticated"),
		)
	}
	return user.Value(), nil
}

func (u *validateUpdateInputPublicMeUseCase) validateEmail(
	ctx context.Context,
	instance *model.User,
	i *useCaseModel.PublicMeUseCaseUpdateInput,
) error {
	if i.Email != nil && instance.Email != *i.Email {
		_, err := u.repository.Get(ctx, &model.UserWhereInput{Email: i.Email})
		if err == nil {
			return _wrapValidateUpdateInputPublicMeUseCaseEmailIsAlreadyRegisteredError(
				fmt.Errorf("User is already registered with email %s", *i.Email),
			)
		}
	}
	return nil
}

func (u *validateUpdateInputPublicMeUseCase) validateUsername(
	ctx context.Context,
	instance *model.User,
	i *useCaseModel.PublicMeUseCaseUpdateInput,
) error {
	if i.Username != nil && instance.Username != *i.Username {
		_, err := u.repository.Get(ctx, &model.UserWhereInput{Username: i.Username})
		if err == nil {
			return _wrapValidateUpdateInputPublicMeUseCaseUsernameIsAlreadyRegisteredError(
				fmt.Errorf("User is already registered with username %s", *i.Username),
			)
		}
	}
	return nil
}

func (u *validateUpdateInputPublicMeUseCase) Validate(
	ctx context.Context,
	instance *model.User,
	i *useCaseModel.PublicMeUseCaseUpdateInput,
) (*model.UserUpdateInput, error) {
	// validate steps
	steps := append(
		make(
			[]func(context.Context, *model.User, *useCaseModel.PublicMeUseCaseUpdateInput) error,
			0,
			2,
		),
		u.validateEmail,
		u.validateUsername,
	)

	// Validate
	for _, step := range steps {
		err := step(ctx, instance, i)
		if err != nil {
			return nil, err
		}
	}
	result := new(model.UserUpdateInput)
	copygen.PublicMeUseCaseUpdateInputToUserUpdateInput(result, i)
	return result, nil
}

func NewPublicMeUseCase(
	getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput],
	updateRepository repository.UpdateModelRepository[*model.User, *model.UserUpdateInput],
	getURL func(*url.URL, url.Values, ...any) string,
) interface {
	GetModelUseCase[*model.User, *struct{}]
	GetAndUpdateModelUseCase[*model.User, *struct{}, *useCaseModel.PublicMeUseCaseUpdateInput]
	SerializeModelUseCase[*model.User, map[string]any]
} {
	columns := []string{
		user.FieldID,
		user.FieldCreateTime,
		user.FieldUpdateTime,
		user.FieldUsername,
		user.FieldFirstName,
		user.FieldLastName,
		user.FieldEmail,
		user.FieldIsStaff,
		user.FieldIsSuperuser,
		user.FieldIsActive,
	}

	// HAL me Serialize
	HALField := newHALIDModelSerializerField[*model.User](getURL, nil)
	customColumns := map[string]func(context.Context, *model.User) any{
		_HALIDPublicMeColumnName: func(ctx context.Context, u *model.User) any {
			return HALField.Serialize(ctx, u)
		},
	}

	// get me usecase
	getter := &getPublicMeUseCase{}

	// update me usecase
	v := &validateUpdateInputPublicMeUseCase{repository: getRepository}
	updater := &updateModelUseCase[*model.User, *struct{}, *useCaseModel.PublicMeUseCaseUpdateInput, *model.UserUpdateInput]{
		getUseCase:          getter,
		repository:          updateRepository,
		validateFunc:        v.Validate,
		wrapUpdateErrorFunc: _wrapUpdatePublicMeUseCaseUpdateError,
	}

	return &struct {
		*getPublicMeUseCase
		*updateModelUseCase[*model.User, *struct{}, *useCaseModel.PublicMeUseCaseUpdateInput, *model.UserUpdateInput]
		*model.UserSerializer
	}{
		getter,
		updater,
		ent.NewUserSerializer(customColumns, columns...),
	}
}
