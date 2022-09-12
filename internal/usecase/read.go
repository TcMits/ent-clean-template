package usecase

import (
	"context"

	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/pkg/tool/generic"
)

type ConverFunc[FromType, ToType any] func(FromType) (ToType, error)

type getModelUseCase[ModelType, WhereInputType, RepoWhereInputType any] struct {
	repository           repository.GetModelRepository[ModelType, RepoWhereInputType]
	toRepoWhereInputFunc ConverFunc[WhereInputType, RepoWhereInputType]
	wrapGetErrorFunc     func(error) error
}

type countModelUseCase[WhereInputType, RepoWhereInputType any] struct {
	repository           repository.CountModelRepository[RepoWhereInputType]
	toRepoWhereInputFunc ConverFunc[WhereInputType, RepoWhereInputType]
	wrapCountErrorFunc   func(error) error
}

type listModelUseCase[ModelType, OrderInputType, WhereInputType, RepoOrderInputType, RepoWhereInputType any] struct {
	repository           repository.ListModelRepository[ModelType, RepoOrderInputType, RepoWhereInputType]
	toRepoWhereInputFunc ConverFunc[WhereInputType, RepoWhereInputType]
	toRepoOrderInputFunc ConverFunc[OrderInputType, RepoOrderInputType]
	wrapListErrorFunc    func(error) error
}

func (u *getModelUseCase[ModelType, FilterInputType, _]) Get(
	ctx context.Context, input FilterInputType,
) (ModelType, error) {
	whereInput, err := u.toRepoWhereInputFunc(input)
	if err != nil {
		return generic.Zero[ModelType](), err
	}
	instance, err := u.repository.Get(ctx, whereInput)
	if err != nil {
		return generic.Zero[ModelType](), u.wrapGetErrorFunc(err)
	}
	return instance, nil
}

func (u *countModelUseCase[FilterInputType, _]) Count(
	ctx context.Context, input FilterInputType,
) (int, error) {
	whereInput, err := u.toRepoWhereInputFunc(input)
	if err != nil {
		return 0, err
	}
	count, err := u.repository.Count(ctx, whereInput)
	if err != nil {
		return 0, u.wrapCountErrorFunc(err)
	}
	return count, nil
}

func (u *listModelUseCase[ModelType, OrderInputType, WhereInputType, _, _]) List(
	ctx context.Context, limit int, offset int, orderInput OrderInputType, whereInput WhereInputType,
) ([]ModelType, error) {
	repoOrderInput, err := u.toRepoOrderInputFunc(orderInput)
	if err != nil {
		return nil, err
	}
	repoWhereInput, err := u.toRepoWhereInputFunc(whereInput)
	if err != nil {
		return nil, err
	}
	instance, err := u.repository.List(ctx, limit, offset, repoOrderInput, repoWhereInput)
	if err != nil {
		return nil, u.wrapListErrorFunc(err)
	}
	return instance, nil
}