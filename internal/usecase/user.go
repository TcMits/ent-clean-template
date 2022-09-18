package usecase

import (
	"context"
	"errors"

	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
)

func NewIsAuthenticatedPermissionChecker() UserPermissionCheckerUseCase[*model.User] {
	return NewDisallowZeroPermissionChecker[*model.User](useCaseModel.NewUseCaseError(
		errors.New("usecase - NewIsAuthenticatedPermissionChecker: User is not authenticated"),
		"internal.usecase.user.NewIsAuthenticatedPermissionChecker.",
		"Permission denied",
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
					errors.New("usecase - NewIsSuperuserPermissionChecker: User is not superuser"),
					"internal.usecase.user.NewIsSuperuserPermissionChecker.",
					"Permission denied",
					PermissionDeniedError,
				)
			},
		},
	)
}

// var (
// 	_wrapReadUserUserAsSuperuserUseCaseListError = func(err error) error {
// 		return useCaseModel.NewUseCaseError(
// 			fmt.Errorf("usecase - NewReadUserUserAsSuperuserUseCase: %w", err),
// 			"internal.usecase.user.NewReadUserUserAsSuperuserUseCase.ReadUserUserAsSuperuserUseCaseListError",
// 			"Can't get users",
// 			DBError,
// 		)
// 	}
// 	_wrapReadUserUserAsSuperuserUseCaseGetError = func(err error) error {
// 		return useCaseModel.NewUseCaseError(
// 			fmt.Errorf("usecase - NewReadUserUserAsSuperuserUseCase: %w", err),
// 			"internal.usecase.user.NewReadUserUserAsSuperuserUseCase.ReadUserUserAsSuperuserUseCaseGetError",
// 			"Can't get user",
// 			DBError,
// 		)
// 	}
// 	_wrapReadUserUserAsSuperuserUseCaseCountError = func(err error) error {
// 		return useCaseModel.NewUseCaseError(
// 			fmt.Errorf("usecase - NewReadUserUserAsSuperuserUseCase: %w", err),
// 			"internal.usecase.user.NewReadUserUserAsSuperuserUseCase.ReadUserUserAsSuperuserUseCaseCountError",
// 			"Can't count user",
// 			DBError,
// 		)
// 	}
// )
//
// func NewReadUserUserAsSuperuserUseCase(repo interface {
// 	repository.ListModelRepository[*model.User, *model.UserOrderInput, *model.UserWhereInput]
// 	repository.GetModelRepository[*model.User, *model.UserWhereInput]
// 	repository.CountModelRepository[*model.UserWhereInput]
// }) interface {
// 	ListModelUseCase[*model.User, *struct{}, *model.UserWhereInput]
// 	GetModelUseCase[*model.User, *model.UserWhereInput]
// 	CountModelUseCase[*model.UserWhereInput]
// 	SerializeModelUseCase[*model.User, map[string]any]
// } {
// 	columns := []string{
// 		user.FieldID,
// 		user.FieldCreateTime,
// 		user.FieldUpdateTime,
// 		user.FieldUsername,
// 		user.FieldFirstName,
// 		user.FieldLastName,
// 		user.FieldEmail,
// 		user.FieldIsStaff,
// 		user.FieldIsSuperuser,
// 		user.FieldIsActive,
// 	}
//
// 	return &struct {
// 		*listModelUseCase[*model.User, *struct{}, *model.UserWhereInput, *model.UserOrderInput, *model.UserWhereInput]
// 		*getModelUseCase[*model.User, *model.UserWhereInput, *model.UserWhereInput]
// 		*countModelUseCase[*model.UserWhereInput, *model.UserWhereInput]
// 		*model.UserSerializer
// 	}{
// 		&listModelUseCase[*model.User, *struct{}, *model.UserWhereInput, *model.UserOrderInput, *model.UserWhereInput]{
// 			repository:           repo,
// 			toRepoWhereInputFunc: func(ctx context.Context, i *model.UserWhereInput) (*model.UserWhereInput, error) { return i, nil },
// 			toRepoOrderInputFunc: func(ctx context.Context, s *struct{}) (*model.UserOrderInput, error) {
// 				return model.DefaultUserOrderInput, nil
// 			},
// 			wrapListErrorFunc: _wrapReadUserUserAsSuperuserUseCaseListError,
// 		},
// 		&getModelUseCase[*model.User, *model.UserWhereInput, *model.UserWhereInput]{
// 			repository:           repo,
// 			toRepoWhereInputFunc: func(ctx context.Context, i *model.UserWhereInput) (*model.UserWhereInput, error) { return i, nil },
// 			wrapGetErrorFunc:     _wrapReadUserUserAsSuperuserUseCaseGetError,
// 		},
// 		&countModelUseCase[*model.UserWhereInput, *model.UserWhereInput]{
// 			repository:           repo,
// 			toRepoWhereInputFunc: func(ctx context.Context, i *model.UserWhereInput) (*model.UserWhereInput, error) { return i, nil },
// 			wrapCountErrorFunc:   _wrapReadUserUserAsSuperuserUseCaseCountError,
// 		},
// 		ent.NewUserSerializer(nil, columns...),
// 	}
// }
