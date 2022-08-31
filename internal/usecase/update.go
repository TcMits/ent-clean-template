package usecase

import (
	"context"
	"fmt"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/internal/repository"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/generic"
)

var (
	_wrapStartUpdateTransactionError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("updateModelInTransactionUseCase - Update - u.transactionRepository.Start: %w", err),
			"internal.usecase.update.updateModelInTransactionUseCase.Update.StartUpdateTransactionError",
			"Can't update now",
			DBError,
		)
	}
	_wrapCommitUpdateError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("updateModelInTransactionUseCase - Update - u.transactionRepository.Commit: %w", err),
			"internal.usecase.update.updateModelInTransactionUseCase.Update.CommitUpdateError",
			"Can't update now",
			DBError,
		)
	}
	_wrapRollbackUpdateError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("updateModelInTransactionUseCase - Update - u.transactionRepository.Rollback: %w", err),
			"internal.usecase.update.updateModelInTransactionUseCase.Update.RollbackUpdateError",
			"Can't update now",
			DBError,
		)
	}
)

type UpdateValidateFunc[ModelType, CreateInput, RepoCreateInput any] func(ModelType, CreateInput) (RepoCreateInput, error)
type UpdateInTransactionValidateFunc[ModelType, CreateInput, RepoCreateInput any] func(ModelType, CreateInput, *ent.Client) (RepoCreateInput, error)

type updateModelUseCase[ModelType, WhereInputType, UpdateInputType, RepoWhereInputType, RepoUpdateInputType any] struct {
	repository           repository.UpdateModelRepository[ModelType, RepoUpdateInputType]
	getRepository        repository.GetModelRepository[ModelType, RepoWhereInputType]
	toRepoWhereInputFunc ConverFunc[WhereInputType, RepoWhereInputType]
	validateFunc         UpdateValidateFunc[ModelType, UpdateInputType, RepoUpdateInputType]
	wrapGetErrorFunc     func(error) error
	wrapUpdateErrorFunc  func(error) error
}

type updateModelInTransactionUseCase[ModelType, WhereInputType, UpdateInputType, RepoWhereInputType, RepoUpdateInputType any] struct {
	repository            repository.UpdateWithClientModelRepository[ModelType, RepoUpdateInputType]
	getRepository         repository.GetWithClientModelRepository[ModelType, RepoWhereInputType]
	transactionRepository repository.TransactionRepository
	toRepoWhereInputFunc  ConverFunc[WhereInputType, RepoWhereInputType]
	validateFunc          UpdateInTransactionValidateFunc[ModelType, UpdateInputType, RepoUpdateInputType]
	selectForUpdate       bool
	wrapGetErrorFunc      func(error) error
	wrapUpdateErrorFunc   func(error) error
}

func (u *updateModelUseCase[ModelType, WhereInputType, UpdateInputType, _, _]) GetAndUpdate(
	ctx context.Context, whereInput WhereInputType, updateInput UpdateInputType,
) (ModelType, error) {
	repoWhereInput, err := u.toRepoWhereInputFunc(whereInput)
	if err != nil {
		return generic.Zero[ModelType](), err
	}
	instance, err := u.getRepository.Get(ctx, repoWhereInput)
	if err != nil {
		return generic.Zero[ModelType](), u.wrapGetErrorFunc(err)
	}
	repoUpdateInput, err := u.validateFunc(instance, updateInput)
	if err != nil {
		return generic.Zero[ModelType](), err
	}
	instance, err = u.repository.Update(ctx, instance, repoUpdateInput)
	if err != nil {
		return generic.Zero[ModelType](), u.wrapUpdateErrorFunc(err)
	}
	return instance, nil
}

func (u *updateModelInTransactionUseCase[ModelType, WhereInputType, UpdateInputType, _, _]) GetAndUpdate(
	ctx context.Context, whereInput WhereInputType, updateInput UpdateInputType,
) (ModelType, error) {
	tx, err := u.transactionRepository.Start(ctx)
	if err != nil {
		return generic.Zero[ModelType](), _wrapStartUpdateTransactionError(err)
	}

	client := tx.Client()
	repoWhereInput, err := u.toRepoWhereInputFunc(whereInput)
	if err != nil {
		return generic.Zero[ModelType](), err
	}

	// get instance
	instance, err := u.getRepository.GetWithClient(ctx, client, repoWhereInput, u.selectForUpdate)
	if err != nil {
		if rerr := u.transactionRepository.Rollback(tx); rerr != nil {
			return generic.Zero[ModelType](), _wrapRollbackUpdateError(err)
		}
		return generic.Zero[ModelType](), u.wrapGetErrorFunc(err)
	}

	// validate input with old instance
	repoUpdateInput, err := u.validateFunc(instance, updateInput, client)
	if err != nil {
		if rerr := u.transactionRepository.Rollback(tx); rerr != nil {
			return generic.Zero[ModelType](), _wrapRollbackUpdateError(err)
		}
		return generic.Zero[ModelType](), err
	}

	// update instance
	instance, err = u.repository.UpdateWithClient(ctx, client, instance, repoUpdateInput)
	if err != nil {
		if rerr := u.transactionRepository.Rollback(tx); rerr != nil {
			return generic.Zero[ModelType](), _wrapRollbackUpdateError(err)
		}
		return generic.Zero[ModelType](), u.wrapUpdateErrorFunc(err)
	}

	if err = u.transactionRepository.Commit(tx); err != nil {
		return generic.Zero[ModelType](), _wrapCommitUpdateError(err)
	}
	return instance, nil
}
