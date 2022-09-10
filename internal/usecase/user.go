package usecase

import (
	"context"
	"errors"

	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
)

func NewIsAuthenticatedPermissionChecker() UserPermissionCheckerUseCase[*model.User] {
	return NewDisallowZeroPermissionChecker[*model.User](useCaseModel.NewUseCaseError(
		errors.New("usecase - NewIsAuthenticatedPermissionChecker: "),
		"internal.usecase.user.NewIsAuthenticatedPermissionChecker.",
		"Your username or password is incorrect",
		PermissionDeniedError,
	))
}

func NewIsSuperuserPermissionChecker() UserPermissionCheckerUseCase[*model.User] {
	return NewIsAuthenticatedPermissionChecker().And(
		&basePermissionCheckerUseCase[*model.User]{
			checkFunc: func(ctx context.Context, u *model.User) error {
				if u.IsSuperuser {
					return nil
				}
				return useCaseModel.NewUseCaseError(
					errors.New("usecase - NewIsSuperuserPermissionChecker: "),
					"internal.usecase.user.NewIsSuperuserPermissionChecker.",
					"Permission denied",
					PermissionDeniedError,
				)
			},
		},
	)
}
