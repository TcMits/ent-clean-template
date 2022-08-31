package usecase

import (
	"context"
	"fmt"

	"github.com/TcMits/ent-clean-template/internal/repository"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
)

var (
	_wrapStartDeleteTransactionError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("deleteModelInTransactionUseCase - Delete - u.transactionRepository.Start: %w", err),
			"internal.usecase.delete.deleteModelInTransactionUseCase.Delete.StartDeleteTransactionError",
			"Can't delete now",
			DBError,
		)
	}
	_wrapCommitDeleteError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("deleteModelInTransactionUseCase - Delete - u.transactionRepository.Commit: %w", err),
			"internal.usecase.delete.deleteModelInTransactionUseCase.Delete.CommitDeleteError",
			"Can't delete now",
			DBError,
		)
	}
	_wrapRollbackDeleteError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("deleteModelInTransactionUseCase - Delete - u.transactionRepository.Rollback: %w", err),
			"internal.usecase.delete.deleteModelInTransactionUseCase.Delete.RollbackDeleteError",
			"Can't delete now",
			DBError,
		)
	}
)

type deleteModelUseCase[ModelType, WhereInputType, RepoWhereInputType any] struct {
	repository           repository.DeleteModelRepository[ModelType]
	getRepository        repository.GetModelRepository[ModelType, RepoWhereInputType]
	toRepoWhereInputFunc ConverFunc[WhereInputType, RepoWhereInputType]
	wrapGetErrorFunc     func(error) error
	wrapDeleteErrorFunc  func(error) error
}

type deleteModelInTransactionUseCase[ModelType, WhereInputType, RepoWhereInputType any] struct {
	repository            repository.DeleteWithClientModelRepository[ModelType]
	getRepository         repository.GetWithClientModelRepository[ModelType, RepoWhereInputType]
	transactionRepository repository.TransactionRepository
	toRepoWhereInputFunc  ConverFunc[WhereInputType, RepoWhereInputType]
	selectForUpdate       bool
	wrapGetErrorFunc      func(error) error
	wrapDeleteErrorFunc   func(error) error
}

func (u *deleteModelUseCase[ModelType, WhereInputType, _]) GetAndDelete(
	ctx context.Context, input WhereInputType,
) error {
	repoWhereInput, err := u.toRepoWhereInputFunc(input)
	if err != nil {
		return err
	}
	instance, err := u.getRepository.Get(ctx, repoWhereInput)
	if err != nil {
		return u.wrapGetErrorFunc(err)
	}
	err = u.repository.Delete(ctx, instance)
	if err != nil {
		return u.wrapDeleteErrorFunc(err)
	}
	return nil
}

func (u *deleteModelInTransactionUseCase[ModelType, WhereInputType, _]) GetAndDelete(
	ctx context.Context, input WhereInputType,
) error {
	tx, err := u.transactionRepository.Start(ctx)
	if err != nil {
		return _wrapStartDeleteTransactionError(err)
	}
	client := tx.Client()
	repoWhereInput, err := u.toRepoWhereInputFunc(input)
	if err != nil {
		return err
	}

	// get instance
	instance, err := u.getRepository.GetWithClient(ctx, client, repoWhereInput, u.selectForUpdate)
	if err != nil {
		if rerr := u.transactionRepository.Rollback(tx); rerr != nil {
			return _wrapRollbackDeleteError(err)
		}
		return u.wrapGetErrorFunc(err)
	}

	// delete instance
	err = u.repository.DeleteWithClient(ctx, client, instance)
	if err != nil {
		if rerr := u.transactionRepository.Rollback(tx); rerr != nil {
			return _wrapRollbackDeleteError(err)
		}
		return u.wrapDeleteErrorFunc(err)
	}

	if err = u.transactionRepository.Commit(tx); err != nil {
		return _wrapCommitDeleteError(err)
	}
	return nil
}
