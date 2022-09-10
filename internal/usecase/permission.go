package usecase

import (
	"context"

	"github.com/TcMits/ent-clean-template/pkg/tool/generic"
)

type basePermissionCheckerUseCase[UserType any] struct {
	checkFunc func(context.Context, UserType) error
}

func (c *basePermissionCheckerUseCase[UserType]) Check(ctx context.Context, u UserType) error {
	return c.checkFunc(ctx, u)
}

func (c *basePermissionCheckerUseCase[UserType]) And(
	checker UserPermissionCheckerUseCase[UserType],
) UserPermissionCheckerUseCase[UserType] {
	return &basePermissionCheckerUseCase[UserType]{
		checkFunc: func(ctx context.Context, u UserType) error {
			if err := c.Check(ctx, u); err != nil {
				return err
			}
			if err := checker.Check(ctx, u); err != nil {
				return err
			}
			return nil
		},
	}
}

func (c *basePermissionCheckerUseCase[UserType]) Or(
	checker UserPermissionCheckerUseCase[UserType],
) UserPermissionCheckerUseCase[UserType] {
	return &basePermissionCheckerUseCase[UserType]{
		checkFunc: func(ctx context.Context, u UserType) error {
			var err error
			if err = c.Check(ctx, u); err == nil {
				return nil
			}
			if err = checker.Check(ctx, u); err == nil {
				return nil
			}
			return err
		},
	}
}

func NewAllowAnyPermissionChecker[T any]() UserPermissionCheckerUseCase[T] {
	return &basePermissionCheckerUseCase[T]{
		checkFunc: func(ctx context.Context, u T) error {
			return nil
		},
	}
}

func NewDisallowAnyPermissionChecker[T any](err error) UserPermissionCheckerUseCase[T] {
	return &basePermissionCheckerUseCase[T]{
		checkFunc: func(ctx context.Context, u T) error {
			return err
		},
	}
}

func NewDisallowZeroPermissionChecker[T comparable](err error) UserPermissionCheckerUseCase[T] {
	return &basePermissionCheckerUseCase[T]{
		checkFunc: func(ctx context.Context, u T) error {
			if u != generic.Zero[T]() {
				return nil
			}
			return err
		},
	}
}
